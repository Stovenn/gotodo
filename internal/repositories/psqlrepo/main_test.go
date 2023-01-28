package psqlrepo

import (
	"github.com/stovenn/gotodo/pkg/util"
	"log"
	"testing"
)

var r *todoRepository

func TestMain(m *testing.M) {
	config, err := util.SetupConfig("../../..")
	if err != nil {
		log.Fatalf("cannot load config: %v\n", err)
	}

	r = NewTodoRepository(config.DBDriver, config.DBURL)

	m.Run()
}
