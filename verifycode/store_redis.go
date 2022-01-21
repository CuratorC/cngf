package verifycode

import (
	"github.com/curatorc/cngf/app"
	"github.com/curatorc/cngf/cache"
	"github.com/curatorc/cngf/config"
	"time"
)

// RedisStore 实现 verifycode.Store interface
type RedisStore struct {
	KeyPrefix string
}

// Set 实现 verifycode.Store interface 的 Set 方法
func (s *RedisStore) Set(key string, value string) bool {

	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))
	// 本地环境方便调试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	}

	cache.Set(s.KeyPrefix+key, value, ExpireTime)
	return true
}

// Get 实现 verifycode.Store interface 的 Get 方法
func (s *RedisStore) Get(key string, clear bool) (value string) {
	key = s.KeyPrefix + key
	val := cache.GetString(key)
	if clear {
		cache.Forget(key)
	}
	return val
}

// Verify 实现 verifycode.Store interface 的 Verify 方法
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
