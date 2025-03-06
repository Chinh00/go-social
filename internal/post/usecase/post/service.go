package post

import (
	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"go-social/internal/pkg/database"
	"go-social/internal/pkg/repository"
	"go-social/internal/pkg/repository/postgres"
	"go-social/internal/post/domain"
	"go-social/internal/post/usecase/models"
	"time"
)

var UsecaseSet = wire.NewSet(NewService)

var PostPgRepositorySet = wire.NewSet(NewPgRepository)

func NewPgRepository(dbContext database.DataContext) repository.Repository[domain.Post] {
	return postgres.NewRepository[domain.Post](dbContext)
}

type service struct {
	repository repository.Repository[domain.Post]
}

func NewService(repository repository.Repository[domain.Post]) Usecase {
	return &service{
		repository: repository,
	}
}

func (s *service) GetPosts(ctx echo.Context) ([]*domain.Post, error) {
	posts, err := s.repository.FindAll(ctx, struct {
	}{})
	if err != nil {
		return nil, err
	}
	return posts, nil

}

func (s service) GetPostDetail() (*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) CreatePost(ctx echo.Context, post *models.PostCreate) (*domain.Post, error) {
	p, err := s.repository.AddEntity(ctx, &domain.Post{
		Title:    post.Title,
		AuthorId: post.AuthorId,
		TagId:    post.TagId,
		Content:  post.Content,
		PostId:   uuid.New(),
		CreateAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *service) UpdatePost() (*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) RemovePost() (*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}
