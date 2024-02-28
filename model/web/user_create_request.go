package web

type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `gorm:"type:varchar(50); not null" json:"password"`
	RoleID   int    `json:"role_id"`
}
