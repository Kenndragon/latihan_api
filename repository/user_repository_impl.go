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
	if err := db.Where("username = ?", username).
		Preload("Role").
		First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repository *UserRepositoryImpl) FindAll(db *gorm.DB) []domain.User {
	var users []domain.User
	if err := db.Preload("Role").Find(&users).Error; err != nil {
		helper.PanicError(err)
	}
	return users
}
