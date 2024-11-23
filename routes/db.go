package route

import (
	controllers "assignment/renie/controllers"

	"github.com/gofiber/fiber/v2"
)

func DbSetupRoutes(app *fiber.App) {
	app.Post("/api/init/db", controllers.DB_init)
}
