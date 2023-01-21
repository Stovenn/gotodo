package todoservice

import (
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/core/ports"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s ports.TodoService

func init() {
	s = NewTodoService(&todoRepositoryMock{})
}

func TestTodoService_AddTodo(t *testing.T) {
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

func TestTodoService_ListTodos(t *testing.T) {
	expectedResponse := []*domain.TodoResponse{
		{ID: "1", Title: "todo 1", Order: 1, Completed: false, Url: ""},
		{ID: "2", Title: "todo 2", Order: 2, Completed: false, Url: ""},
		{ID: "3", Title: "todo 3", Order: 3, Completed: false, Url: ""},
	}

	response, err := s.ListTodos()

	assert.NotEmpty(t, response)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}

func TestTodoService_FindTodoByID(t *testing.T) {
	t.Run("given a todo id should return a todo response", func(t *testing.T) {
		id := "1"
		response, err := s.FindTodoByID(id)
		assert.NotEmpty(t, response)
		assert.NoError(t, err)
		assert.Equal(t, id, response.ID)
	})

	t.Run("given an unknown todo id should return an error", func(t *testing.T) {
		id := "unknown"
		response, err := s.FindTodoByID(id)
		assert.Empty(t, response)
		assert.Error(t, err)
		assert.EqualError(t, err, "todoservice.FindTodoByID: todo not found")
	})
}
func TestTodoService_DeleteTodo(t *testing.T) {
	t.Run("given a todo id should not return an error", func(t *testing.T) {
		err := s.DeleteTodo("id")

		assert.NoError(t, err)
	})
	t.Run("given an unknown todo id should return an error", func(t *testing.T) {
		err := s.DeleteTodo("unknown")

		assert.Error(t, err)
		assert.EqualError(t, err, "todoservice.DeleteTodo: todo not found")
	})
}
