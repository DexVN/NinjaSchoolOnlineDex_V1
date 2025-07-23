package app

import (
	"fmt"

	_ "github.com/lib/pq"
	"nso-server/internal/config"
	"nso-server/internal/infra"
	"nso-server/internal/lang"
	"nso-server/internal/model"
	netio "nso-server/internal/net"
	"nso-server/internal/net/handler"
)

func Bootstrap() error {
	// ✅ Load config
	config.LoadEnv()
	config.LoadConfig()

	infra.Log.Infof("🌐 Loaded config: Env=%s, LogLevel=%s", config.Config.AppEnv, config.Config.LogLevel)

	// ✅ Init language
	defaultLang := config.Config.DefaultLanguage
	if err := lang.Init(defaultLang); err != nil {
		infra.Log.WithError(err).Error("❌ Failed to load language")
		return fmt.Errorf("load language failed: %w", err)
	}
	infra.Log.Infof("🌍 Language loaded: %s", lang.GetLangDisplayName(defaultLang))

	// ✅ Init database
	infra.InitDatabase()
	infra.Log.Info("🔗 Database initialized")

	// ✅ Auto migrate
	if err := autoMigrateModels(); err != nil {
		infra.Log.WithError(err).Error("❌ Auto migration failed")
		return fmt.Errorf("auto-migrate failed: %w", err)
	}
	infra.Log.Info("📦 Models auto-migrated")

	// ✅ Seed data in development
	if config.Config.AppEnv == "development" {
		infra.Log.Warn("🌱 Development mode: running seed")
		Seed()
	}

	// ✅ Start game server
	srv, err := netio.NewServer(":14444", handler.RouteMessage)
	if err != nil {
		infra.Log.WithError(err).Error("❌ Server startup failed")
		return fmt.Errorf("server error: %w", err)
	}

	infra.Log.Info("🚀 NSO Server started on :14444")
	srv.Start()

	return nil
}

func autoMigrateModels() error {
	db := infra.DB
	infra.Log.Info("📦 Starting auto-migration...")

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
