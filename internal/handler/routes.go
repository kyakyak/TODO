package handler

import "github.com/labstack/echo/v4"

func RegisterRoutes(
	e *echo.Echo,
	healthHandler *HealthHandler,
) {
	e.GET("/health", healthHandler.Health)
}
