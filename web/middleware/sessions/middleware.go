package sessions

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	contextSessions = contextKey("sessions")
)

// NewMiddleware is used to expose render template to request
func NewMiddleware(name string, store sessions.Store) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			session, _ := store.Get(r, name)
			ctx := context.WithValue(r.Context(), contextSessions, session)
			h.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

// FromContext returns renderer from context object
func FromContext(ctx context.Context) (*sessions.Session, bool) {
	value, ok := ctx.Value(contextSessions).(*sessions.Session)
	return value, ok
}
