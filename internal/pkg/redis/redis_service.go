package redis

import (
	"github.com/redis/go-redis/v9"
	"time"
)

//var _ RedisService = (*redis_service)(nil)

type redis_service[T any] struct {
	client *redis.Client
}

func NewRedisService[T any](client *redis.Client) RedisService[T] {
	return &redis_service[T]{client: client}
}

func (r redis_service[T]) Set(key string, value *T, expiration time.Duration) {
	//TODO implement me
	panic("implement me")
}

func (r redis_service[T]) Get(key string) (*T, error) {
	//TODO implement me
	panic("implement me")
}
