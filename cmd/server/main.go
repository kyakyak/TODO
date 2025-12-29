package main

import (
	"go.uber.org/fx"

	"TODO/internal/config"
	"TODO/internal/handler"
	httpinfra "TODO/internal/infrastructure/http"
)

func main() {
	fx.New(
		fx.Provide(
			config.NewConfig,
		),
		httpinfra.Module,
		handler.Module,
	).Run()
}
