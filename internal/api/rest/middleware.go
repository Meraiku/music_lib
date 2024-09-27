package rest

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/meraiku/music_lib/pkg/logging"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

func (i *Implementation) logRequest(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := wrapResponseWriter(w)

		ctx := logging.WithLogRequestID(r.Context(), uuid.NewString())
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)

		done := time.Since(start)
		i.log.InfoContext(r.Context(), "Request",
			slog.String("from", r.RemoteAddr),
			slog.String("method", r.Method),
			slog.String("path", r.URL.String()),
			slog.Int("code", rw.Status()),
			slog.Duration("latency", done),
		)
	})
}
