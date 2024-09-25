package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (i *Implementation) setRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(i.logRequest)
	r.Use(middleware.Recoverer)

	r.Get("/healthz", statusCheck)

	return r
}
