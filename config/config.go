package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type APIConfig struct {
	Port      string
	JWTSecret string
	DBURL     string
	WorkEnv   string
}

func LoadConfig() (*APIConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("WARNING: assuming default configuration. .env unreadable: %v", err)
	}

	cfg := &APIConfig{
		Port:      getEnv("PORT"),
		JWTSecret: getEnv("JWT_SECRET"),
		DBURL:     getEnv("DATABASE_URL"),
		WorkEnv:   getEnv("WORKENV"),
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

type fatalError struct {
	s string
}

func (e *fatalError) Error() string {
	return "FATAL: " + e.s
}

func newFatalError(text string) error {
	return &fatalError{text}
}
