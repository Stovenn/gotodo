package psqlrepo

import (
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/stovenn/gotodo/internal/core/domain"
	"log"

	_ "github.com/lib/pq"
)

type todoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository() *todoRepository {
	conn, err := sqlx.Connect(viper.GetString("DB_DRIVER"), viper.GetString("DB_URL"))
	if err != nil {
		log.Panicln(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Panicln(err)
	}

	return &todoRepository{db: conn}
}

func (t todoRepository) FindAll() ([]*domain.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t todoRepository) FindByID(id string) (*domain.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t todoRepository) FindByOrder(order int) (*domain.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t todoRepository) Save(todo *domain.Todo) (*domain.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t todoRepository) DeleteByID(id string) error {
	//TODO implement me
	panic("implement me")
}
