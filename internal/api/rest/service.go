package rest

import (
	"net/http"

	"go.uber.org/zap"
)

type Implementation struct {
	Router http.Handler
	log    *zap.Logger
}

func NewImplementation(log *zap.Logger) *Implementation {
	return &Implementation{
		Router: SetRoutes(),
		log:    log,
	}
}
