package repository

import (
	"ares/model"
	"ares/pkg/session"

	"gorm.io/gorm"
)

type Repository struct {
	session *session.Session
	Psql    *gorm.DB
}

type RepositoryInt interface {
	SetSession(session *session.Session)
	GetClient(username string) (model.Client, error)
	GetRuleByClientId() ([]model.Rule, error)
	GetActiveRules() (rules model.RuleSet, err error)
}

func New(db *gorm.DB) RepositoryInt {
	return &Repository{
		Psql: db,
	}
}

func (r *Repository) SetSession(s *session.Session) {
	r.session = s
}
