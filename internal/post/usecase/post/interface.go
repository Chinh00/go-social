package post

import (
	"github.com/labstack/echo/v4"
	"go-social/internal/post/domain"
	"go-social/internal/post/usecase/models"
)

type Usecase interface {
	GetPosts(ctx echo.Context) ([]*domain.Post, error)
	GetPostDetail() (*domain.Post, error)
	CreatePost(ctx echo.Context, post *models.PostCreate) (*domain.Post, error)
	UpdatePost() (*domain.Post, error)
	RemovePost() (*domain.Post, error)
}
