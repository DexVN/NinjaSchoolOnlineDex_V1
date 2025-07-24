package lang

import (
	"go.uber.org/fx"
	"nso-server/internal/pkg/config"
)

var Module = fx.Options(
	fx.Provide(NewI18n), // <-- Inject I18n
	fx.Invoke(func(cfg *config.Config) {
		if err := Init(cfg.DefaultLanguage); err != nil {
			panic("❌ Cannot init language: " + err.Error())
		}
	}),
)
