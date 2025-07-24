package database

import (
	"fmt"
	"time"

	"nso-server/internal/pkg/config"
	"nso-server/internal/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(cfg *config.Config) (*Database, error) {
	logger.Infof("🔗 Connecting to database: %s", cfg.DbName)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbSSL)

	db, err := gorm.Open(postgres.Open(dsn), NewGormConfig())
	if err != nil {
		logger.WithError(err).Error("❌ Cannot connect to database")
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	logger.Info("✅ Database connection established")
	return &Database{DB: db}, nil
}
