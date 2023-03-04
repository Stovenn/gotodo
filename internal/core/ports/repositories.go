package ports

import "github.com/stovenn/gotodo/internal/core/domain"

// TodoRepository defines a set of methods to interact with todos in a database
type TodoRepository interface {
	FindAll() ([]*domain.Todo, error)
	FindByID(id string) (*domain.Todo, error)
	FindByOrder(order int) (*domain.Todo, error)
	Create(todo *domain.Todo) (*domain.Todo, error)
	Update(todo *domain.Todo) (*domain.Todo, error)
	DeleteByID(id string) error
}

// UserRespository defines a set of methods to interact with users in a database
type UserRepository interface {
	FindAll() ([]*domain.User, error)
	FindByID(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	DeleteByID(id string) error
}
