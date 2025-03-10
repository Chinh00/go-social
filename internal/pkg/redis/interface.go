package redis

import "time"

type RedisService[T any] interface {
	Set(key string, value *T, expiration time.Duration)
	Get(key string) (*T, error)
}
