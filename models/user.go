package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primary_key" json:"user_id"`
	Name      string         `json:"username"`
	Role      UserRole       `gorm:"foreignKey:RoleID;" json:"role_details"`
	RoleID    uint           `json:"role_id"`
	Email     string         `json:"email_address"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"last_updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
