package app

import (
	"nso-server/internal/pkg/config"
	"nso-server/internal/pkg/database"
	"nso-server/internal/pkg/di"
	"nso-server/internal/pkg/logger"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var InfraModule = fx.Options(
	fx.Provide(
		config.LoadConfig,    // *config.Config
		logger.InitZapLogger, // *zap.SugaredLogger
		database.NewDatabase, // *gorm.DB
		di.NewDependencies,   // *di.Dependencies
		func(db *database.Database) *gorm.DB {
			return db.DB // ✅ Provide thêm *gorm.DB
		},
	),
	fx.Invoke(func(l *zap.SugaredLogger) {
		logger.Log = l // gán global
	}),
)
