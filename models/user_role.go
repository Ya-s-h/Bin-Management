package model

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	Name string `json:"role_name"`
}
