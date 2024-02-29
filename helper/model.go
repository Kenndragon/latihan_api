package helper

import (
	"latihan_api/model/domain"
	"latihan_api/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		RoleID:   user.RoleID,
		RoleName: user.Role.Name,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var usersResponses []web.UserResponse
	for _, user := range users {
		usersResponses = append(usersResponses, ToUserResponse(user))
	}
	return usersResponses
}
