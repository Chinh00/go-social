package user

import (
	"github.com/google/wire"
	"go-social/internal/pkg/redis"
	"go-social/internal/user/domain"
)

type service struct {
	service redis.Service[domain.User]
}

func NewRedisServiceSet(option *redis.RedisOption) redis.Service[domain.User] {
	return redis.NewRedisService[domain.User](option)
}

var NewServiceSet = wire.NewSet(NewService)

func NewService(redisService redis.Service[domain.User]) Usecase {
	return service{
		service: redisService,
	}
}
func (s service) CreateUser() {

}
