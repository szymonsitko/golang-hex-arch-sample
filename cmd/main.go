package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/ssitko/hex-domain/config"
	"github.com/ssitko/hex-domain/internal/handlers"
	"github.com/ssitko/hex-domain/internal/infrastructure/persistence"
	"github.com/ssitko/hex-domain/internal/repositories"
	"github.com/ssitko/hex-domain/internal/routers"
	"github.com/ssitko/hex-domain/internal/services"
	"github.com/ssitko/hex-domain/pkg/logger"
)

var (
	db            persistence.DB
	envPath       string
	serviceLogger logger.Logger
)

func init() {
	// Read & parse cmd args
	cmd()

	// Load app config
	err := config.LoadConfig(envPath)
	if err != nil {
		log.Fatalf("invalid config provided %s", err)
	}

	// Setup persistence layer
	db = persistence.NewPersistenceLayer()

	// Setup logger
	serviceLogger = logger.NewLogger()
}

func main() {
	r := gin.Default()

	// Add logger middleware
	r.Use(func(c *gin.Context) {
		serviceLogger.Info(fmt.Sprintf("Request method: %s, time: %s, path: %s", c.Request.Method, time.Now().UTC().Format(time.RFC3339), c.Request.URL))
		c.Next()
	})

	// Initialize layers
	repo := repositories.NewGormAlbumRepository(db)
	service := services.NewAlbumService(repo)
	handler := handlers.NewAlbumHandler(service)

	// Router
	routers.RegisterAlbumHandlers(r, handler)

	r.Run(fmt.Sprintf(":%s", config.GetConfigValue(config.PORT)))
}

func cmd() {
	// Define the root command
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: ".env absolute path location",
		Run: func(cmd *cobra.Command, args []string) {
			if envPath == "" {
				log.Fatal("Error: --env-path is required")
			}
			fmt.Printf("Using environment file at: %s\n", envPath)
		},
	}

	// Add the --env-path flag to the root command
	rootCmd.Flags().StringVar(&envPath, "env-path", "", "Path to the environment file")
	rootCmd.MarkFlagRequired("env-path")

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
