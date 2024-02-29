package service

import (
	"latihan_api/model/web"
)

type UserService interface {
	Register(request web.UserCreateRequest) web.UserResponse
	Login(request web.UserLoginRequest) (web.UserResponse, error)
	FindAll() []web.UserResponse
}
