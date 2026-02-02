package service

import (
	"ares/pkg/redis"
	"ares/pkg/session"
	"ares/proto"
	"ares/repository"
)

type Service struct {
	session *session.Session
	repo    repository.RepositoryInt
	redis   redis.RedisInt
}

type ServiceInt interface {
	SetSession(session *session.Session)
	Transaction(request *proto.Request) (*proto.Response, error)
	AuthValidation(auth string) error
}

func New(repo repository.RepositoryInt, redis redis.RedisInt) ServiceInt {
	return &Service{
		repo:  repo,
		redis: redis,
	}
}

func (s *Service) SetSession(session *session.Session) {
	s.session = session
}
