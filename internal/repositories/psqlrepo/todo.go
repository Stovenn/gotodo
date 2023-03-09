package psqlrepo

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/stovenn/gotodo/internal/core/domain"
)

type todoRepository struct {
}

var (
	ErrTodoNotFound = errors.New("could not find Todo")
)

func NewTodoRepository() *todoRepository {
	return &todoRepository{}
}

func (t todoRepository) FindAll() ([]*domain.Todo, error) {
	var todos []*domain.Todo
	err := db.Select(&todos, "SELECT * FROM todos ORDER BY created_at;")
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t todoRepository) FindByID(id string) (*domain.Todo, error) {
	foundTodo := domain.Todo{}
	err := db.Get(&foundTodo, "SELECT * FROM todos WHERE id = $1;", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTodoNotFound
		}
		return nil, err
	}
	return &foundTodo, nil
}

func (t todoRepository) FindByOrder(order int) (*domain.Todo, error) {
	foundTodo := domain.Todo{}
	err := db.Get(&foundTodo, "SELECT * FROM todos WHERE item_order = $1;", order)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTodoNotFound
		}
		return nil, err
	}
	return &foundTodo, nil
}

func (t todoRepository) Create(todo *domain.Todo) (*domain.Todo, error) {
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		log.Fatal(err)
	}
	var maxOrder int
	err = db.Get(&maxOrder, "SELECT count(item_order) FROM todos")
	if err != nil {
		_ = tx.Rollback()
	}
	row := db.QueryRowx("INSERT INTO todos (title, completed, item_order, assigned_to) VALUES ($1, false, $2, $3) RETURNING id;", todo.Title, maxOrder+1, todo.AssignedTo)

	var insertedRowID string
	err = row.Scan(&insertedRowID)
	if err != nil {
		_ = tx.Rollback()
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	return t.FindByID(insertedRowID)
}

func (t todoRepository) Update(todo *domain.Todo) (*domain.Todo, error) {
	row := db.QueryRowx("UPDATE todos SET title = $1, completed = $2, item_order = $3, updated_at = $4 where id = $5 RETURNING *", todo.Title, todo.Completed, todo.Order, time.Now().In(time.UTC), todo.ID)
	err := row.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.Order,
		&todo.AssignedTo,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t todoRepository) DeleteByID(id string) error {
	_, err := db.Exec("DELETE FROM todos where id = $1;", id)
	if err != nil {
		return err
	}
	return nil
}
