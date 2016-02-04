package registrar

import "sync"

// default package instance
var (
	instance Registrar
	once     sync.Once
)

// Registry returns the registry singleton
func Registry() Registrar {
	once.Do(func() {
		instance = newRegistrar()
	})
	return instance
}
