package app

import (
	"github.com/meraiku/music_lib/internal/api/rest"
	"go.uber.org/zap"
)

type serviceProvider struct {
	log      *zap.Logger
	restImpl *rest.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Logger() *zap.Logger {
	if s.log == nil {
		s.log, _ = zap.NewDevelopment()
	}

	return s.log
}

func (s *serviceProvider) RestImpl() *rest.Implementation {
	if s.restImpl == nil {
		s.restImpl = rest.NewImplementation(s.Logger())
	}

	return s.restImpl
}
