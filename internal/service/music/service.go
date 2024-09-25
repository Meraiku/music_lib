package music

import (
	"github.com/meraiku/music_lib/internal/repo"
	"go.uber.org/zap"
)

type service struct {
	repo *repo.MusicRepository
	log  *zap.Logger
}

func NewService(
	repo *repo.MusicRepository,
	log *zap.Logger,
) *service {
	return &service{
		repo: repo,
		log:  log,
	}
}
