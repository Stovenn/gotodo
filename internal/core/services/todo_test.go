package services

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/core/ports"
	mockdb "github.com/stovenn/gotodo/internal/repositories/mock"
	"github.com/stovenn/gotodo/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s ports.TodoService

func TestTodoService_CreateTodo(t *testing.T) {
	arg := domain.TodoCreationRequest{
		Title: "new todo",
	}
	todo := &domain.Todo{Title: arg.Title, Completed: false, Order: 1}
	expectedResponse := todo.ToResponse()

	ctrl := gomock.NewController(t)
	repository := mockdb.NewMockTodoRepository(ctrl)
	s = NewTodoService(repository)

	repository.EXPECT().
		Create(&domain.Todo{Title: arg.Title}).
		Times(1).
		Return(todo, nil)

	response, err := s.CreateTodo(arg)

	assert.NotEmpty(t, response)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}

func TestTodoService_DisplayAllTodos(t *testing.T) {
	todos := util.CreateRandomTodos(3)

	ctrl := gomock.NewController(t)
	repository := mockdb.NewMockTodoRepository(ctrl)
	s = NewTodoService(repository)

	repository.EXPECT().
		FindAll().
		Times(1).
		Return(todos, nil)

	response, err := s.DisplayAllTodos()

	assert.NotEmpty(t, response)
	assert.NoError(t, err)
	assert.Equal(t, len(todos), len(response))
}

func TestTodoService_DisplayTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := mockdb.NewMockTodoRepository(ctrl)
	s = NewTodoService(repository)

	t.Run("given a todo id should return a todo response", func(t *testing.T) {
		todo := util.CreateRandomTodo(1)
		expected := todo.ToResponse()

		repository.EXPECT().
			FindByID(todo.ID).
			Times(1).
			Return(todo, nil)

		response, err := s.DisplayTodo(todo.ID)

		assert.NotEmpty(t, response)
		assert.NoError(t, err)
		assert.Equal(t, expected, response)
	})

	t.Run("given an unknown todo id should return an error", func(t *testing.T) {
		id := "unknown"

		repository.EXPECT().
			FindByID(id).
			Times(1).
			Return(nil, fmt.Errorf("todo not found"))

		response, err := s.DisplayTodo(id)

		assert.Empty(t, response)
		assert.Error(t, err)
		assert.EqualError(t, err, "todoservice.DisplayTodo: todo not found")
	})
}

func TestTodoService_UpdateTodo(t *testing.T) {
	todo := util.CreateRandomTodo(1)
	updateRequest := domain.TodoUpdateRequest{Title: "updated title", Completed: true, Order: 2}
	todoUpdate := domain.Todo{Title: updateRequest.Title, Completed: updateRequest.Completed, Order: updateRequest.Order}

	updatedTodo := &domain.Todo{ID: todo.ID, Title: todoUpdate.Title, Completed: todoUpdate.Completed, Order: todoUpdate.Order}

	expected := updatedTodo.ToResponse()

	ctrl := gomock.NewController(t)
	repository := mockdb.NewMockTodoRepository(ctrl)
	s = NewTodoService(repository)

	findByIDCall := repository.EXPECT().FindByID(todo.ID).Times(1).Return(todo, nil)
	findByOrderCall := repository.EXPECT().FindByOrder(updateRequest.Order).Times(1).Return(nil, nil)
	saveCall := repository.EXPECT().Update(todo.ID, updatedTodo).Times(1).Return(updatedTodo, nil)
	gomock.InOrder(findByIDCall, findByOrderCall, saveCall)

	response, err := s.UpdateTodo(todo.ID, updateRequest)

	assert.NotEmpty(t, response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestTodoService_PartiallyUpdateTodo(t *testing.T) {
	todo := util.CreateRandomTodo(1)
	updateRequest := domain.TodoPartialUpdateRequest{Order: 2}
	todoUpdate := domain.Todo{Title: updateRequest.Title, Completed: updateRequest.Completed, Order: updateRequest.Order}

	updatedTodo := &domain.Todo{ID: todo.ID, Title: todo.Title, Completed: todo.Completed, Order: todoUpdate.Order}

	expected := updatedTodo.ToResponse()

	ctrl := gomock.NewController(t)
	repository := mockdb.NewMockTodoRepository(ctrl)
	s = NewTodoService(repository)

	findByIDCall := repository.EXPECT().FindByID(todo.ID).Times(1).Return(todo, nil)
	findByOrderCall := repository.EXPECT().FindByOrder(updateRequest.Order).Times(1).Return(nil, nil)
	saveCall := repository.EXPECT().Update(todo.ID, updatedTodo).Times(1).Return(updatedTodo, nil)
	gomock.InOrder(findByIDCall, findByOrderCall, saveCall)

	response, err := s.PartiallyUpdateTodo(todo.ID, updateRequest)

	assert.NotEmpty(t, response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestTodoService_DeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := mockdb.NewMockTodoRepository(ctrl)
	s = NewTodoService(repository)

	t.Run("given a todo id should not return an error", func(t *testing.T) {
		id := "1"

		repository.EXPECT().
			DeleteByID("1").
			Times(1).
			Return(nil)

		err := s.DeleteTodo(id)

		assert.NoError(t, err)
	})
	t.Run("given an unknown todo id should return an error", func(t *testing.T) {
		id := "unknown"

		repository.EXPECT().
			DeleteByID(id).
			Times(1).
			Return(fmt.Errorf("todo not found"))

		err := s.DeleteTodo(id)

		assert.Error(t, err)
		assert.EqualError(t, err, "todoservice.DeleteOneTodo: todo not found")
	})
}
