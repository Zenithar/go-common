package cache

import (
	"errors"
	"time"
)

const (
	DEFAULT = time.Duration(0)
	FOREVER = time.Duration(-1)
)

var (
	CachePrefix   = "default"
	ErrCacheMiss  = errors.New("cache: key not found.")
	ErrNotStored  = errors.New("cache: not stored.")
	ErrNotSupport = errors.New("cache: not support.")
)

// CacheStore is the basic interface for cache implementations
type CacheStore interface {
	Get(key string, value interface{}) error
	Set(key string, value interface{}, expire time.Duration) error
	Add(key string, value interface{}, expire time.Duration) error
	Replace(key string, data interface{}, expire time.Duration) error
	Delete(key string) error
	Increment(key string, data uint64) (uint64, error)
	Decrement(key string, data uint64) (uint64, error)
	Flush() error
}
