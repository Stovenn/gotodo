package psqlrepo

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/pkg/util"
	"github.com/stretchr/testify/assert"
)

func createRandomTodo(t *testing.T) *domain.Todo {
	timestamp := time.Now().UTC()
	arg := &domain.Todo{
		ID:         "",
		Title:      util.RandomString(12),
		Order:      0,
		Completed:  false,
		AssignedTo: sql.NullString{},
	}
	ch := make(chan int, 1)
	go maxOrder(ch)
	expected := &domain.Todo{
		ID:         "",
		Title:      arg.Title,
		Order:      <-ch + 1,
		Completed:  false,
		AssignedTo: sql.NullString{String: "", Valid: false},
		CreatedAt:  timestamp,
		UpdatedAt:  timestamp,
	}

	createdTodo, err := todoRepo.Create(arg)
	assertCreation(t, expected, createdTodo, err)

	return createdTodo
}

func maxOrder(ch chan int) {
	var order int
	db.QueryRowx("SELECT count(item_order) from todos").Scan(&order)
	ch <- order
}

func TestTodoRepository_Create(t *testing.T) {
	t.Run("One insert", func(t *testing.T) {
		createRandomTodo(t)
	})
	t.Run("Concurrent inserts", func(t *testing.T) {

	})
}

func assertCreation(t *testing.T, expected, got *domain.Todo, err error) {
	t.Helper()

	assert.NotEmpty(t, got)
	assert.NoError(t, err)
	assert.Equal(t, expected.Title, got.Title)
	assert.Equal(t, expected.Completed, got.Completed)
	assert.Equal(t, expected.Order, got.Order)
	assert.Equal(t, expected.AssignedTo, got.AssignedTo)
	assert.WithinDuration(t, expected.CreatedAt, got.CreatedAt, time.Second)
	assert.WithinDuration(t, expected.UpdatedAt, got.UpdatedAt, time.Second)
	assert.NotZero(t, got.ID)
}

func TestTodoRepository_Update(t *testing.T) {
	timestamp := time.Now().UTC()
	todo := createRandomTodo(t)
	arg := &domain.Todo{
		ID:         todo.ID,
		Title:      "updated title",
		Order:      todo.Order,
		Completed:  true,
		AssignedTo: todo.AssignedTo,
	}
	expected := &domain.Todo{
		ID:         arg.ID,
		Title:      arg.Title,
		Order:      arg.Order,
		Completed:  arg.Completed,
		AssignedTo: arg.AssignedTo,
		CreatedAt:  todo.CreatedAt,
		UpdatedAt:  timestamp,
	}

	updatedTodo, err := todoRepo.Update(arg)
	assertUpdate(t, expected, updatedTodo, err)
}

func assertUpdate(t *testing.T, expected, got *domain.Todo, err error) {
	t.Helper()

	assert.NotEmpty(t, got)
	assert.NoError(t, err)
	assert.Equal(t, expected.Title, got.Title)
	assert.Equal(t, expected.Completed, got.Completed)
	assert.Equal(t, expected.Order, got.Order)
	assert.Equal(t, expected.AssignedTo, got.AssignedTo)
	assert.WithinDuration(t, expected.CreatedAt, got.CreatedAt, time.Second)
	assert.WithinDuration(t, expected.UpdatedAt, got.UpdatedAt, time.Second)
}

func TestTodoRepository_FindByID(t *testing.T) {
	todo := createRandomTodo(t)

	foundTodo, err := todoRepo.FindByID(todo.ID)
	assert.NotEmpty(t, foundTodo)
	assert.NoError(t, err)
	assert.Equal(t, todo.ID, foundTodo.ID)
}

func TestTodoRepository_FindByOrder(t *testing.T) {
	todo := createRandomTodo(t)

	foundTodo, err := todoRepo.FindByOrder(todo.Order)
	assert.NotEmpty(t, foundTodo)
	assert.NoError(t, err)
	assert.Equal(t, todo.Order, foundTodo.Order)
}
func TestTodoRepository_FindAll(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomTodo(t)
	}

	todos, err := todoRepo.FindAll()
	assert.NotEmpty(t, todos)
	assert.NoError(t, err)
}

func TestTodoRepository_DeleteByID(t *testing.T) {
	todo := createRandomTodo(t)

	err := todoRepo.DeleteByID(todo.ID)
	assert.NoError(t, err)

	foundTodo, err := todoRepo.FindByID(todo.ID)
	assert.Empty(t, foundTodo)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrTodoNotFound)
}
