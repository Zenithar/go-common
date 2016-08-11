package middleware

import (
	"net/http"

	"goji.io"

	"golang.org/x/net/context"

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
func RenderMiddleware(h goji.Handler) goji.Handler {
	fn := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		ctx = context.WithValue(ctx, "render", rend)
		h.ServeHTTPC(ctx, w, r)
	}
	return goji.HandlerFunc(fn)
}
