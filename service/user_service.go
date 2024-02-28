package service

import (
	"latihan_api/model/web"
)

type UserService interface {
	Register(request web.UserCreateRequest) web.UserResponse
	Login(request web.UserCreateRequest) web.UserResponse
}
