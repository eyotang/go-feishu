package feishu

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/patrickmn/go-cache"
)

const (
	defaultExpire  = 5 * time.Minute
	defaultCleanup = 10 * time.Minute
)

type TokenCache interface {
	Set(string, interface{}, time.Duration)
	Get(string) (interface{}, bool)
	Lock()
	Unlock()
}

type RedisCache struct {
	client *redis.Client
	sync.Mutex
}

func (rc *RedisCache) Set(key string, value interface{}, expire time.Duration) {
	_ = rc.client.Set(context.Background(), key, value, expire).Err()
}

func (rc *RedisCache) Get(key string) (interface{}, bool) {
	var (
		err   error
		value string
	)
	if value, err = rc.client.Get(context.Background(), key).Result(); err != nil {
		return "", false
	}
	return value, true
}

type LocalCache struct {
	*cache.Cache
	sync.Mutex
}

var (
	tokenCache *LocalCache
	once       sync.Once
)

// 全局唯一，保证多个client操作同一个appId，也不会各自刷token。
func LocalTokenCache() *LocalCache {
	once.Do(func() {
		tokenCache = &LocalCache{
			Cache: cache.New(defaultExpire, defaultCleanup),
		}
	})
	return tokenCache
}

type CacheOptionFunc func(*accessTokenManagerService)

// 使用同一个缓存
func WithLocalCache() CacheOptionFunc {
	return func(s *accessTokenManagerService) {
		s.Cache = LocalTokenCache()
	}
}

func WithRedisCache(addr, password string) CacheOptionFunc {
	return func(s *accessTokenManagerService) {
		s.Cache = &RedisCache{
			client: redis.NewClient(&redis.Options{
				Addr:     addr,
				Password: password,
				DB:       0,
			}),
		}
	}
}
