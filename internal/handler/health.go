package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c echo.Context) error {
	dbStatus := checkDatabaseConnection()

	status := "ok"
	if dbStatus != "ok" {
		status = "unhealthy"
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   status,
		"database": dbStatus,
	})
}

func checkDatabaseConnection() string {

	return "ok"
}
