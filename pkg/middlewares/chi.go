package middlewares

import (
	"net/http"

	"github.com/pzabolotniy/logging/pkg/logging"
)

// WithLogger is a chi-router middleware to inject logger into the context.
func WithLogger(logger logging.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctxWithLogger := logging.WithContext(ctx, logger)
			r = r.WithContext(ctxWithLogger)
			next.ServeHTTP(w, r)
		})
	}
}
