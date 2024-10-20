package music

import (
	"github.com/meraiku/music_lib/internal/lib/fetcher"
	"github.com/meraiku/music_lib/internal/repo"
	"github.com/meraiku/music_lib/pkg/logging"
)

type service struct {
	repo    repo.MusicRepository
	log     *logging.Logger
	fetcher fetcher.Fetcher
}

func NewService(
	repo repo.MusicRepository,
	log *logging.Logger,
	fetcher fetcher.Fetcher,
) *service {
	return &service{
		repo:    repo,
		log:     log,
		fetcher: fetcher,
	}
}
