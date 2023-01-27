package ports

import "github.com/stovenn/gotodo/internal/core/domain"

type TodoRepository interface {
	FindAll() ([]*domain.Todo, error)
	FindByID(id string) (*domain.Todo, error)
	FindByOrder(order int) (*domain.Todo, error)
	Save(todo *domain.Todo) (*domain.Todo, error)
	DeleteByID(id string) error
}
