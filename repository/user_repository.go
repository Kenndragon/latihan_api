package repository

import (
	"latihan_api/model/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(db *gorm.DB, user domain.User) domain.User
	Login(db *gorm.DB, username string) (domain.User, error)
	FindAll(db *gorm.DB) []domain.User
}
