package app

import (
	"github.com/meraiku/music_lib/internal/api/rest"
	"github.com/meraiku/music_lib/internal/config"
	"github.com/meraiku/music_lib/internal/repo"
	"github.com/meraiku/music_lib/internal/repo/postgre"
	"github.com/meraiku/music_lib/internal/service"
	"github.com/meraiku/music_lib/internal/service/music"
	"go.uber.org/zap"
)

type serviceProvider struct {
	log          *zap.Logger
	restImpl     *rest.Implementation
	cfg          *config.RESTConfig
	repo         repo.MusicRepository
	musicService service.MusicService
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Config() *config.RESTConfig {
	if s.cfg == nil {
		s.cfg = config.NewRESTConfig()
	}

	return s.cfg
}

func (s *serviceProvider) Logger() *zap.Logger {
	if s.log == nil {
		s.log, _ = zap.NewDevelopment()
	}

	return s.log
}

func (s *serviceProvider) RestImpl() *rest.Implementation {
	if s.restImpl == nil {
		s.restImpl = rest.NewImplementation(s.MusicService(), s.Logger())
	}

	return s.restImpl
}

func (s *serviceProvider) Repository() repo.MusicRepository {
	if s.repo == nil {
		var err error
		s.repo, err = postgre.New()
		if err != nil {
			zap.L().Error("connecting db",
				zap.String("error", err.Error()),
			)
		}
	}

	return s.repo
}

func (s *serviceProvider) MusicService() service.MusicService {
	if s.musicService == nil {
		s.musicService = music.NewService(s.Repository(), s.Logger())
	}

	return s.musicService
}
