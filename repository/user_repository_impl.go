package repository

import (
	"latihan_api/helper"
	"latihan_api/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Register(db *gorm.DB, user domain.User) domain.User {
	err := db.Create(&user).Error
	helper.PanicError(err)

	return user
}

func (repository *UserRepositoryImpl) Login(db *gorm.DB, username string) (domain.User, error) {
	var user domain.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
