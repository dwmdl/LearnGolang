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
	*Service
}

type Handler struct {
	*configs.Config
	*Service
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		Config:  deps.Config,
		Service: deps.Service,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, req)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
		}

		_, err = handler.Service.Register(body.Email, body.Password, body.Name)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
		}
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
