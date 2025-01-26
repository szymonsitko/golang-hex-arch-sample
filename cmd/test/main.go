package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	testMode string
	envPath  string
)

type TestMode string

const (
	Integration TestMode = "integration"
	Unit        TestMode = "unit"
)

func main() {
	// Define the root command
	var rootCmd = &cobra.Command{
		Use:   "test",
		Short: "Run tests",
		Run: func(cmd *cobra.Command, args []string) {
			if envPath == "" {
				log.Fatal("Error: --env-path is required")
			}
			if testMode == "" {
				log.Fatal("Error: --mode is required")
			}

			fmt.Printf("Using environment file at: %s\n", envPath)
			fmt.Printf("Running %s tests...\n", testMode)

			// Set the environment variable
			os.Setenv("TEST_CONFIG_FILE_PATH", envPath)

			// Run the tests
			var testCmd *exec.Cmd
			if TestMode(testMode) == Unit {
				testCmd = exec.Command("go", "test", "-v", "-short", "./...")
			} else if TestMode(testMode) == Integration {
				testCmd = exec.Command("go", "test", "-v", "-run", "Integration", "./...")
			} else {
				log.Fatalf("Invalid test mode: %s", testMode)
			}

			testCmd.Stdout = os.Stdout
			testCmd.Stderr = os.Stderr
			err := testCmd.Run()
			if err != nil {
				log.Fatalf("Error running tests: %s", err)
			}
		},
	}

	// Add the --env-path flag to the root command
	rootCmd.Flags().StringVar(&envPath, "env-path", "", "Path to the environment file")
	rootCmd.MarkFlagRequired("env-path")

	// Add the --mode flag to the root command
	rootCmd.Flags().StringVar(&testMode, "mode", "", "Mode of tests to run (unit|integration)")
	rootCmd.MarkFlagRequired("mode")

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
