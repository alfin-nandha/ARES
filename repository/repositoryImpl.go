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

func (r *Repository) GetActiveRules() (rules model.RuleSet, err error) {
	err = r.Psql.WithContext(r.session.Ctx).
		Joins("JOIN fraud_rulses on rule_sets.id = fraud_rules.rule_set_id").
		Preload("Rules", "is_active = ?", true).
		Where("rule_sets.client_id = ? and rule_sets.is_active = true and fraud_rules.is_active = true", r.session.ClientId).
		Find(&rules).Error
	return
}
