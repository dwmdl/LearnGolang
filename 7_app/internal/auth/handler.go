package auth

import (
	"api/configs"
	"api/pkg/jwt"
	"api/pkg/request"
	"api/pkg/response"
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
			return
		}

		email, err := handler.Service.Register(body.Email, body.Password, body.Name)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.Data{Email: email})
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := RegisterResponse{Token: token}

		response.Json(w, data, http.StatusCreated)
	}
}

func (handler *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[LoginRequest](&w, req)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		email, err := handler.Service.Login(body.Email, body.Password)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.Data{Email: email})
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := LoginResponse{
			Token: token,
		}

		response.Json(w, data, http.StatusOK)
	}
}
