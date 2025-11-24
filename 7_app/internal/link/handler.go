package link

import (
	"api/configs"
	"api/pkg/middleware"
	"api/pkg/request"
	"api/pkg/response"
	"net/http"
	"strconv"
)

type HandlerDeps struct {
	LinkRepo *Repository
	Config   *configs.Config
}

type Handler struct {
	Repository *Repository
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{deps.LinkRepo}
	router.HandleFunc("POST /link", handler.Create())
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.Update(), deps.Config))
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

		for {
			existedLink, _ := handler.Repository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}

			link.GenerateHash()
		}

		createdLink, err := handler.Repository.Create(link)
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
		link, err := handler.Repository.GetByHash(hash)
		if err != nil {
			response.Json(w, err, http.StatusNotFound)
			return
		}

		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[UpdateRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		link, err := handler.Repository.Update(&Link{
			ID:   uint(id),
			Url:  body.Url,
			Hash: body.Hash,
		})
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
		}

		response.Json(w, link, http.StatusOK)
	}
}

func (handler *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = handler.Repository.GetById(id)
		if err != nil {
			response.Json(w, err.Error(), http.StatusNotFound)
			return
		}

		err = handler.Repository.Delete(id)
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Json(w, nil, http.StatusOK)
	}
}
