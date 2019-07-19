package store

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/patrickmn/go-cache"
)

type Cache interface {
	Get()
	Set()
	Delete()
}

func NewRedisCache() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}

type RedisCache struct {
	client *redis.Client
}

func (c *RedisCache) Get() {
	c.client.Get("dog")
}

func (c *RedisCache) Set() {
	c.client.Set("dog2", "dog3", 0)
}

type InMemoryCache struct {
	cache *cache.Cache
}

func (c *InMemoryCache) Get(key string) {
	c.cache.Get(key)
}
func (c *InMemoryCache) Set(key string, val interface{}) {
	c.cache.Set(key, val, cache.DefaultExpiration)
}
func (c *InMemoryCache) Delete(key string) {
	c.cache.Delete(key)
}

func NewInMemoryCache() *InMemoryCache {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &InMemoryCache{
		cache: c,
	}
}
