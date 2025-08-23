package auth

import (
	"api/configs"
	"api/pkg/response"
	"encoding/json"
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

func (handler *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("register")
		fmt.Println(handler.Config.Auth.Secret)
	}
}

func (*Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var payload LoginRequest

		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			response.Json(w, err.Error(), 402)
		}
		fmt.Println(payload)

		res := LoginResponse{Token: "123"}
		response.Json(w, res, 200)
	}
}
