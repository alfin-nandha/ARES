package repository

import (
	"ares/model"
)

func (r *Repository) GetClient(username string) (client model.Client, err error) {
	err = r.Psql.WithContext(r.session.Ctx).Where("username = ?", username).First(&client).Error
	return
}
func (r *Repository) GetRuleByClientId() (rules []model.Rule, err error) {
	err = r.Psql.WithContext(r.session.Ctx).Raw("select * from public.fraud_rules where client_id = ?", r.session.ClientId).Scan(&rules).Error
	return
}
