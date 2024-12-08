package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Warning: .env file not found. Using default values")
	}

	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "3700")
	}

	if os.Getenv("JWT_SECRET") == "" {
		os.Setenv("JWT_SECRET", "your-secret-key")
	}

	if os.Getenv("REFRESH_SECRET") == "" {
		os.Setenv("REFRESH_SECRET", "your-refresh-secret-key")
	}
}
