package app

import (
	"nso-server/internal/pkg/config"
	"nso-server/internal/pkg/database"
	"nso-server/internal/seeder"
)

func SeedData(cfg *config.Config, db *database.Database) {
	if cfg.AppEnv == "development" {
		seeder.Seed(db)
	}
}
