package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	post "go-social/cmd/post/config"
	"go-social/internal/pkg/core"
	"go-social/internal/pkg/database"
	"go-social/internal/pkg/repository"
	"go-social/internal/post/domain"
	uc "go-social/internal/post/usecase/post"
)

type App struct {
	database.DataContext
	echo *echo.Echo
	core.RouteRoot
	*post.Config
	postRepo repository.Repository[domain.Post]
	usecase  uc.Usecase
}

func (app *App) Run() error {
	err := app.DataContext.Migrations(&domain.Tag{}, &domain.Post{})
	if err != nil {
		return err
	}

	app.RouteRoot.MapEndpoint()
	err = app.echo.Start(fmt.Sprintf("%s:%d", app.Http.Host, app.Http.Port))
	if err != nil {
		return err
	}
	return nil
}

func NewApp(config *post.Config, echo *echo.Echo, root core.RouteRoot, ctx database.DataContext, repository2 repository.Repository[domain.Post], usecase uc.Usecase) *App {
	return &App{
		Config:      config,
		echo:        echo,
		RouteRoot:   root,
		DataContext: ctx,
		postRepo:    repository2,
		usecase:     usecase,
	}
}
