package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUri    string
	Port     string
	JWTSecret string
}

// LoadConfig loads configuration from .env file or environment variables.
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading environment variables directly.")
	}

	config := &Config{
		DBUri:     getEnv("DB_URI", "mongodb://localhost:27017/todoapp"),
		Port:      getEnv("PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "supersecret"),
	}

	return config, nil
}

// Helper function to retrieve environment variables or default values.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

