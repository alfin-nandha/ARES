package service

import (
	"ares/helper/vo"
	"ares/pkg/session"
	"ares/proto"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) AuthValidation(session *session.Session, auth string) (err error) {

	authData := vo.AuthDecode(auth)
	if !authData.IsValid {
		err = fmt.Errorf("Authorization Invalid")
		return
	}

	client, err := s.repo.GetClient(session, authData.Username)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(client.PasswordHash), []byte(authData.Password))
	if err != nil {
		return err
	}

	session.SetClientId(client.Id)
	return nil
}
func (s *Service) Transaction(session *session.Session, request *proto.Request) (response *proto.Response, err error) {

	return
}
