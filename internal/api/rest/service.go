package rest

import (
	"net/http"

	"github.com/meraiku/music_lib/internal/service"
	"github.com/meraiku/music_lib/pkg/logging"
)

type Implementation struct {
	router       http.Handler
	log          *logging.Logger
	musicService service.MusicService
}

func NewImplementation(musicService service.MusicService, log *logging.Logger) *Implementation {
	return &Implementation{
		log:          log,
		musicService: musicService,
	}
}

func (i *Implementation) Handler() http.Handler {
	if i.router == nil {
		i.router = i.setRoutes()
	}

	return i.router
}
