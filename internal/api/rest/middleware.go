package rest

import (
	"net/http"
	"time"

	"go.uber.org/zap"
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
		defer i.log.Sync()

		start := time.Now()

		rw := wrapResponseWriter(w)

		next.ServeHTTP(rw, r)

		done := time.Since(start)

		i.log.Info("Request",
			zap.String("from", r.RemoteAddr),
			zap.String("method", r.Method),
			zap.String("path", r.URL.String()),
			zap.Int("code", rw.Status()),
			zap.Duration("latency", done),
		)
	})
}
