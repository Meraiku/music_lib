package rest

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func (i *Implementation) logRequest(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer i.log.Sync()

		start := time.Now()

		next.ServeHTTP(w, r)

		done := time.Since(start)

		i.log.Info("Request",
			zap.String("from", r.RemoteAddr),
			zap.String("method", r.Method),
			zap.String("path", r.URL.String()),
			zap.Duration("in", done),
		)
	})
}
