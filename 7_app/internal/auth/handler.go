package auth

import (
	"api/configs"
	"api/pkg/request"
	"api/pkg/response"
	"fmt"
	"net/http"
)

type HandlerDeps struct {
	*configs.Config
}

type Handler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{Config: deps.Config}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (*Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, req)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
		}

		fmt.Println(*body)
	}
}

func (*Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[LoginRequest](&w, req)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(*body)

		res := LoginResponse{Token: "123"}
		response.Json(w, res, http.StatusOK)
	}
}
