package logging

import (
	"context"
	"log/slog"
)

type Logger struct {
	*slog.Logger
}

func Init(env string) *Logger {
	return &Logger{
		initSlog(env),
	}
}

func WithLogRequestID(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, key, logCtx{RequestId: reqID})
}
