package snowflake

import (
	"strconv"

	"github.com/sdming/gosnow"
)

type Snowflake struct {
	gosnow *gosnow.SnowFlake
}

func New(index int) *Snowflake {
	gosnow, err := gosnow.NewSnowFlake(uint32(index))
	if err != nil {
		panic(err)
	}
	return &Snowflake{gosnow}
}

func (g *Snowflake) Generate() string {
	result, _ := g.gosnow.Next()
	return strconv.FormatInt(int64(result), 10)
}
