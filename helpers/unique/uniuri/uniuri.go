package uniuri

import "stablelib.com/v1/uniuri"

type UniURI struct {
}

func New() *UniURI {
	return &UniURI{}
}

func (g *UniURI) Generate() string {
	return uniuri.New()
}
