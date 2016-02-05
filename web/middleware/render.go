package middleware

import (
	"net/http"

	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
)

var rend *render.Render

func init() {
	rend = render.New(render.Options{
		Layout:    "layout",
		Directory: "./views",
	})
}

// RenderMiddleware is used to expose render template to request
func RenderMiddleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c.Env["render"] = rend
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
