package main

import (
	pg "assignment/renie/db"
	route "assignment/renie/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	pg.ConnectToDb()
	route.UserSetupRoutes(app)
	log.Fatal(app.Listen(":3000"))

}
