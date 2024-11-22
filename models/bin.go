package model

import (
	"time"

	"gorm.io/gorm"
)

type Bin struct {
	ID             uint           `json:"bin_id"`
	BinArea        Area           `gorm:"foreignKey:AreaID" json:"bin_area"`
	AreaID         uint           `json:"area_id"`
	WasteCollected int            `json:"waste_collected_in_gram"`
	BinOwner       User           `gorm:"foreignKey:UserID" json:"area_owned_by"`
	UserID         uint           `json:"user_id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"last_updated"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
