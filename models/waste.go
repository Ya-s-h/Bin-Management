package model

import "gorm.io/gorm"

type Waste struct {
	gorm.Model
	CollectedAt string `json:"timestamp"`
	BinWaste    Bin    `gorm:"constraint:OnDelete:CASCADE;foreignKey:BinID"`
	BinID       uint   `json:"bin_id"`
	Weight      int    `json:"weight_in_kgs"`
	Earnings    int    `json:"earnings"`
}
