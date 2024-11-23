package controller

import (
	database "assignment/renie/db"
	models "assignment/renie/models"
	validation "assignment/renie/validations"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(fiber_context *fiber.Ctx) error {
	db := database.ConnectToDb()
	user := new(models.User)
	if err := fiber_context.BodyParser(user); err != nil {
		return fiber_context.Status(400).JSON(fiber.Map{
			"error": "Invalid Format",
		})
	}
	validation_error := validation.DuplicateEmailAddress(user)
	if len(validation_error) > 0 && validation_error[0].Error {
		return fiber_context.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("%s already exists: %s",
				validation_error[0].FailedField,
				validation_error[0].FailedFieldValue,
			),
		})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber_context.Status(500).JSON(fiber.Map{
			"error": "Error hashing password",
		})
	}
	user.Password = string(hashedPassword)

	result := db.Create(user)
	if err := result.Error; err != nil {
		return err
	}

	return fiber_context.Status(200).JSON(user)

}

func UpdateUser(fiber_context *fiber.Ctx) error {
	id := fiber_context.Query("id")
	fmt.Println("Here")
	user := new(models.User)
	db := database.ConnectToDb()
	if err := fiber_context.BodyParser(user); err != nil {
		return fiber_context.Status(400).JSON(fiber.Map{
			"error": "Invalid Format",
		})
	}
	var existingUser models.User
	if err := db.First(&existingUser, "id = ?", id).Error; err != nil {
		return fiber_context.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Update the record with new data
	if err := db.Model(&existingUser).Updates(user).Error; err != nil {
		return fiber_context.Status(500).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}
	// fmt.Println(existingUser)
	return fiber_context.Status(200).JSON(existingUser)
}

func DeleteUser(fiber_context *fiber.Ctx) error {
	id := fiber_context.Query("id")
	user := new(models.User)
	db := database.ConnectToDb()
	if err := fiber_context.BodyParser(user); err != nil {
		return fiber_context.Status(400).JSON(fiber.Map{
			"error": "Invalid Format",
		})
	}
	var existingUser models.User
	if err := db.First(&existingUser, "id = ?", id).Error; err != nil {
		return fiber_context.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Update the record with new data
	if err := db.Delete(&existingUser).Error; err != nil {
		return fiber_context.Status(500).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}
	// fmt.Println(existingUser)
	return fiber_context.Status(200).JSON(existingUser)
}
