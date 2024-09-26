package logging

import "log/slog"

type Logger struct {
	*slog.Logger
}

func Init(env string) *Logger {
	return &Logger{
		initSlog(env),
	}
}
