package uniuri

import "github.com/dchest/uniuri"

type UniURI struct {
}

func New() *UniURI {
	return &UniURI{}
}

func (g *UniURI) Generate() string {
	return uniuri.New()
}
