package controller

import (
	database "assignment/renie/db"
	init_db "assignment/renie/db_init"

	"github.com/gofiber/fiber/v2"
)

func DB_init(fiber_context *fiber.Ctx) error {
	connection := database.ConnectToDb()
	init_db.DeleteTables(connection)
	init_db.CreateTables(connection)
	init_db.CreateProcedures(connection)
	init_db.SeedTables(connection)
	// init_db.InsertRolesData(connection)
	// init_db.InsertUserData(connection)

	return fiber_context.Status(200).JSON(fiber.Map{
		"message": "Db Initialize",
	})

}
