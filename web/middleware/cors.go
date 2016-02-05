package middleware

import (
	"net/http"

	"github.com/martini-contrib/cors"
)

// CORS is a middleware to handle CORS requests
func CORS(h http.Handler) http.Handler {
	f := cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"PUT", "PATCH", "DELETE", "POST"},
	})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
		h.ServeHTTP(w, r)
	})
}
