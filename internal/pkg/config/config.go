package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() (*Config, error) {
	// 🧪 Load .env nếu có
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env file not found, fallback to OS ENV")
	}
	fmt.Println("✅ ENV loaded manually:")
	fmt.Println("➤ APP_ENV =", os.Getenv("APP_ENV"))
	fmt.Println("➤ DB_HOST =", os.Getenv("DB_HOST"))
	fmt.Println("➤ DB_PORT =", os.Getenv("DB_PORT"))
	fmt.Println("➤ DB_USER =", os.Getenv("DB_USER"))
	fmt.Println("➤ DB_NAME =", os.Getenv("DB_NAME"))

	viper.AutomaticEnv()
	// 🛠️ Set default value
	viper.SetDefault("app_env", "development")
	viper.SetDefault("log_level", "debug")
	viper.SetDefault("db_ssl", "disable")
	viper.SetDefault("app_port", "14444")
	viper.SetDefault("server_port", ":14444")
	viper.SetDefault("redis_url", "redis://localhost:6379")
	viper.SetDefault("server_code", "iron")
	viper.SetDefault("default_language", "vi")

	for _, key := range []string{"db_host", "db_port", "db_name", "db_user", "db_password"} {
		viper.BindEnv(key)
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("❌ Failed to unmarshal config: %v", err)
		return nil, err
	}

	// ✅ Kiểm tra các biến bắt buộc
	required := []string{"db_host", "db_port", "db_name", "db_user", "db_password"}
	missing := []string{}
	for _, key := range required {
		if viper.GetString(key) == "" {
			missing = append(missing, key)
		}
	}
	if len(missing) > 0 {
		log.Printf("⚠️  Missing required ENV: %s", strings.Join(missing, ", "))
	}

	// 📋 Log full config (ẩn password nếu muốn)
	log.Println("✅ Loaded configuration:")
	log.Printf("  ➤ APP_ENV: %s", cfg.AppEnv)
	log.Printf("  ➤ DB_HOST: %s", cfg.DbHost)
	log.Printf("  ➤ DB_PORT: %s", cfg.DbPort)
	log.Printf("  ➤ DB_NAME: %s", cfg.DbName)
	log.Printf("  ➤ DB_USER: %s", cfg.DbUser)
	log.Printf("  ➤ REDIS_URL: %s", cfg.RedisUrl)
	log.Printf("  ➤ LANG: %s | LOG_LEVEL: %s | SERVER_PORT: %s", cfg.DefaultLanguage, cfg.LogLevel, cfg.ServerPort)

	return &cfg, nil
}
