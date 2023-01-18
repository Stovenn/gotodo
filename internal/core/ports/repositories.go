package ports

import "github.com/stovenn/gotodo/internal/core/domain"

type TodoRepository interface {
	FindAll() ([]*domain.Todo, error)
	FindByID(id string) (*domain.Todo, error)
	Create(todo domain.Todo) (*domain.Todo, error)
	Update(todo domain.Todo) (*domain.Todo, error)
	DeleteByID(todo domain.Todo) (*domain.Todo, error)
}
