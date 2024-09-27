package music

import (
	"github.com/meraiku/music_lib/internal/repo"
	"github.com/meraiku/music_lib/pkg/logging"
)

type service struct {
	repo repo.MusicRepository
	log  *logging.Logger
}

func NewService(
	repo repo.MusicRepository,
	log *logging.Logger,
) *service {
	return &service{
		repo: repo,
		log:  log,
	}
}
