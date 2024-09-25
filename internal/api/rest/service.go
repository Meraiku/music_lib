package rest

import (
	"net/http"

	"github.com/meraiku/music_lib/internal/service"
	"go.uber.org/zap"
)

type Implementation struct {
	router       http.Handler
	log          *zap.Logger
	musicService service.MusicService
}

func NewImplementation(musicService service.MusicService, log *zap.Logger) *Implementation {
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
