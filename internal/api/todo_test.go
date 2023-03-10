package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stovenn/gotodo/internal/core/domain"
	mockservice "github.com/stovenn/gotodo/internal/core/services/mock"
	"github.com/stovenn/gotodo/pkg/token"
	"github.com/stovenn/gotodo/pkg/util"
	"github.com/stretchr/testify/require"
)

func setupAuth(t *testing.T, request *http.Request, tokenMaker token.Maker) {
	addAuthorization(t, request, tokenMaker, authTypeBearer, "user", time.Minute)
}

func TestTodoHandler_HandleCreateTodo(t *testing.T) {
	todoResponse := util.CreateRandomTodoResponse(1)
	ctrl := gomock.NewController(t)

	service := mockservice.NewMockTodoService(ctrl)
	service.EXPECT().
		CreateTodo(gomock.Any()).
		Times(1).
		Return(todoResponse, nil)

	server := newTestServer(t, service, nil)
	body := strings.NewReader(`{"title": "new todo"}`)
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/api/todos/", body)
	require.NoError(t, err)

	setupAuth(t, request, server.tokenMaker)
	server.Server.Handler.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusCreated, recorder.Code)
	requireBodyMatchTodoResponse(t, recorder.Body, todoResponse)
}

func TestTodoHandler_HandleListTodo(t *testing.T) {
	todoResponses := util.CreateRandomTodoResponses(4)
	ctrl := gomock.NewController(t)

	service := mockservice.NewMockTodoService(ctrl)
	service.EXPECT().
		DisplayAllTodos().
		Times(1).
		Return(todoResponses, nil)

	server := newTestServer(t, service, nil)
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/api/todos/", nil)
	require.NoError(t, err)

	setupAuth(t, request, server.tokenMaker)
	server.Server.Handler.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchTodoResponses(t, recorder.Body, todoResponses)
}

func TestTodoHandler_HandleFindTodoByID(t *testing.T) {
	todoResponse := util.CreateRandomTodoResponse(1)
	ctrl := gomock.NewController(t)

	service := mockservice.NewMockTodoService(ctrl)
	// build stubs
	service.EXPECT().
		DisplayTodo(todoResponse.ID).
		Times(1).
		Return(todoResponse, nil)

	// start test server and send request
	server := newTestServer(t, service, nil)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/api/todos/%s", todoResponse.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	setupAuth(t, request, server.tokenMaker)
	server.Server.Handler.ServeHTTP(recorder, request)
	// check response
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchTodoResponse(t, recorder.Body, todoResponse)
}

func TestHandler_HandlePutTodo(t *testing.T) {
	todoResponse := util.CreateRandomTodoResponse(1)
	update := domain.TodoUpdateRequest{
		Title:      "updated todo",
		Completed:  true,
		Order:      1,
		AssignedTo: "user",
	}
	expectedResponse := &domain.TodoResponse{
		ID:         todoResponse.ID,
		Title:      update.Title,
		Completed:  update.Completed,
		Order:      update.Order,
		AssignedTo: update.AssignedTo,
		URL:        todoResponse.URL,
	}

	ctrl := gomock.NewController(t)
	service := mockservice.NewMockTodoService(ctrl)
	// build stubs
	service.EXPECT().
		UpdateTodo(todoResponse.ID, update).
		Times(1).
		Return(expectedResponse, nil)

	// start test server and send request
	server := newTestServer(t, service, nil)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/api/todos/%s", todoResponse.ID)
	body := strings.NewReader(`{"title": "updated todo", "order": 1, "completed": true, "assigned_to": "user"}`)
	request, err := http.NewRequest(http.MethodPut, url, body)
	require.NoError(t, err)

	setupAuth(t, request, server.tokenMaker)
	server.Server.Handler.ServeHTTP(recorder, request)
	// check response
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchTodoResponse(t, recorder.Body, expectedResponse)
}

func TestHandler_HandlePatchTodo(t *testing.T) {
	todoResponse := util.CreateRandomTodoResponse(1)
	update := domain.TodoPartialUpdateRequest{
		Order: 2,
	}
	expectedResponse := &domain.TodoResponse{
		ID:        todoResponse.ID,
		Title:     todoResponse.Title,
		Completed: todoResponse.Completed,
		Order:     update.Order,
		URL:       todoResponse.URL,
	}

	ctrl := gomock.NewController(t)
	service := mockservice.NewMockTodoService(ctrl)
	// build stubs
	service.EXPECT().
		PartiallyUpdateTodo(todoResponse.ID, update).
		Times(1).
		Return(expectedResponse, nil)

	// start test server and send request
	server := newTestServer(t, service, nil)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/api/todos/%s", todoResponse.ID)
	body := strings.NewReader(`{"order": 2}`)
	request, err := http.NewRequest(http.MethodPatch, url, body)
	require.NoError(t, err)

	setupAuth(t, request, server.tokenMaker)
	server.Server.Handler.ServeHTTP(recorder, request)
	// check response
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchTodoResponse(t, recorder.Body, expectedResponse)
}

func TestHandler_HandleDeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	service := mockservice.NewMockTodoService(ctrl)
	// build stubs
	service.EXPECT().DeleteTodo(gomock.Any()).Times(1).Return(nil)

	// start test server and send request
	server := newTestServer(t, service, nil)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/api/todos/%s", "1")
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	require.NoError(t, err)

	setupAuth(t, request, server.tokenMaker)
	server.Server.Handler.ServeHTTP(recorder, request)
	// check response
	require.Equal(t, http.StatusOK, recorder.Code)
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
