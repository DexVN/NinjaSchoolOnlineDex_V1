package main

import (
	"nso-server/internal/app"
	"nso-server/internal/infra"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// ✅ Init logger
	infra.InitLogger("debug", "logs/nso.log")

	infra.Log.Info("🚀 Server starting...")

	if err := app.Bootstrap(); err != nil {
		infra.Log.WithError(err).Fatal("❌ Bootstrap failed")
	}
}
