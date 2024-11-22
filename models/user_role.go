package model

type UserRole struct {
	ID   uint   `gorm:"primary_key" json:"role_id"`
	Name string `json:"role_name"`
}
