package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/stovenn/gotodo/internal/core/domain"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt returns a random integer between min and max (inclusive)
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString returns a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomEmail returns a random email
func RandomEmail(n int) string {
	return fmt.Sprintf("%s@email.com", RandomString(n))
}

// CreateRandomTodo creates a random Todo
func CreateRandomTodo(order int) *domain.Todo {
	return &domain.Todo{
		ID:        RandomString(15),
		Title:     RandomString(25),
		Order:     order,
		Completed: false,
	}
}

// CreateRandomTodos creates a batch of n random Todo items
func CreateRandomTodos(n int) []*domain.Todo {
	var todos []*domain.Todo
	for i := 0; i < n; i++ {
		todos = append(todos, CreateRandomTodo(i))
	}
	return todos
}

// CreateRandomTodoResponse creates a random TodoResponse
func CreateRandomTodoResponse(order int) *domain.TodoResponse {
	return CreateRandomTodo(order).ToResponse()
}

// CreateRandomTodoResponses creates a batch of n random TodoResponse items
func CreateRandomTodoResponses(n int) []*domain.TodoResponse {
	var todoResponses []*domain.TodoResponse
	for i := 0; i < n; i++ {
		todoResponses = append(todoResponses, CreateRandomTodoResponse(i))
	}
	return todoResponses
}
