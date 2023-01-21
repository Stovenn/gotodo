package todoservice

import (
	"fmt"
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/core/ports"
)

var db = []*domain.Todo{
	{ID: "1", Title: "todo 1", Order: 1, Completed: false, Url: ""},
	{ID: "2", Title: "todo 2", Order: 2, Completed: false, Url: ""},
	{ID: "3", Title: "todo 3", Order: 3, Completed: false, Url: ""},
}

type todoRepositoryMock struct {
	ports.TodoRepository
}

func (r *todoRepositoryMock) Create(todo domain.Todo) (*domain.Todo, error) {
	return &domain.Todo{ID: "1", Title: todo.Title, Order: 1, Completed: false, Url: ""}, nil
}

func (r *todoRepositoryMock) FindAll() ([]*domain.Todo, error) {
	return db, nil
}

func (r *todoRepositoryMock) FindByID(id string) (*domain.Todo, error) {
	switch id {
	case "unknown":
		return nil, fmt.Errorf("todo not found")
	default:
		return &domain.Todo{ID: id, Title: "todo 1", Order: 1, Completed: false, Url: ""}, nil
	}
}

func (r *todoRepositoryMock) DeleteByID(id string) error {
	switch id {
	case "unknown":
		return fmt.Errorf("todo not found")
	default:
		return nil
	}
}
