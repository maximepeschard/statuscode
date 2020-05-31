package middleware

import (
	"net/http"

	"github.com/felixge/httpsnoop"
	"github.com/rs/zerolog/log"
)

// Logging returns a log-wrapped http.Handler.
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(next, w, r)
		log.Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("code", m.Code).
			Dur("duration", m.Duration).
			Send()
	})
}
