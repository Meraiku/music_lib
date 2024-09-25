package app

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/meraiku/music_lib/internal/config"
)

type App struct {
	serviceProvider *serviceProvider
}

func NewApp(ctx context.Context) (*App, error) {

	a := &App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runRestServer()
}

func (a *App) initDeps(ctx context.Context) error {

	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvicer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	if os.Getenv("ENV") != "" {
		return nil
	}

	return config.Load(".env")
}

func (a *App) initServiceProvicer(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *App) runRestServer() error {

	s := http.Server{
		Handler: a.serviceProvider.RestImpl().Router,
		Addr:    a.serviceProvider.Config().Address(),
	}

	a.serviceProvider.log.Info(fmt.Sprintf("Starting server at %s", s.Addr))

	return s.ListenAndServe()
}
