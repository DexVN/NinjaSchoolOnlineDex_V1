package app

import (
	"nso-server/internal/model"
	"nso-server/internal/pkg/database"
)

func MigrateModels(db *database.Database) {
	db.AutoMigrate(
		&model.Account{},
		&model.Character{},
		&model.NClass{},
		&model.SkillTemplate{},
		&model.Skill{},
		&model.SkillOption{},
		&model.SkillOptionTemplate{},
		&model.DataVersion{},
		&model.ClientSession{},
	)
}
