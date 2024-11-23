package model

import (
	"gorm.io/gorm"
)

type Bin struct {
	gorm.Model
	BinArea        Area `gorm:"constraint:OnDelete:CASCADE;foreignKey:AreaID" json:"bin_area"`
	AreaID         uint `json:"area_id"`
	WasteCollected int  `json:"waste_collected_in_kilogram"`
	BinOwner       User `gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID" json:"area_owned_by"`
	UserID         uint `json:"user_id"`
}
