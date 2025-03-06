//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"go-social/cmd/post/config"
	"go-social/internal/pkg/database"
	"go-social/internal/post/app/route"
	"go-social/internal/post/usecase/post"
)

func Init(config *config.Config) (*App, error) {
	panic(
		wire.Build(
			echo.New,
			route.RouteSet,
			post.PostPgRepositorySet,
			database.PgContextSet,
			post.UsecaseSet,
			CatPostgresConfig,
			NewApp,
		))
}

func CatPostgresConfig(config *config.Config) *database.PostgresConfig {
	return &config.PostgresConfig
}
