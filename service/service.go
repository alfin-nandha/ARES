package service

import (
	"ares/pkg/redis"
	"ares/repository"
)

type Service struct {
	repo  repository.RepositoryInt
	redis redis.RedisInt
}

type ServiceInt interface {
}

func New(repo repository.RepositoryInt, redis redis.RedisInt) ServiceInt {
	return &Service{
		repo:  repo,
		redis: redis,
	}
}
