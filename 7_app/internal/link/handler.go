package link

import (
	"api/pkg/request"
	"api/pkg/response"
	"net/http"
)

type HandlerDeps struct {
	LinkRepo *Repository
}

type Handler struct {
	LinkRepo *Repository
}

func NewLinkHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{deps.LinkRepo}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.GoTo())
}

func (handler *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[CreateRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		link := NewLink(body.Url)
		createdLink, err := handler.LinkRepo.Create(link)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, createdLink, http.StatusCreated)
	}
}

func (handler *Handler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.LinkRepo.GetByHash(hash)
		if err != nil {
			response.Json(w, err, http.StatusNotFound)
			return
		}

		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
