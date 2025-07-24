package di

import (
	"nso-server/internal/lang"
	"nso-server/internal/pkg/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Dependencies struct {
	Config *config.Config
	DB     *gorm.DB
	Log    *zap.SugaredLogger
	I18n   lang.I18n
}

func NewDependencies(
	cfg *config.Config,
	db *gorm.DB,
	log *zap.SugaredLogger,
	i18n lang.I18n,
) *Dependencies {
	return &Dependencies{
		Config: cfg,
		DB:     db,
		Log:    log,
		I18n:   i18n,
	}
}
