package psqlrepo

import (
	"github.com/stovenn/gotodo/pkg/util"
	"log"
	"testing"
)

var todoRepo *todoRepository
var userRepo *userRepository

func TestMain(m *testing.M) {
	config, err := util.SetupConfig("../../..")
	if err != nil {
		log.Fatalf("cannot load config: %v\n", err)
	}
	err = OpenDB(config.DBDriver, config.DBUrl)
	if err != nil {
		log.Fatalf("cannot connect to DB: %v\n", err)
	}
	todoRepo = NewTodoRepository()
	userRepo = NewUserRepository()
	m.Run()
}
