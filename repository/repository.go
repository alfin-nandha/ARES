package repository

import "gorm.io/gorm"

type Repository struct {
	DbClient *gorm.DB
}

type RepositoryInt interface {
}

func New(db *gorm.DB) RepositoryInt {
	return &Repository{
		DbClient: db,
	}
}
