package route

import (
	controllers "assignment/renie/controllers"

	"github.com/gofiber/fiber/v2"
)

func WasteSetupRoutes(app *fiber.App) {
	app.Post("/api/add/waste", controllers.AddWaste)
}
