package renderer

import (
	"context"
	"net/http"

	"github.com/unrolled/render"
)

const (
	contextRenderer = contextKey("renderer")
)

// NewMiddleware is used to expose render template to request
func NewMiddleware(render *render.Render) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), contextRenderer, render)
			h.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

// FromContext returns renderer from context object
func FromContext(ctx context.Context) (*render.Render, bool) {
	value, ok := ctx.Value(contextRenderer).(*render.Render)
	return value, ok
}
