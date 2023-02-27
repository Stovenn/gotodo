package ports

import (
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/dto/request"
	"github.com/stovenn/gotodo/internal/dto/response"
	"github.com/stovenn/gotodo/pkg/token"
	"github.com/stovenn/gotodo/pkg/util"
)

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
	SignUp(r request.UserCreationRequest) (*response.UserResponse, error)
	Login(uc request.UserCredentials, m token.Maker, c util.Config) (*response.LoginResponse, error)
}
