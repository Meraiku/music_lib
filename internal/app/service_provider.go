package app

import (
	"log"
	"os"

	"github.com/meraiku/music_lib/internal/api/rest"
	"github.com/meraiku/music_lib/internal/config"
	"github.com/meraiku/music_lib/internal/repo"
	"github.com/meraiku/music_lib/internal/repo/postgre"
	"github.com/meraiku/music_lib/internal/service"
	"github.com/meraiku/music_lib/internal/service/music"
	"github.com/meraiku/music_lib/pkg/logging"
)

type serviceProvider struct {
	log          *logging.Logger
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

func (s *serviceProvider) Logger() *logging.Logger {
	if s.log == nil {
		s.log = logging.Init(os.Getenv("ENV"))
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
		s.repo, err = postgre.New(s.Logger())
		if err != nil {
			log.Fatalf("error connecting db: %s", err)
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
