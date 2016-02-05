package middleware

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

// XRequestID is a goji middleware to track requestID
func XRequestID(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if c.Env == nil {
			c.Env = make(map[interface{}]interface{})
		}

		if len(r.Header.Get("X-Request-ID")) > 0 {
			c.Env["reqID"] = r.Header.Get("X-Request-ID")
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
