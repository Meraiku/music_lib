package music

import "go.uber.org/zap"

type service struct {
	repo repository.Repository
	log  *zap.Logger
}

func NewService(
	repo repository.Repository,
	log *zap.Logger,
) *service {
	return &service{
		repo: repo,
		log:  log,
	}
}
