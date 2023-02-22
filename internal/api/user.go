package api

import (
	"encoding/json"
	"github.com/stovenn/gotodo/internal/core/domain"
	"net/http"
)

func (s *Server) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	var request domain.UserCreationRequest
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
	var credentials domain.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		handleError(w, http.StatusUnauthorized, err)
	}

	_, err = s.tokenMaker.CreateToken(credentials.Email, s.config.TokenDuration)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
	}

}
