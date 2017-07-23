package renderer

type contextKey string

func (c contextKey) String() string {
	return "go.zenithar.org/common/web/middleware/" + string(c)
}
