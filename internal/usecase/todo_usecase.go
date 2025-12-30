package usecase

import (
	"TODO/internal/models"
	"TODO/internal/repository"
	"errors"
	"unicode/utf8"
)

type TodoUsecase struct {
	repo *repository.TodoRepository
}

func NewTodoUsecase(repo *repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{repo: repo}
}

func (u *TodoUsecase) Create(todo *models.Todo) error {
	if utf8.RuneCountInString(todo.Title) == 0 {
		return errors.New("title is required")
	}

	if utf8.RuneCountInString(todo.Title) > 50 {
		return errors.New("title must be 50 characters or less")
	}

	todo.Completed = false
	return u.repo.CreateTodo(todo)
}

func (u *TodoUsecase) GetAll() ([]models.Todo, error) {
	return u.repo.GetAllTodos()
}

func (u *TodoUsecase) GetByID(id uint) (*models.Todo, error) {
	return u.repo.GetTodoByID(id)
}

func (u *TodoUsecase) Update(todo *models.Todo) error {
	if todo.Completed {
		return errors.New("You cannot modify Todo that is already completed")
	}
	return u.repo.UpdateTodo(todo)
}

func (u *TodoUsecase) Delete(id uint) error {
	return u.repo.DeleteTodo(id)
}
