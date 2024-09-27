package logging

import (
	"fmt"
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
)

func initSlog(_ string) *slog.Logger {

	f := openLoggingFile()

	h := slogmulti.Fanout(
		slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}),
		slog.NewJSONHandler(f, &slog.HandlerOptions{Level: slog.LevelInfo, ReplaceAttr: replaceAttr}),
	)

	l := slog.New(NewHandlerMiddleware(h))

	return l
}

func openLoggingFile() *os.File {
	if err := os.MkdirAll("logs", 0755); err != nil {
		panic(fmt.Errorf("error creating 'logs' directory: %s", err))
	}

	f, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("error oppening logs file: %s", err))
	}
	return f
}
