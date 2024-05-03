package server

import (
	"fmt"
	"net/http"

	"github.com/aliflazuardi/cats-social/internal/auth"
)

type AuthHandler struct{}

func (authHandler *AuthHandler) HellowWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Cats")
}

func (authHandler *AuthHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	auth.Register(w, r)
}

func (authHandler *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	auth.Login(w, r)
}

type CatHandler struct{}

type MatchHandler struct{}
