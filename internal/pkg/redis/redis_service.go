package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

//var _ Service = (*redis_service)(nil)

type redisService[T any] struct {
	client *redis.Client
}

func (r redisService[T]) Set(ctx context.Context, key string, value *T, expiration time.Duration) (*T, error) {
	val := r.client.Set(ctx, key, value, expiration)
	if val.Err() != nil {
		return nil, val.Err()
	}
	return value, nil
}

func (r redisService[T]) Get(ctx context.Context, key string) (*T, error) {
	//TODO implement me
	panic("implement me")
}

func (r redisService[T]) Remove(ctx context.Context, key string) (*T, error) {
	//TODO implement me
	panic("implement me")
}

func NewRedisService[T any](options *RedisOption) Service[T] {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", options.Host, options.Port),
	})

	return &redisService[T]{client: client}
}
