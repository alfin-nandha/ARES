package service

import (
	"ares/pkg/redis"
	"ares/pkg/session"
	"ares/proto"
	"ares/repository"
)

type Service struct {
	repo  repository.RepositoryInt
	redis redis.RedisInt
}

type ServiceInt interface {
	Transaction(session *session.Session, request *proto.Request) (*proto.Response, error)
}

func New(repo repository.RepositoryInt, redis redis.RedisInt) ServiceInt {
	return &Service{
		repo:  repo,
		redis: redis,
	}
}
