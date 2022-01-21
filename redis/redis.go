// Package redis 工具包
package redis

import (
	"context"
	"github.com/curatorc/cngf/logger"
	"time"

	"github.com/go-redis/redis/v8"
)

// Client Redis 服务
type Client struct {
	Client  *redis.Client
	Context context.Context
}

// New 创建一个新的 redis 连接
func New(address string, username string, password string, db int) *Client {
	// 初始化自定义的 Client 实例
	rds := &Client{
		Context: context.Background(),
		Client: redis.NewClient(&redis.Options{
			Addr:     address,
			Username: username,
			Password: password,
			DB:       db,
		}),
	}

	// 测试一下连接
	err := rds.Ping()
	logger.LogIf(err)

	return rds
}

// Ping 用以测试 redis 连接是否正常
func (rds Client) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

// Set 存储 key 对应的 value, 且设置 expiration 过期时间
func (rds Client) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		logger.ErrorString("Redis", "set", err.Error())
		return false
	}
	return true
}

// Get 获取 key 对应的 value
func (rds Client) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return ""
	}
	return result
}

// Has 判断一个 key 是否存在，内部错误和 redis.Nil 都返回 false
func (rds Client) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return false
	}
	return true
}

// Del 删除存储在 redis 里的数据，支持多个 key 传参
func (rds Client) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}
	return true
}

// FlushDB 清空当前 redis db 里的所有数据
func (rds Client) FlushDB() bool {
	if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
		logger.ErrorString("Redis", "FlushDB", err.Error())
		return false
	}
	return true
}

func (rds Client) Increment(key string) bool {
	if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
		logger.ErrorString("Redis", "Increment", err.Error())
		return false
	}
	return true
}

func (rds Client) Increments(key string, count int64) bool {
	if err := rds.Client.IncrBy(rds.Context, key, count).Err(); err != nil {
		logger.ErrorString("Redis", "Increment", err.Error())
	}
	return true
}

func (rds Client) Decrement(key string) bool {
	if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
		logger.ErrorString("Redis", "Decrement", err.Error())
		return false
	}
	return true
}

func (rds Client) Decrements(key string, count int64) bool {
	if err := rds.Client.DecrBy(rds.Context, key, count).Err(); err != nil {
		logger.ErrorString("Redis", "Decrement", err.Error())
		return false
	}
	return true
}
