package model

import (
	"time"

	"gorm.io/gorm"
)

type Area struct {
	gorm.Model
	Name      string    `json:"area_name"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"timestamp"`
	AreaOwner User      `gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID" json:"area_owned_by"`
	UserID    uint      `json:"user_id"`
}
