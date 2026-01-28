package repository

import (
	"ares/model"
	"ares/pkg/session"
)

func (r *Repository) GetClient(session *session.Session, username string) (client model.Client, err error) {
	err = r.Psql.WithContext(session.Ctx).Where("username = ?", username).First(&client).Error
	return
}
func (r *Repository) GetRuleByClientId(session *session.Session) (rules []model.Rule, err error) {
	err = r.Psql.WithContext(session.Ctx).Raw("select * from public.fraud_rules where client_id = ?", session.ClientId).Scan(&rules).Error
	return
}
