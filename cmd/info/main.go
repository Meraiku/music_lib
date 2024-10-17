package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/meraiku/music_lib/internal/api/rest"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/info", handleInfo)

	srv := http.Server{
		Addr:    ":2000",
		Handler: r,
	}

	fmt.Printf("Info service starting at %s\n", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func handleInfo(w http.ResponseWriter, r *http.Request) {
	var req InfoRequest

	if err := decodeIntoStruct(r, &req); err != nil {
		if errors.Is(err, rest.ErrNoBody) {
			JSON(w, http.StatusBadRequest, struct{}{}) //nolint:errcheck
			return
		}
		JSON(w, http.StatusInternalServerError, struct{}{}) //nolint:errcheck
		return
	}

	if err := req.validate(); err != nil {
		JSON(w, http.StatusBadRequest, struct{}{}) //nolint:errcheck
		return
	}

	resp := InfoResponse{
		ReleaseDate: "16.07.2006",
		Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
		Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
	}

	JSON(w, http.StatusOK, resp) //nolint:errcheck
}
