package app

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"nso-server/internal/config"
	"nso-server/internal/infra"
	"nso-server/internal/model"
	netio "nso-server/internal/net"
	"nso-server/internal/net/handler"
	
)

func Bootstrap() error {
	config.LoadEnv()
	config.LoadConfig()

	log.Println("🌐 Env:", config.Config.AppEnv)
	log.Println("🔗 DB:", config.Config.DBUrl)

	infra.InitDatabase()

	// Chạy AutoMigrate
	if err := autoMigrateModels(); err != nil {
		return fmt.Errorf("auto-migrate failed: %w", err)
	}

	// Chạy Seed
	if config.Config.AppEnv == "development" {
		Seed()
	}

	// Start server
	srv, err := netio.NewServer(":14444", handler.RouteMessage)
	if err != nil {
		return fmt.Errorf("server error: %w", err)
	}
	log.Println("🚀 NSO Server up")
	srv.Start()

	return nil
}

func autoMigrateModels() error {
	db := infra.DB

	log.Println("📦 Auto migrating models...")

	return db.AutoMigrate(
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