package handler

import (
	"TODO/internal/models"
	"TODO/internal/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	Usecase *usecase.TodoUsecase
}

func NewTodoHandler(uc *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{Usecase: uc}
}

// CreateTodo godoc
// @Summary      Create todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo body models.Todo true "Todo"
// @Success      201 {object} models.Todo
// @Router       /todos [post]
func (h *TodoHandler) CreateTodo(c echo.Context) error {
	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	if err := h.Usecase.Create(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, todo)
}

// GetTodos godoc
// @Summary      Get todos
// @Description  Get all todos
// @Tags         todos
// @Produce      json
// @Success      200 {array} models.Todo
// @Router       /todos [get]
func (h *TodoHandler) GetTodos(c echo.Context) error {
	todos, err := h.Usecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, todos)
}

// GetTodoByID godoc
// @Summary      Get todo by ID
// @Tags         todos
// @Produce      json
// @Param        id path int true "Todo ID"
// @Success      200 {object} models.Todo
// @Router       /todos/{id} [get]
func (h *TodoHandler) GetTodoByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	todo, err := h.Usecase.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, todo)
}

// UpdateTodo godoc
// @Summary      Update todo
// @Description  Update todo by ID
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      int           true  "Todo ID"
// @Param        todo body      models.Todo   true  "Update Todo"
// @Success      200  {object}  models.Todo
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /todos/{id} [put]
func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	if err := h.Usecase.Update(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, todo)
}

// DeleteTodo godoc
// @Summary      Delete todo
// @Description  Delete todo by ID
// @Tags         todos
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      204
// @Failure      404  {object}  map[string]string
// @Router       /todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	if err := h.Usecase.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
