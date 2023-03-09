package api

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/stovenn/gotodo/internal/core/ports"
	"github.com/stovenn/gotodo/pkg/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, todoService ports.TodoService, userService ports.UserService) *Server {
	config := util.Config{
		SymmetricKey:  util.RandomString(32),
		TokenDuration: time.Minute,
	}
	server, err := NewServer(config, todoService, userService, log.Default(), log.Default())
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
