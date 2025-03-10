package route

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-social/internal/pkg/core"
	"go-social/internal/user/usecase/user"
)

var RouteSet = wire.NewSet(NewRoute)

type route struct {
	echo    *echo.Echo
	Usecase user.Usecase
}

func NewRoute(echo *echo.Echo, usecase user.Usecase) core.RouteRoot {
	return &route{
		echo:    echo,
		Usecase: usecase,
	}
}

// @title Social API
// @version 1.0
// @description API tài liệu cho Social app
// @host localhost:8081
// @BasePath /api/v1
func (r *route) MapEndpoint() {
	r.echo.GET("/swagger/*", echoSwagger.WrapHandler)
	_ = r.echo.Group("/api/v1")
}
