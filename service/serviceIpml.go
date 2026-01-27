package service

import (
	"ares/pkg/session"
	"ares/proto"
)

func (s *Service) Transaction(session *session.Session, request *proto.Request) (response *proto.Response, err error) {

	client, err := s.repo.GetClient(session)
	if err != nil {
		session.LogError("failed get client from db", err)
		return
	}

	session.LogInfo(client)

	return
}
