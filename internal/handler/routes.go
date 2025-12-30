package handler

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRoutes(
	e *echo.Echo,
	healthHandler *HealthHandler,
	todoHandler *TodoHandler,
) {
	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Health
	e.GET("/health", healthHandler.Health)

	// Todos
	e.POST("/todos", todoHandler.CreateTodo)
	e.GET("/todos", todoHandler.GetTodos)
	e.GET("/todos/:id", todoHandler.GetTodoByID)
	e.PUT("/todos/:id", todoHandler.UpdateTodo)
	e.DELETE("/todos/:id", todoHandler.DeleteTodo)
}
