package config

import (
	"net"
	"os"
)

type RESTConfig struct {
	host string
	port string
}

func (cfg RESTConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

func NewRESTConfig() *RESTConfig {
	cfg := &RESTConfig{}

	cfg.host = os.Getenv("HOST")

	cfg.port = os.Getenv("PORT")

	if cfg.port == "" {
		cfg.port = "8080"
	}

	return cfg
}
