package model

import (
	"time"
)

type Area struct {
	ID        uint      `json:"area_id"`
	Name      string    `json:"area_name"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"timestamp"`
	AreaOwner User      `gorm:"foreignKey:UserID" json:"area_owned_by"`
	UserID    uint      `json:"user_id"`
}
