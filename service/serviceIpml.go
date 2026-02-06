package service

import (
	"ares/helper/vo"
	"ares/pkg/grule"
	"ares/proto"
	"ares/service/core"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) Transaction(request *proto.Request) (resp *proto.Response, err error) {
	ruleSet, err := s.repo.GetActiveRules()
	if err != nil {
		s.session.LogError("failed get active rules", err.Error())
		return
	} else if len(ruleSet.Rules) == 0 {
		return
	}

	coreTrx := core.New(request, s.redis, s.repo)
	cr := grule.GetNewEngine()
	cr.SetIDataContext(coreTrx.GetKeyContext(), &coreTrx)

	for _, rule := range ruleSet.Rules {
		b, _ := base64.StdEncoding.DecodeString(rule.Drl_content)
		cr.SetKnowledgeBase(rule.RuleName, rule.Version, string(b))
	}

	// execute customGrule
	cr.ExecuteRule()
	resp = &proto.Response{
		TransactionId: request.TransactionId,
		Decision:      proto.Decision_APPROVE,
		ExecutionContext: &proto.ExecutionContext{
			RuleSetVersion: ruleSet.Version,
			ProcessedAt:    timestamppb.Now(),
			LatencyMs:      0,
		},
		Status: &proto.ResponseStatus{
			Code:         "00",
			Description:  "SUCCESS",
			InternalCode: 0,
		},
	}

	fmt.Println(coreTrx)

	return
}
func (s *Service) AuthValidation(auth string) (err error) {

	authData := vo.AuthDecode(auth)
	if !authData.IsValid {
		err = fmt.Errorf("Authorization Invalid")
		return
	}

	client, err := s.repo.GetClient(authData.Username)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(client.PasswordHash), []byte(authData.Password))
	if err != nil {
		return err
	}

	s.session.SetClientId(client.Id)
	return nil
}
