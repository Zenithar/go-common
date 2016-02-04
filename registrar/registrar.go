package registrar

import "sync"

// Registrar defines the default registry contract
type Registrar interface {
	Register(string, interface{})
	Lookup(string) interface{}
}

// registrar is like spring registry but for go
type registrar struct {
	values map[string]interface{}
	sync.Mutex
}

// NewRegistrar returns an empty registy
func newRegistrar() Registrar {
	return &registrar{
		values: map[string]interface{}{},
	}
}

// Register a bean in the registry
func (r *registrar) Register(k string, v interface{}) {
	if r == nil {
		return
	}

	r.Lock()
	defer r.Unlock()
	r.values[k] = v
}

// Lookup a bean from the registry
func (r *registrar) Lookup(k string) interface{} {
	if r == nil {
		return nil
	}

	r.Lock()
	defer r.Unlock()
	return r.values[k]
}
