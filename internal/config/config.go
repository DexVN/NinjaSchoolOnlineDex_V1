package config

import (

	"os"

	"nso-server/internal/infra"
	"github.com/joho/godotenv"
)

var log = infra.Log.WithField("module", "config")

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env file not found, using system env")
	}
	if os.Getenv("APP_ENV") == "" {
		log.Println("⚠️  APP_ENV not set, defaulting to 'development'")
		os.Setenv("APP_ENV", "development")
	}
}

func LoadConfig() {
	Config = AppConfig{
		AppEnv:   getEnv("APP_ENV", "development"),
		DBUrl:    getEnv("DB_URL", "postgres://postgres:123456@localhost:5432/nso_db?sslmode=disable"),
		Port:     getEnv("APP_PORT", "14444"),
		RedisUrl: getEnv("REDIS_URL", "redis://localhost:6379"),
		DefaultLanguage: getEnv("DEFAULT_LANGUAGE", "vi"),
		LogLevel: getEnv("LOG_LEVEL", "debug"),
	}
}

func getEnv(key string, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
