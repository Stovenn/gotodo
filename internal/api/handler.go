package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/core/ports"
	"net/http"
)

type Handler struct {
	S ports.TodoService
}

func NewHandler(todoService ports.TodoService) *Handler {
	return &Handler{S: todoService}
}

func (t *Handler) HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
	var request domain.TodoCreationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, err)
		return
	}

	response, err := t.S.AddTodo(request)
	if err != nil {
		handleError(w, err)
		return
	}
	withJSON(w, http.StatusCreated, response)
}

func (t *Handler) HandleListTodo(w http.ResponseWriter, r *http.Request) {
	response, err := t.S.ListTodos()
	if err != nil {
		handleError(w, err)
		return
	}
	withJSON(w, http.StatusOK, response...)
}

func (t *Handler) HandleFindTodoByID(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["id"]
	response, err := t.S.FindTodoByID(todoId)
	if err != nil {
		handleError(w, err)
		return
	}
	withJSON(w, http.StatusOK, response)
}

func (t *Handler) HandlePutTodo(w http.ResponseWriter, r *http.Request) {
	var request domain.TodoUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, err)
		return
	}
	todoID := mux.Vars(r)["id"]
	response, err := t.S.UpdateTodo(todoID, request)
	if err != nil {
		handleError(w, err)
		return
	}
	fmt.Println(response)
	withJSON(w, http.StatusOK, response)
}

func (t *Handler) HandlePatchTodo(w http.ResponseWriter, r *http.Request) {
	var request domain.TodoPartialUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, err)
		return
	}
	fmt.Println(request)
	todoID := mux.Vars(r)["id"]
	response, err := t.S.PartiallyUpdateTodo(todoID, request)
	if err != nil {
		handleError(w, err)
		return
	}
	fmt.Println(response)
	withJSON(w, http.StatusOK, response)
}

func (t *Handler) HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["id"]
	err := t.S.DeleteTodo(todoId)
	if err != nil {
		handleError(w, err)
	}
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func withJSON(w http.ResponseWriter, statusCode int, response ...*domain.TodoResponse) {
	var b []byte
	var err error
	if len(response) == 1 {
		b, err = json.Marshal(&response[0])
		if err != nil {
			handleError(w, err)
			return
		}
	} else {
		b, err = json.Marshal(&response)
		if err != nil {
			handleError(w, err)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(b)
}
