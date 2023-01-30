package api

import (
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

}

func (u UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {

}
