package controller

import (
	database "assignment/renie/db"
	models "assignment/renie/models"

	"github.com/gofiber/fiber/v2"
)

func CreateArea(fiber_context *fiber.Ctx) error {
	connection := database.ConnectToDb()
	new_area := new(models.Area)
	if err := fiber_context.BodyParser(new_area); err != nil {
		return fiber_context.Status(400).JSON(fiber.Map{
			"error": "Invalid Format",
		})
	}
	invalid := false
	if new_area.AreaOwner != (models.User{}) {
		invalid = true
		for _, value := range []uint{2, 4} {
			if new_area.AreaOwner.RoleID == value {
				invalid = false
				break
			}
		}
	}

	if invalid {
		return fiber_context.Status(400).JSON(fiber.Map{
			"error": "User Can't Onw Area",
		})
	}
	result := connection.Create(&new_area)
	if err := result.Error; err != nil {
		return fiber_context.Status(500).JSON(fiber.Map{
			"error": "Failed to create Bin",
		})
	}
	return fiber_context.Status(200).JSON(new_area)
}

func DeleteArea(fiber_context *fiber.Ctx) error {
	payload := struct {
		AreaID uint `json:"area_id"`
	}{}
	if err := fiber_context.BodyParser(&payload); err != nil {
		return fiber_context.Status(404).JSON(fiber.Map{
			"error": err,
		})
	}
	id := payload.AreaID
	db := database.ConnectToDb()

	var existingBin models.Bin
	if err := db.First(&existingBin, "id = ?", id).Error; err != nil {
		return fiber_context.Status(404).JSON(fiber.Map{
			"error": "Bin not found",
		})
	}

	// Update the record with new data
	if err := db.Delete(&existingBin).Error; err != nil {
		return fiber_context.Status(500).JSON(fiber.Map{
			"error": "Failed to delete Bin",
		})
	}
	return fiber_context.Status(200).JSON(existingBin)
}

func AssingAreaToUser(fiber_context *fiber.Ctx) error {
	payload := struct {
		UserID uint `json:"user_id"`
		AserID uint `json:"area_id"`
	}{}
	if err := fiber_context.BodyParser(&payload); err != nil {
		return fiber_context.Status(404).JSON(fiber.Map{
			"error": err,
		})
	}
	user_id := payload.UserID
	areaID := payload.AserID
	var existingArea models.Area
	var existingUser models.User
	db := database.ConnectToDb()

	if err := db.First(&existingArea, "id = ?", areaID).Error; err != nil {
		return fiber_context.Status(404).JSON(fiber.Map{
			"error": "Invalid Area ID",
		})
	}

	if err := db.First(&existingUser, "id = ?", user_id).Error; err != nil {
		return fiber_context.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	invalid := false
	if existingUser.Role != (models.UserRole{}) {
		invalid = true
		for _, value := range []uint{2, 4} {
			if existingUser.Role.ID == value {
				invalid = false
				break
			}
		}
	}
	if invalid {
		return fiber_context.Status(400).JSON(fiber.Map{
			"error": "User Can't Onw Area",
		})
	}
	existingArea.UserID = user_id
	existingArea.AreaOwner = existingUser

	if err := db.Save(&existingArea).Error; err != nil {
		return fiber_context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to assign bin to user",
		})
	}
	return fiber_context.Status(200).JSON(existingArea)
}
