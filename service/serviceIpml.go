package service

import (
	"ares/helper/vo"
	"ares/model"
	"ares/pkg/grule"
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
	// get rule
	rule := `rule IsKaya "Apply orkay discount" salience 10 {
		when
			Trx.Get("amount") >= 1000
		then
			Trx.Set("category", "kaya");
			Trx.Set("serviceFee", 5.0);
			Complete();
		}`

	// parse the request
	trx := model.FromRequest(request)

	// initiate customGrule
	cr := grule.GetNewEngine()
	cr.SetIDataContext(trx.GetKeyContext(), &trx)
	cr.SetKnowledgeBase("IsKaya", "1", rule)

	// execute customGrule
	cr.ExecuteRule()

	fmt.Printf("Trx Amount: %v,Category: %v, Fee: $%v\n", trx.Get("amount"), trx.Get("category"), trx.Get("serviceFee"))
	return
}
