package http

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	"TODO/internal/config"
)

func StartServer(
	lc fx.Lifecycle,
	e *echo.Echo,
	cfg config.Config,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				e.Start(":" + cfg.Port)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			return e.Shutdown(shutdownCtx)
		},
	})
}
