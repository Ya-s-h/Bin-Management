package model

import (
	"gorm.io/gorm"
)

type Bin struct {
	gorm.Model
	BinArea        Area `gorm:"foreignKey:AreaID" json:"bin_area"`
	AreaID         uint `json:"area_id"`
	WasteCollected int  `json:"waste_collected_in_gram"`
	BinOwner       User `gorm:"foreignKey:UserID" json:"area_owned_by"`
	UserID         uint `json:"user_id"`
}
