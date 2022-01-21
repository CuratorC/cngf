package cache

import "time"

type Store interface {
	Set(key string, value string, expireTime time.Duration)
	Get(key string) string
	Has(key string) bool
	Forget(key string)
	Forever(key string, value string)
	Flush()

	IsAlive() error

	Increment(key string)
	Increments(key string, count int64)

	Decrement(key string)
	Decrements(key string, count int64)
}
