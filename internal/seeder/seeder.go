package seeder

import "nso-server/internal/pkg/database"

func Seed(db *database.Database) {
	SeedServer(db)
	SeedSkillOptionTemplate(db)
	SeedNClass(db)
}
