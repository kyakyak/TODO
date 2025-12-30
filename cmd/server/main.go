package main

import (
	"go.uber.org/fx"

	_ "TODO/docs"
	"TODO/internal/config"
	"TODO/internal/handler"
	"TODO/internal/infrastructure/http"
	"TODO/internal/repository"
	"TODO/internal/usecase"
)

// @title TODO API
// @version 1.0
// @description Simple TODO API
// @host localhost:8080
// @BasePath /
func main() {
	fx.New(
		fx.Provide(
			config.NewConfig,
			config.NewDBConnection,
			repository.NewTodoRepository,
			usecase.NewTodoUsecase,
			// handler.NewHealthHandler,
			// handler.NewTodoHandler,
		),
		http.Module,
		handler.Module,
	).Run()
}
