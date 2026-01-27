package repository

import (
	"ares/model"
	"ares/pkg/session"

	"gorm.io/gorm"
)

type Repository struct {
	Psql *gorm.DB
}

type RepositoryInt interface {
	GetClient(session *session.Session) (model.Client, error)
	GetRuleByClientId(session *session.Session) ([]model.Rule, error)
}

func New(db *gorm.DB) RepositoryInt {
	return &Repository{
		Psql: db,
	}
}
