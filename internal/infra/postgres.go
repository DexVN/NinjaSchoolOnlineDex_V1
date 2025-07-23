package infra

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	ssl := os.Getenv("DB_SSL")
	if ssl == "" {
		ssl = "disable"
	}

	// Tạo DSN từ từng biến
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, pass, host, port, name, ssl)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		Log.Fatalf("❌ Cannot connect to database: %v", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
	Log.Info("✅ Connected to PostgreSQL")
}
