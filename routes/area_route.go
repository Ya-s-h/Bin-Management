package route

import (
	controllers "assignment/renie/controllers"

	"github.com/gofiber/fiber/v2"
)

func AreaSetupRoutes(app *fiber.App) {
	app.Post("/api/create/area", controllers.CreateArea)
	app.Post("/api/delete/area", controllers.DeleteArea)
	app.Post("/api/assign_area/bin", controllers.AssingAreaToUser)
}
