package web

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	RoleID   int    `json:"role_id"`
	RoleName string `json:"role_name"`
}
