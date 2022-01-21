package cache

import (
	"github.com/curatorc/cngf/config"
	"github.com/curatorc/cngf/redis"
	"time"
)

// RedisStore 实现 cache.Store interface
type RedisStore struct {
	RedisClient *redis.Client
	KeyPrefix   string
}

func NewRedisStore(address string, username string, password string, db int) *RedisStore {
	rs := &RedisStore{}
	rs.RedisClient = redis.New(address, username, password, db)
	rs.KeyPrefix = config.GetString("app.name") + ":cache:"
	return rs
}

func (s *RedisStore) Set(key string, value string, expireTime time.Duration) {
	s.RedisClient.Set(s.KeyPrefix+key, value, expireTime)
}

func (s *RedisStore) Get(key string) string {
	return s.RedisClient.Get(s.KeyPrefix + key)
}

func (s *RedisStore) Has(key string) bool {
	return s.RedisClient.Has(s.KeyPrefix + key)
}

func (s *RedisStore) Forget(key string) {
	s.RedisClient.Del(s.KeyPrefix + key)
}

func (s *RedisStore) Forever(key string, value string) {
	s.RedisClient.Set(s.KeyPrefix+key, value, 0)
}

func (s *RedisStore) Flush() {
	s.RedisClient.FlushDB()
}

func (s *RedisStore) Increment(key string) {
	s.RedisClient.Increment(key)
}

func (s *RedisStore) Increments(key string, count int64) {
	s.RedisClient.Increments(key, count)
}

func (s *RedisStore) Decrement(key string) {
	s.RedisClient.Decrement(key)
}

func (s *RedisStore) Decrements(key string, count int64) {
	s.RedisClient.Decrements(key, count)
}

func (s *RedisStore) IsAlive() error {
	return s.RedisClient.Ping()
}
