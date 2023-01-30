package psqlrepo

import (
	"database/sql"
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createRandomTodo(t *testing.T) *domain.Todo {
	arg := &domain.Todo{
		ID:        "",
		Title:     util.RandomString(12),
		Order:     0,
		Completed: false,
	}
	expected := &domain.Todo{
		ID:        "",
		Title:     arg.Title,
		Order:     maxOrder() + 1,
		Completed: false,
	}

	createdTodo, err := r.Create(arg)
	assertCreation(t, expected, createdTodo, err)

	return createdTodo
}

func maxOrder() int {
	var order int
	r.db.QueryRowx("SELECT MAX(item_order) from todos").Scan(&order)
	return order
}

func TestTodoRepository_Create(t *testing.T) {
	createRandomTodo(t)
}

func assertCreation(t *testing.T, expected, got *domain.Todo, err error) {
	t.Helper()

	assert.NotEmpty(t, got)
	assert.NoError(t, err)
	assert.Equal(t, expected.Title, got.Title)
	assert.Equal(t, expected.Order, got.Order)
	assert.NotZero(t, got.ID)
}

func TestTodoRepository_Update(t *testing.T) {
	todo := createRandomTodo(t)
	arg := &domain.Todo{
		ID:        todo.ID,
		Title:     "updated title",
		Order:     todo.Order,
		Completed: todo.Completed,
	}
	expected := &domain.Todo{
		ID:        arg.ID,
		Title:     arg.Title,
		Order:     arg.Order,
		Completed: arg.Completed,
	}

	updatedTodo, err := r.Update(todo.ID, arg)
	assertUpdate(t, expected, updatedTodo, err)
}

func assertUpdate(t *testing.T, expected, got *domain.Todo, err error) {
	t.Helper()

	assert.NotEmpty(t, got)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestTodoRepository_FindByID(t *testing.T) {
	todo := createRandomTodo(t)

	foundTodo, err := r.FindByID(todo.ID)
	assert.NotEmpty(t, foundTodo)
	assert.NoError(t, err)
	assert.Equal(t, todo, foundTodo)
}

func TestTodoRepository_FindByOrder(t *testing.T) {
	todo := createRandomTodo(t)

	foundTodo, err := r.FindByOrder(todo.Order)
	assert.NotEmpty(t, foundTodo)
	assert.NoError(t, err)
	assert.Equal(t, todo, foundTodo)
}
func TestTodoRepository_FindAll(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomTodo(t)
	}

	todos, err := r.FindAll()
	assert.NotEmpty(t, todos)
	assert.NoError(t, err)
}

func TestTodoRepository_DeleteByID(t *testing.T) {
	todo := createRandomTodo(t)

	err := r.DeleteByID(todo.ID)
	assert.NoError(t, err)

	foundTodo, err := r.FindByID(todo.ID)
	assert.Empty(t, foundTodo)
	assert.Error(t, err)
	assert.ErrorIs(t, err, sql.ErrNoRows)
}
