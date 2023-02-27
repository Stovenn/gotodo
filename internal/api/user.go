package api

import (
	"encoding/json"
	"net/http"

	"github.com/stovenn/gotodo/internal/dto/request"
)

func (s *Server) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	var request request.UserCreationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	response, err := s.UserService.SignUp(request)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	withJSON(w, http.StatusCreated, response.JSON())
}

func (s *Server) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var credentials request.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		handleError(w, http.StatusUnauthorized, err)
	}

	response, err := s.UserService.Login(credentials, s.tokenMaker, s.config)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
	}

	withJSON(w, http.StatusOK, response.JSON())
}
