package controller

import (
	database "assignment/renie/db"
	models "assignment/renie/models"

	"github.com/gofiber/fiber/v2"
)

func AddWaste(fiber_context *fiber.Ctx) error {
	connection := database.ConnectToDb()
	new_waste := new(models.Waste)
	if err := fiber_context.BodyParser(new_waste); err != nil {
		return fiber_context.Status(400).JSON(fiber.Map{
			"error": "Invalid Format",
		})
	}
	result := connection.Create(&new_waste)
	if err := result.Error; err != nil {
		return err
	}
	return fiber_context.Status(200).JSON(new_waste)
}
