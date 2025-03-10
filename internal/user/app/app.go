package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-social/cmd/user/config"
	"go-social/internal/pkg/core"
	"go-social/internal/pkg/redis"
	"go-social/internal/user/domain"
	"go-social/internal/user/usecase/user"
)

type App struct {
	*config.Config
	echo *echo.Echo
	user.Usecase
	redis.Service[domain.User]
	core.RouteRoot
}

func (a *App) Run() error {
	err := a.echo.Start(fmt.Sprintf("%s:%d", a.Http.Host, a.Http.Port))
	if err != nil {
		return err
	}
	return nil
}

func NewApp(config *config.Config, usecase user.Usecase, echo *echo.Echo, service redis.Service[domain.User], root core.RouteRoot) (*App, error) {
	return &App{
		Config:    config,
		echo:      echo,
		Usecase:   usecase,
		Service:   service,
		RouteRoot: root,
	}, nil
}
