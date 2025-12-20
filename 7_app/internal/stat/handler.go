package stat

import (
	"api/configs"
	"api/pkg/middleware"
	"fmt"
	"net/http"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type HandlerDeps struct {
	StatRepo *Repository
	Config   *configs.Config
}

type Handler struct {
	Repository *Repository
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{deps.StatRepo}
	router.Handle("GET /stat", middleware.IsAuthed(handler.GetStat(), deps.Config))
}

func (handler *Handler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid 'from' param", http.StatusBadRequest)
			return
		}
		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid 'to' param", http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "Invalid 'by' param", http.StatusBadRequest)
			return
		}
		fmt.Println(from, to, by)
	}
}
