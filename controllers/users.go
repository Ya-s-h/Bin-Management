package controller

import (
	database "assignment/renie/db"
	models "assignment/renie/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(fiber_context *fiber.Ctx) error {
	user := new(models.User)
	if err := fiber_context.BodyParser(user); err != nil {
		return fiber_context.Status(400).JSON(fiber.Map{
			"error": "Invalid Format",
		})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber_context.Status(500).JSON(fiber.Map{
			"error": "Error hashing password",
		})
	}
	user.Password = string(hashedPassword)

	result := database.ConnectToDb().Create(user)
	if err := result.Error; err != nil {
		return err
	}

	return fiber_context.Status(200).JSON(user)

}
