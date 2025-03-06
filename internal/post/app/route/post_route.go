package route

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "go-social/docs"
	"go-social/internal/pkg/core"
	"go-social/internal/post/usecase/models"
	"go-social/internal/post/usecase/post"
	"net/http"
)

var RouteSet = wire.NewSet(NewRoute)

type route struct {
	echo *echo.Echo
	post.Usecase
}

func NewRoute(echo *echo.Echo, usecase post.Usecase) core.RouteRoot {
	return &route{echo: echo, Usecase: usecase}
}

// @title Social API
// @version 1.0
// @description API tài liệu cho Social app
// @host localhost:8080
// @BasePath /api/v1
func (r *route) MapEndpoint() {
	r.echo.GET("/swagger/*", echoSwagger.WrapHandler)
	group := r.echo.Group("/api/v1")
	group.GET("/posts", r.HandlerGetPosts)
	group.POST("/posts", r.HandlerCreatePost)
}

// HandlerGetPosts godoc
// @Summary Lấy danh sách posts
// @Description Trả về danh sách bài viết
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {object} domain.Post
// @Router /posts [get]
func (r *route) HandlerGetPosts(ctx echo.Context) error {
	posts, err := r.Usecase.GetPosts(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, posts)
}

// HandlerCreatePost godoc
// @Summary Tạo mới post
// @Description Trả về post created
// @Tags posts
// @Accept json
// @Param post body models.PostCreate true "Dữ liệu bài viết"
// @Produce json
// @Success 200 {object} domain.Post
// @Failure 400 {object} map[string]string
// @Router /posts [post]
func (r *route) HandlerCreatePost(ctx echo.Context) error {
	p := &models.PostCreate{}
	err := ctx.Bind(p)
	if err != nil {
		return err
	}
	createPost, err := r.Usecase.CreatePost(ctx, p)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, createPost)
}
