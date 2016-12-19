package middleware

import (
	"net/http"

	"golang.org/x/net/context"
)

// XRequestID is a goji middleware to track requestID
func XRequestID(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("X-Request-ID")) > 0 {
			ctx := context.WithValue(r.Context(), "reqID", r.Header.Get("X-Request-ID"))
			h.ServeHTTP(w, r.WithContext(ctx))
		} else {
			h.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(fn)
}
