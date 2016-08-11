package middleware

import (
	"net/http"

	"goji.io"

	"golang.org/x/net/context"
)

// XRequestID is a goji middleware to track requestID
func XRequestID(h goji.Handler) goji.Handler {
	fn := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("X-Request-ID")) > 0 {
			ctx = context.WithValue(ctx, "reqID", r.Header.Get("X-Request-ID"))
		}

		h.ServeHTTPC(ctx, w, r)
	}

	return goji.HandlerFunc(fn)
}
