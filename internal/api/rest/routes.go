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

	r.Route("/api", func(r chi.Router) {

		r.Route("/songs", func(r chi.Router) {
			r.Get("/", i.GetSongs)
			r.Post("/", i.PostSong)
			r.Put("/", i.UpdateSong)
			r.Delete("/", i.DeleteSong)

			r.Get("/text", i.GetText)
		})
	})

	return r
}
