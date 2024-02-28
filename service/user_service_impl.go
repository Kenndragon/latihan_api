package service

import (
	"latihan_api/helper"
	"latihan_api/model/domain"
	"latihan_api/model/web"
	"latihan_api/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, db *gorm.DB, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Register(request web.UserCreateRequest) web.UserResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	hashedPassword, err := helper.HashPassword(request.Password)
	helper.PanicError(err)

	user := domain.User{
		Username: request.Username,
		Password: hashedPassword,
		RoleID:   request.RoleID,
	}
	user = service.UserRepository.Register(tx, user)
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Login(request web.UserCreateRequest) web.UserResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	user, err := service.UserRepository.Login(tx, request.Username)
	if err != nil {
		return web.UserResponse{}
	}
	check := helper.CheckPasswordHash(request.Password, user.Password)
	if !check {
		return web.UserResponse{}
	}
	return helper.ToUserResponse(user)
}
