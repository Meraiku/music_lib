package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetRoutes() http.Handler {
	r := chi.NewRouter()

	r.With(middleware.Recoverer)

	r.Get("/healthz", statusCheck)

	return r
}
