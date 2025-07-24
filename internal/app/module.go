package app

import (
	"nso-server/internal/lang"

	"go.uber.org/fx"
)

var Module = fx.Options(
	InfraModule,   // 📦 config, logger, DB, Redis
	HandlerModule, // 📡 Message router / handler
	ServiceModule, // 🧠 Game logic layer
	lang.Module,   //
	fx.Invoke(

		MigrateModels, // 👷 DB migrate
		SeedData,      // 🌱 Seed trong dev
		StartServer,   // 🚀 Start TCP
	),
)
