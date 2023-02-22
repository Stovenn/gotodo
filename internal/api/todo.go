package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stovenn/gotodo/internal/core/domain"
	"net/http"
)

func (s *Server) HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
	var request domain.TodoCreationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	response, err := s.TodoService.CreateTodo(request)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	withJSON(w, http.StatusCreated, response.JSON())
}

func (s *Server) HandleListTodo(w http.ResponseWriter, r *http.Request) {
	response, err := s.TodoService.DisplayAllTodos()
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

func (s *Server) HandleFindTodoByID(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["id"]
	response, err := s.TodoService.DisplayTodo(todoId)
	if err != nil {
		handleError(w, http.StatusNotFound, err)
		return
	}
	withJSON(w, http.StatusOK, response.JSON())
}

func (s *Server) HandlePutTodo(w http.ResponseWriter, r *http.Request) {
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
	response, err := s.TodoService.UpdateTodo(todoID, request)
	if err != nil {
		handleError(w, http.StatusNotFound, err)
		return
	}
	withJSON(w, http.StatusOK, response.JSON())
}

func (s *Server) HandlePatchTodo(w http.ResponseWriter, r *http.Request) {
	var request domain.TodoPartialUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	todoID := mux.Vars(r)["id"]
	response, err := s.TodoService.PartiallyUpdateTodo(todoID, request)
	if err != nil {
		handleError(w, http.StatusNotFound, err)
		return
	}
	withJSON(w, http.StatusOK, response.JSON())
}

func (s *Server) HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["id"]
	err := s.TodoService.DeleteTodo(todoId)
	if err != nil {
		handleError(w, http.StatusNotFound, err)
	}
}
