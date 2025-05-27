package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mambo-dev/kenya-locations/internal/database"
)

type APIConfig struct {
	Port       string
	TAuthToken string
	DBURL      string
	WorkEnv    string
	Db         *database.Queries
}

func LoadConfig() (*APIConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("WARNING: assuming default configuration. .env unreadable: %v", err)
	}

	cfg := &APIConfig{
		Port:       getEnv("PORT"),
		DBURL:      getEnv("DATABASE_URL"),
		WorkEnv:    getEnv("WORK_ENV"),
		TAuthToken: getEnv("TAUTH_TOKEN"),
	}

	return cfg, nil
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {

		log.Fatalf("FATAL: %s environment variable is not set", key)
	}
	return value
}
