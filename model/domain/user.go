package domain

type Role struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"size:50"`
}

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	Password string `gorm:"type:varchar(255)"`
	RoleID   int
	Role     Role `gorm:"foreignKey:RoleID"`
}
