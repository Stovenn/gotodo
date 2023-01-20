package todorepo

import (
	"database/sql"
	"github.com/stovenn/gotodo/internal/core/domain"
	"log"

	_ "github.com/lib/pq"
)

const (
	driver = "postgres"
	url    = ""
)

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository() *todoRepository {
	conn, err := sql.Open(driver, url)
	if err != nil {
		log.Panicln(err)
	}
	conn.Ping()
	if err != nil {
		log.Panicln(err)
	}
	return &todoRepository{db: conn}
}

func (t *todoRepository) FindAll() ([]*domain.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t *todoRepository) FindByID(id string) (*domain.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t *todoRepository) Create(todo domain.Todo) (*domain.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t *todoRepository) Update(id string, todo domain.Todo) (*domain.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t *todoRepository) DeleteByID(id string) error {
	//TODO implement me
	panic("implement me")
}
