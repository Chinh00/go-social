//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"go-social/cmd/user/config"
	"go-social/internal/pkg/redis"
	"go-social/internal/user/app/route"
	"go-social/internal/user/usecase/user"
)

func Init(config *config.Config) (*App, error) {
	panic(wire.Build(
		echo.New,
		user.NewRedisServiceSet,
		route.RouteSet,
		CatRedisOption,
		user.NewServiceSet,
		NewApp,
	))
}

func CatRedisOption(config *config.Config) *redis.RedisOption {
	return &config.RedisOption
}
