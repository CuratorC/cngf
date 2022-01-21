// Package cache 缓存工具类，可以缓存各种类型包括 struct 对象
package cache

import (
	"encoding/json"
	"github.com/curatorc/cngf/helpers"
	"github.com/curatorc/cngf/logger"
	"sync"
	"time"

	"github.com/spf13/cast"
)

type Service struct {
	Store Store
}

var once sync.Once
var Cache *Service

func InitWithCacheStore(store Store) {
	once.Do(func() {
		Cache = &Service{
			Store: store,
		}
	})
}

func Set(key string, obj interface{}, expireTime time.Duration) {
	b, err := json.Marshal(&obj)
	logger.LogIf(err)
	Cache.Store.Set(key, string(b), expireTime)
}

func Get(key string) interface{} {
	stringValue := Cache.Store.Get(key)
	var wanted interface{}
	err := json.Unmarshal([]byte(stringValue), &wanted)
	logger.LogIf(err)
	return wanted
}

func Remember(key string, expireTime time.Duration, fu func() interface{}) (response interface{}) {
	// 取数据
	val := Cache.Store.Get(key)
	if len(val) > 0 { // 读取缓存
		err := json.Unmarshal([]byte(val), &response)
		logger.LogIf(err)
	} else {
		// 执行查询
		response = fu()
		if helpers.Empty(response) {
			return
		}
		// 设置缓存
		Set(key, response, expireTime)
	}
	return
}

func Has(key string) bool {
	return Cache.Store.Has(key)
}

// GetObject 应该传地址，用法如下:
//     model := user.User{}
//     cache.GetObject("key", &model)
func GetObject(key string, wanted interface{}) {
	val := Cache.Store.Get(key)
	if len(val) > 0 {
		err := json.Unmarshal([]byte(val), &wanted)
		logger.LogIf(err)
	}
}

func RememberObject(key string, expireTime time.Duration, wanted interface{}, fu func() interface{}) {

	// 取数据
	val := Cache.Store.Get(key)
	if len(val) > 0 { // 读取缓存
		err := json.Unmarshal([]byte(val), &wanted)
		logger.LogIf(err)
	} else {
		// 执行查询
		wanted = fu()
		if helpers.Empty(wanted) {
			return
		}
		// 设置缓存
		Set(key, wanted, expireTime)
	}
}

func GetString(key string) string {
	return cast.ToString(Get(key))
}

func GetBool(key string) bool {
	return cast.ToBool(Get(key))
}

func GetInt(key string) int {
	return cast.ToInt(Get(key))
}

func GetInt32(key string) int32 {
	return cast.ToInt32(Get(key))
}

func GetInt64(key string) int64 {
	return cast.ToInt64(Get(key))
}

func GetUint(key string) uint {
	return cast.ToUint(Get(key))
}

func GetUint32(key string) uint32 {
	return cast.ToUint32(Get(key))
}

func GetUint64(key string) uint64 {
	return cast.ToUint64(Get(key))
}

func GetFloat64(key string) float64 {
	return cast.ToFloat64(Get(key))
}

func GetTime(key string) time.Time {
	return cast.ToTime(Get(key))
}

func GetDuration(key string) time.Duration {
	return cast.ToDuration(Get(key))
}

func GetIntSlice(key string) []int {
	return cast.ToIntSlice(Get(key))
}

func GetStringSlice(key string) []string {
	return cast.ToStringSlice(Get(key))
}

func GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(Get(key))
}

func GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(Get(key))
}

func GetStringMapStringSlice(key string) map[string][]string {
	return cast.ToStringMapStringSlice(Get(key))
}

func Forget(key string) {
	Cache.Store.Forget(key)
}

func Forever(key string, value string) {
	Cache.Store.Set(key, value, 0)
}

func Flush() {
	Cache.Store.Flush()
}

func Increment(key string) {
	Cache.Store.Increment(key)
}

func Increments(key string, count int64) {
	Cache.Store.Increments(key, count)
}

func Decrement(key string) {
	Cache.Store.Decrement(key)
}

func Decrements(key string, count int64) {
	Cache.Store.Decrements(key, count)
}

func IsAlive() error {
	return Cache.Store.IsAlive()
}
