package redis

import (
	"context"
	"time"
)

type Service[T any] interface {
	Set(ctx context.Context, key string, value *T, expiration time.Duration) (*T, error)
	Get(ctx context.Context, key string) (*T, error)
	Remove(ctx context.Context, key string) (*T, error)
}
