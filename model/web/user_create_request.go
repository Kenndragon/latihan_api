package web

type UserCreateRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" gorm:"type:varchar(255); not null" json:"password"`
	RoleID   int    `json:"role_id"`
}
