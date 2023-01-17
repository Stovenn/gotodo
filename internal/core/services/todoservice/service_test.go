package todoservice

import (
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/core/ports"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s *todoService

type todoRepositoryMock struct {
	ports.TodoRepository
}

func (todoRepositoryMock) Create(todo domain.Todo) (*domain.Todo, error) {
	return &domain.Todo{
		Id:    "1",
		Title: todo.Title,
		Order: 1,
	}, nil
}

func TestTodoService_AddTodo(t *testing.T) {
	s = NewTodoService(&todoRepositoryMock{})

	arg := domain.TodoCreationRequest{
		Title: "new todo",
	}
	expectedResponse := domain.TodoResponse{ID: "1", Title: "new todo", Order: 1, Completed: false, Url: ""}

	response, err := s.AddTodo(arg)

	assert.NotEmpty(t, response)
	assert.NoError(t, err)

	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.Title, response.Title)
	assert.Equal(t, expectedResponse.Order, response.Order)
	assert.Equal(t, expectedResponse.Completed, response.Completed)
	assert.Equal(t, expectedResponse.Url, response.Url)
}
