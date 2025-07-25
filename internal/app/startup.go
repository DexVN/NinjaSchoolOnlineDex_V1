package app

import (
	"context"
	"nso-server/internal/net"

	"go.uber.org/fx"
)

func StartServer(lc fx.Lifecycle, srv *net.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go srv.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			done := make(chan struct{})
			go func() {
				srv.Stop()
				close(done)
			}()

			select {
			case <-done:
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		},
	})
}
