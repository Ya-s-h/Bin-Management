package controller

import (
	database "assignment/renie/db"
	models "assignment/renie/models"
	"strconv"

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
	result := connection.Create(&new_bin)
	if err := result.Error; err != nil {
		return err
	}
	return fiber_context.Status(200).JSON(new_bin)
}

func DeleteBin(fiber_context *fiber.Ctx) error {
	id := fiber_context.Query("id")
	new_bin := new(models.Bin)
	db := database.ConnectToDb()
	if err := fiber_context.BodyParser(new_bin); err != nil {
		return fiber_context.Status(400).JSON(fiber.Map{
			"error": "Invalid Format",
		})
	}
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
	// fmt.Println(existingBin)
	return fiber_context.Status(200).JSON(existingBin)
}

func stringToUint(str string) (uint, error) {
	// Parse the string as a 64-bit integer
	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	// Convert to uint and return
	return uint(val), nil
}

func AssignBinToArea(fiber_context *fiber.Ctx) error {
	bin_id := fiber_context.Query("bin_id")
	area_id_str := fiber_context.Query("area_id")
	var existingBin models.Bin
	db := database.ConnectToDb()

	areaID, err := stringToUint(area_id_str)
	if err != nil {
		return fiber_context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid area_id",
		})
	}

	if err := db.First(&existingBin, "id = ?", bin_id).Error; err != nil {
		return fiber_context.Status(404).JSON(fiber.Map{
			"error": "Bin not found",
		})
	}

	existingBin.AreaID = areaID
	if err := db.Save(&existingBin).Error; err != nil {
		return fiber_context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to assign bin to area",
		})
	}
	return fiber_context.Status(200).JSON(existingBin)
}

func AssignBinToUser(fiber_context *fiber.Ctx) error {
	bin_id := fiber_context.Query("bin_id")
	user_id_str := fiber_context.Query("area_id")
	var existingBin models.Bin
	db := database.ConnectToDb()

	userID, err := stringToUint(user_id_str)
	if err != nil {
		return fiber_context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid area_id",
		})
	}

	if err := db.First(&existingBin, "id = ?", bin_id).Error; err != nil {
		return fiber_context.Status(404).JSON(fiber.Map{
			"error": "Bin not found",
		})
	}

	existingBin.UserID = userID
	if err := db.Save(&existingBin).Error; err != nil {
		return fiber_context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to assign bin to area",
		})
	}
	return fiber_context.Status(200).JSON(existingBin)
}
