package main

import (
	pg "assignment/renie/db"
	route "assignment/renie/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/live",
	}))

	pg.ConnectToDb()
	route.UserSetupRoutes(app)
	log.Fatal(app.Listen(":3000"))

}
