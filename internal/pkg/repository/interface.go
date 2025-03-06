package repository

import (
	"github.com/labstack/echo/v4"
)

type Repository[T any] interface {
	AddEntity(ctx echo.Context, entity *T) (*T, error)
	UpdateEntity(ctx echo.Context, entity *T) (*T, error)
	RemoveEntity(ctx echo.Context, entity *T) (*T, error)
	FindOne(ctx echo.Context, condition interface{}) (*T, error)
	FindAll(ctx echo.Context, condition interface{}) ([]*T, error)
}
