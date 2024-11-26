package controller

import (
	database "assignment/renie/db"
	models "assignment/renie/models"

	"github.com/gofiber/fiber/v2"
)

func CreateBin(fiber_context *fiber.Ctx) error {
	connection := database.ConnectToDb()
	new_bin := new(models.Bin)
	if err := fiber_context.BodyParser(new_bin); err != nil {
		return fiber_context.Status(400).JSON(fiber.Map{
			"error": "Invalid Format",
		})
	}
	if new_bin.BinOwner != (models.User{}) {
		if new_bin.BinOwner.Role.ID < 3 {
			return fiber_context.Status(400).JSON(fiber.Map{
				"error": "User Can't Own Bin",
			})
		}
	}
	new_bin.WasteCollected = 0
	result := connection.Create(&new_bin)
	if err := result.Error; err != nil {
		return fiber_context.Status(500).JSON(fiber.Map{
			"error": "Failed to create Bin",
		})
	}
	return fiber_context.Status(200).JSON(new_bin)
}

func DeleteBin(fiber_context *fiber.Ctx) error {
	payload := struct {
		BinID uint `json:"bin_id"`
	}{}
	if err := fiber_context.BodyParser(&payload); err != nil {
		return fiber_context.Status(404).JSON(fiber.Map{
			"error": err,
		})
	}
	id := payload.BinID
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

func AssignBinToArea(fiber_context *fiber.Ctx) error {
	payload := struct {
		BinID  uint `json:"bin_id"`
		AserID uint `json:"area_id"`
	}{}
	if err := fiber_context.BodyParser(&payload); err != nil {
		return fiber_context.Status(404).JSON(fiber.Map{
			"error": err,
		})
	}
	bin_id := payload.BinID
	areaID := payload.AserID
	var existingBin models.Bin
	db := database.ConnectToDb()

	// areaID, err := stringToUint(area_id_str)
	// if err != nil {
	// 	return fiber_context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": "Invalid area_id",
	// 	})
	// }
	var existingArea models.Area
	if err := db.First(&existingArea, "id = ?", areaID).Error; err != nil {
		return fiber_context.Status(404).JSON(fiber.Map{
			"error": "Invalid Area ID",
		})
	}

	if err := db.First(&existingBin, "id = ?", bin_id).Error; err != nil {
		return fiber_context.Status(404).JSON(fiber.Map{
			"error": "Bin not found",
		})
	}

	existingBin.AreaID = areaID
	existingBin.BinArea = existingArea
	if err := db.Save(&existingBin).Error; err != nil {
		return fiber_context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to assign bin to area",
		})
	}
	return fiber_context.Status(200).JSON(existingBin)
}

func AssignBinToUser(fiber_context *fiber.Ctx) error {
	payload := struct {
		BinID  uint `json:"bin_id"`
		UserID uint `json:"user_id"`
	}{}
	if err := fiber_context.BodyParser(&payload); err != nil {
		return fiber_context.Status(404).JSON(fiber.Map{
			"error": err,
		})
	}
	bin_id := payload.BinID
	userID := payload.UserID
	var existingBin models.Bin
	var existingUser models.User
	db := database.ConnectToDb()

	if err := db.First(&existingBin, "id = ?", bin_id).Error; err != nil {
		return fiber_context.Status(400).JSON(fiber.Map{
			"error": "Bin not found",
		})
	}
	if err := db.First(&existingUser, "id = ?", userID).Error; err != nil {
		return fiber_context.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if existingUser.RoleID < 3 {
		return fiber_context.Status(400).JSON(fiber.Map{
			"error": "User Can't Own Bin",
		})
	}
	existingBin.UserID = userID
	existingBin.BinOwner = existingUser
	if err := db.Save(&existingBin).Error; err != nil {
		return fiber_context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to assign bin to user",
		})
	}
	return fiber_context.Status(200).JSON(existingBin)
}
