package repository

import (
	"gorm.io/gorm"
)

type StartRepository interface {
	Awal(a string) string
}

type startRepository struct {
	connUser *gorm.DB
}

func NewStartRepository(connUser *gorm.DB) StartRepository {
	return &startRepository{
		connUser: connUser,
	}
}

func (s *startRepository) Awal(a string) string {
	return "Test"
}
