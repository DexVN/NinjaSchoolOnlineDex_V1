package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env file not found, using system env")
	}
}

func LoadConfig() {
	Config = AppConfig{
		AppEnv:   getEnv("APP_ENV", "development"),
		DBUrl:    getEnv("DB_URL", "postgres://postgres:123456@localhost:5432/nso_db?sslmode=disable"),
		Port:     getEnv("APP_PORT", "14444"),
		RedisUrl: getEnv("REDIS_URL", "redis://localhost:6379"),
	}
}

func getEnv(key string, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
