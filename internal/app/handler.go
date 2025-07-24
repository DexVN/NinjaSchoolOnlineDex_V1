package app

import (
	"nso-server/internal/net"
	"nso-server/internal/net/handler"
	"nso-server/internal/net/handler/not_login"
	"nso-server/internal/net/handler/not_map"
	"nso-server/internal/net/handler/sub_command"

	"go.uber.org/fx"
)

var HandlerModule = fx.Options(
	fx.Provide(
		not_login.NewRegisterHandler,
		not_login.NewClientInfoHandler,
		not_login.NewRouter,
		not_map.NewRouter,
		sub_command.NewRouter,
		net.NewServer,
		handler.NewRouter,
		func(r *handler.Router) net.RouterFunc {
			return r.Handle
		},
	),
)
