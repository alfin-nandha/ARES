package repository

import (
	"ares/model"
	"ares/pkg/session"
)

func (r *Repository) GetClient(session *session.Session) (client model.Client, err error) {
	err = r.Psql.Where("id = ?", session.ClientId).First(&client).Error
	return
}
func (r *Repository) GetRuleByClientId(session *session.Session) (rules []model.Rule, err error) {
	err = r.Psql.Raw("select * from public.fraud_rules where client_id = ?", session.ClientId).Scan(&rules).Error
	return
}
