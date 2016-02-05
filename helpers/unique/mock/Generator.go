package mock

import "github.com/stretchr/testify/mock"

type Generator struct {
	mock.Mock
}

func (_m *Generator) Generate() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
