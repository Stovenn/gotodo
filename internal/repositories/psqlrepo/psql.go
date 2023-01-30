package psqlrepo

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func OpenDB(driver, url string) error {
	conn, err := sqlx.Connect(driver, url)
	if err != nil {
		return err
	}

	err = conn.Ping()
	if err != nil {
		return err
	}

	db = conn
	return nil
}
