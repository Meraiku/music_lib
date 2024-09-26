package logging

import (
	"fmt"
	"log/slog"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
)

func initSlog(env string) *slog.Logger {

	createLoggingDir()

	cfg := zap.Config{
		Encoding:         "json",
		OutputPaths:      []string{"stdout", "logs/all.log"},
		ErrorOutputPaths: []string{"stderr", "logs/all.log"},
	}

	switch env {
	case "prod":
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	default:
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		cfg.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		cfg.Development = true
	}

	zapL, _ := cfg.Build()

	h := zapslog.NewHandler(zapL.Core(), nil)

	l := slog.New(NewHandlerMiddleware(h))

	return l
}

func createLoggingDir() {
	if err := os.MkdirAll("logs", 0755); err != nil {
		panic(fmt.Errorf("error creating 'logs' directory: %s", err))
	}
}
