package main

import (
	connection "assignment/renie/db"
	model "assignment/renie/models"
	"log"

	"gorm.io/gorm"
)

// db.Model(&models.UserInfo{}).AddForeignKey("u_id", "t_user(id)", "RESTRICT", "RESTRICT")

func CreateTables(db *gorm.DB) {
	error := db.AutoMigrate(&model.UserRole{}, &model.User{}, &model.Area{}, &model.Bin{}, &model.Waste{})

	if error != nil {
		log.Fatalf("Failed to migrate the database1: %v", error)
	}

}

func main() {
	db := connection.ConnectToDb()
	// Create Tables
	CreateTables(db)
	// Create Roles
	userRoles := []model.UserRole{
		{Name: "Area Provider"},
		{Name: "Owners/Investors"},
		{Name: "Both"},
	}
	result := db.Create(&userRoles)
	if err := result.Error; err != nil {
		log.Fatalf("Error while creating: %v\n", err)
	}
}
