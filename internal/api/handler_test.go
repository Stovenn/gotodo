package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stovenn/gotodo/internal/core/domain"
	mockservice "github.com/stovenn/gotodo/internal/core/services/todoservice/mock"
	"github.com/stovenn/gotodo/pkg/util"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTodoHandler_HandleCreateTodo(t *testing.T) {
	todoResponse := util.CreateRandomTodoResponse(1)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockservice.NewMockTodoService(ctrl)
	service.EXPECT().
		AddTodo(gomock.Any()).
		Times(1).
		Return(todoResponse, nil)

	server := NewServer(NewHandler(service))
	body := strings.NewReader(`{"title": "new todo"}`)
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/api/todos/", body)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusCreated, recorder.Code)
	requireBodyMatchTodoResponse(t, recorder.Body, todoResponse)
}

func TestTodoHandler_HandleListTodo(t *testing.T) {
	todoResponses := util.CreateRandomTodoResponses(4)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockservice.NewMockTodoService(ctrl)
	service.EXPECT().
		ListTodos().
		Times(1).
		Return(todoResponses, nil)

	server := NewServer(NewHandler(service))
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/api/todos/", nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchTodoResponses(t, recorder.Body, todoResponses)
}

func TestTodoHandler_HandleFindTodoByID(t *testing.T) {
	todoResponse := util.CreateRandomTodoResponse(1)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockservice.NewMockTodoService(ctrl)
	// build stubs
	service.EXPECT().
		FindTodoByID(todoResponse.ID).
		Times(1).
		Return(todoResponse, nil)

	// start test server and send request
	server := NewServer(NewHandler(service))
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/api/todos/%s", todoResponse.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	// check response
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchTodoResponse(t, recorder.Body, todoResponse)
}

func requireBodyMatchTodoResponse(t *testing.T, body *bytes.Buffer, expected *domain.TodoResponse) {
	b, err := io.ReadAll(body)
	require.NoError(t, err)

	var todoResponse *domain.TodoResponse
	err = json.Unmarshal(b, &todoResponse)
	require.NoError(t, err)
	require.Equal(t, expected, todoResponse)
	require.NotZero(t, todoResponse.ID)
}

func requireBodyMatchTodoResponses(t *testing.T, body *bytes.Buffer, expected []*domain.TodoResponse) {
	b, err := io.ReadAll(body)
	require.NoError(t, err)

	var todoResponses []*domain.TodoResponse
	err = json.Unmarshal(b, &todoResponses)
	require.NoError(t, err)
	require.Equal(t, expected, todoResponses)
}