package uuid

import (
	"encoding/base64"
	"encoding/hex"
	"strings"

	"github.com/satori/go.uuid"
)

type UUIDV1 struct{}

func NewV1() *UUIDV1 {
	return &UUIDV1{}
}

func (g *UUIDV1) Generate() string {
	v1 := uuid.NewV1()
	data := v1.Bytes()
	buf := make([]byte, 32)
	hex.Encode(buf, data)
	return string(buf)
}

type UUIDV4 struct{}

func NewV4() *UUIDV4 {
	return &UUIDV4{}
}

func (g *UUIDV4) Generate() string {
	v4 := uuid.NewV4()
	data := v4.Bytes()
	buf := make([]byte, 32)
	hex.Encode(buf, data)
	return string(buf)
}

type UUID1Base64 struct{}

func NewV1Base64() *UUID1Base64 {
	return &UUID1Base64{}
}

func (g *UUID1Base64) Generate() string {
	id := base64.URLEncoding.EncodeToString(uuid.NewV1().Bytes())
	return strings.TrimRight(id, "=")
}
