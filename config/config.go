package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config() string {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, proceeding with defaults or env vars")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatalf("Failed to load the DATABASE_URL")
	}

	return dbURL
}
