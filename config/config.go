package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {

	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env" // Default to .env if ENV_FILE is not set
	}
	// load .env file
	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
