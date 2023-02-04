package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/core/ports"
	"net/http"
)

type TodoHandler struct {
	S ports.TodoService
}

func NewTodoHandler(todoService ports.TodoService) *TodoHandler {
	return &TodoHandler{S: todoService}
}

func (t *TodoHandler) HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
	var request domain.TodoCreationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	response, err := t.S.CreateTodo(request)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	withJSON(w, http.StatusCreated, response.JSON())
}

func (t *TodoHandler) HandleListTodo(w http.ResponseWriter, r *http.Request) {
	response, err := t.S.DisplayAllTodos()
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	var b []byte
	b = append(b, byte('['))
	for i, r := range response {
		b = append(b, r.JSON()...)
		if i == len(response)-1 {
			break
		}
		b = append(b, byte(','))
	}
	b = append(b, byte(']'))
	withJSON(w, http.StatusOK, b)
}

func (t *TodoHandler) HandleFindTodoByID(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["id"]
	response, err := t.S.DisplayTodo(todoId)
	if err != nil {
		handleError(w, http.StatusNotFound, err)
		return
	}
	withJSON(w, http.StatusOK, response.JSON())
}

func (t *TodoHandler) HandlePutTodo(w http.ResponseWriter, r *http.Request) {
	var request domain.TodoUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	err = validate.Struct(request)
	if err != nil {
		handleError(w, http.StatusBadRequest, err)
		return
	}
	todoID := mux.Vars(r)["id"]
	response, err := t.S.UpdateTodo(todoID, request)
	if err != nil {
		handleError(w, http.StatusNotFound, err)
		return
	}
	withJSON(w, http.StatusOK, response.JSON())
}

func (t *TodoHandler) HandlePatchTodo(w http.ResponseWriter, r *http.Request) {
	var request domain.TodoPartialUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	todoID := mux.Vars(r)["id"]
	response, err := t.S.PartiallyUpdateTodo(todoID, request)
	if err != nil {
		handleError(w, http.StatusNotFound, err)
		return
	}
	withJSON(w, http.StatusOK, response.JSON())
}

func (t *TodoHandler) HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["id"]
	err := t.S.DeleteTodo(todoId)
	if err != nil {
		handleError(w, http.StatusNotFound, err)
	}
}
