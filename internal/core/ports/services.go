package ports

import "github.com/stovenn/gotodo/internal/core/domain"

type TodoService interface {
	CreateTodo(r domain.TodoCreationRequest) (*domain.TodoResponse, error)
	DisplayTodo(id string) (*domain.TodoResponse, error)
	DisplayAllTodos() ([]*domain.TodoResponse, error)
	UpdateTodo(id string, r domain.TodoUpdateRequest) (*domain.TodoResponse, error)
	PartiallyUpdateTodo(id string, r domain.TodoPartialUpdateRequest) (*domain.TodoResponse, error)
	DeleteTodo(id string) error
	DeleteAllTodos() error
}

type UserService interface {
	SignUp(r domain.UserCreationRequest) (*domain.UserResponse, error)
	Login(uc domain.UserCredentials) error
}
