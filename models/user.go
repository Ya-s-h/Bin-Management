package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string   `json:"username"`
	Role     UserRole `gorm:"foreignKey:RoleID;" json:"role_details"`
	RoleID   uint     `json:"role_id"`
	Email    string   `json:"email_address"`
	Password string   `json:"-"`
}
