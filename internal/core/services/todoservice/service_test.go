package todoservice

import (
	"github.com/golang/mock/gomock"
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/core/ports"
	mockdb "github.com/stovenn/gotodo/internal/repositories/mock"
	"github.com/stovenn/gotodo/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s ports.TodoService

func TestTodoService_AddTodo(t *testing.T) {
	arg := domain.TodoCreationRequest{
		Title: "new todo",
	}
	todo := &domain.Todo{Title: arg.Title, Completed: false, Order: 1, Url: ""}
	expectedResponse := todo.ToResponse()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockdb.NewMockTodoRepository(ctrl)
	s = NewTodoService(repository)

	repository.EXPECT().
		Create(domain.Todo{Title: arg.Title}).
		Times(1).
		Return(todo, nil)

	response, err := s.AddTodo(arg)

	assert.NotEmpty(t, response)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}

func TestTodoService_ListTodos(t *testing.T) {
	todos := util.CreateRandomTodos(3)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockdb.NewMockTodoRepository(ctrl)
	s = NewTodoService(repository)

	repository.EXPECT().
		FindAll().
		Times(1).
		Return(todos, nil)

	response, err := s.ListTodos()

	assert.NotEmpty(t, response)
	assert.NoError(t, err)
	assert.Equal(t, len(todos), len(response))
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

func TestTodoService_UpdateTodo(t *testing.T) {

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
