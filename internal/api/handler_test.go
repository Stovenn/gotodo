package api

import (
	"github.com/golang/mock/gomock"
	mockservice "github.com/stovenn/gotodo/internal/core/services/todoservice/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTodoHandler_HandleCreateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockservice.NewMockTodoService(ctrl)

	service.EXPECT().
		AddTodo(gomock.Any()).
		Times(1).
		Return(nil, nil)

	server := NewServer(NewHandler(service))

	recorder := httptest.NewRecorder()
	body := strings.NewReader(`{"title": "new todo"}`)
	request, err := http.NewRequest(http.MethodPost, "/api/todos/", body)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
}

func TestTodoHandler_HandleListTodo(t *testing.T) {
	//t.Error("implement me")
}

func TestTodoHandler_HandleFindTodoByID(t *testing.T) {
	//t.Error("implement me")
}
