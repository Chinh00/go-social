package postgres

import (
	"github.com/labstack/echo/v4"
	"go-social/internal/pkg/database"
	repo "go-social/internal/pkg/repository"
)

type repository[T any] struct {
	dataContext database.DataContext
}

func NewRepository[T any](dbContext database.DataContext) repo.Repository[T] {
	return &repository[T]{dataContext: dbContext}
}
func (r *repository[T]) FindAll(ctx echo.Context, condition interface{}) ([]*T, error) {
	results := make([]*T, 0)
	r.dataContext.GetDB().Find(&results, condition)
	return results, nil
}

func (r *repository[T]) AddEntity(ctx echo.Context, entity *T) (*T, error) {
	r.dataContext.GetDB().Create(entity)
	return entity, nil
}

func (r *repository[T]) UpdateEntity(ctx echo.Context, entity *T) (*T, error) {
	r.dataContext.GetDB().Save(entity)
	return entity, nil
}

func (r *repository[T]) RemoveEntity(ctx echo.Context, entity *T) (*T, error) {
	r.dataContext.GetDB().Delete(entity)
	return entity, nil
}

func (r *repository[T]) FindOne(ctx echo.Context, condition interface{}) (entity *T, err error) {
	r.dataContext.GetDB().Where(condition).First(entity)
	return
}
