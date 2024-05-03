package routes

import (
	"net/http"

	"github.com/aliflazuardi/cats-social/internal/server"
)

func GetRoutesHandler() *http.ServeMux {
	mux := http.NewServeMux()

	h := &server.AuthHandler{}
	getAuthHandler(mux, h)

	return mux
}

func getAuthHandler(mux *http.ServeMux, authHandler *server.AuthHandler) {
	mux.HandleFunc("/", authHandler.HellowWorldHandler)
	mux.HandleFunc("POST /v1/user/register", authHandler.RegisterUserHandler)
	mux.HandleFunc("POST /v1/user/login", authHandler.LoginHandler)
}
