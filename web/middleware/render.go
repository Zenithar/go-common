package middleware

import (
	"context"
	"net/http"

	"github.com/unrolled/render"
)

var rend *render.Render

func init() {
	rend = render.New(render.Options{
		Layout:    "layout",
		Directory: "./views",
	})
}

// RenderMiddleware is used to expose render template to request
func RenderMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "render", rend)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
