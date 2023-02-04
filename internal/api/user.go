package api

import (
	"encoding/json"
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/core/ports"
	"net/http"
)

type UserHandler struct {
	S ports.UserService
}

func NewUserHandler(userService ports.UserService) *UserHandler {
	return &UserHandler{S: userService}
}

func (u UserHandler) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	var request domain.UserCreationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	response, err := u.S.SignUp(request)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	withJSON(w, http.StatusCreated, response.JSON())
}

func (u UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
