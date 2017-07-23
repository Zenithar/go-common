package sessions

type contextKey string

func (c contextKey) String() string {
	return "go.zenithar.org/common/web/middleware/" + string(c)
}
