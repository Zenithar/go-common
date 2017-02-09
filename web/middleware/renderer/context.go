package renderer

type contextKey string

func (c contextKey) String() string {
	return "zenithar.org/go/common/web/middleware/" + string(c)
}
