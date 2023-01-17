package todoservice

import (
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/core/ports"
)

type todoService struct {
	R ports.TodoRepository
}

func (t todoService) ListTodos() ([]*domain.TodoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t todoService) AddTodo(r domain.TodoCreationRequest) (*domain.TodoResponse, error) {
	todo := domain.Todo{Title: r.Title}

	createdTodo, err := t.R.Create(todo)
	if err != nil {
		return nil, err
	}

	return createdTodo.ToResponse(), nil
}

func (t todoService) UpdateTodo(r domain.TodoUpdateRequest) (*domain.TodoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t todoService) PartiallyUpdateTodo(r domain.TodoPartialUpdateRequest) (*domain.TodoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t todoService) DeleteTodo(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewTodoService(r ports.TodoRepository) *todoService {
	return &todoService{R: r}
}
