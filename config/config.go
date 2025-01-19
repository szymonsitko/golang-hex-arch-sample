package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_HOST     = "DB_HOST"
	DB_PORT     = "DB_PORT"
	DB_NAME     = "DB_NAME"
	PORT        = "PORT"
)

var REQUIRED_KEYS = []string{
	"DB_USER",
	"DB_PASSWORD",
	"DB_HOST",
	"DB_PORT",
	"DB_NAME",
	"PORT",
}

func LoadConfig(envFilePath string) error {
	viper.SetConfigFile(envFilePath)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error reading config file, %s", err)
	}

	// Validate required config values
	err = validateConfig(REQUIRED_KEYS)
	if err != nil {
		return err
	}
	return nil
}

func GetConfigValue(key string) string {
	return viper.GetString(key)
}

func validateConfig(keys []string) error {
	for _, key := range keys {
		if !viper.IsSet(key) {
			return fmt.Errorf("config key %s not found in .env file", key)
		}
	}
	return nil
}
