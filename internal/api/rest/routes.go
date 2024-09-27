package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (i *Implementation) setRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(i.logRequest)
	r.Use(middleware.Recoverer)

	r.Get("/healthz", statusCheck)

	r.Get("/swagger/*", httpSwagger.Handler())

	r.Route("/api", func(r chi.Router) {

		r.Route("/songs", func(r chi.Router) {
			r.Get("/", i.Make(i.GetSongs))
			r.Post("/", i.Make(i.PostSong))
			r.Put("/{id}", i.Make(i.UpdateSong))
			r.Delete("/{id}", i.Make(i.DeleteSong))

			r.Get("/{id}/text", i.Make(i.GetText))
		})
	})

	return r
}
