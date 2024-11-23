package route

import (
	controllers "assignment/renie/controllers"

	"github.com/gofiber/fiber/v2"
)

func BinSetupRoutes(app *fiber.App) {
	app.Post("/api/create/bin", controllers.CreateBin)
	app.Post("/api/delete/bin", controllers.DeleteBin)
	app.Post("/api/assign_bin/area", controllers.AssignBinToArea)
	app.Post("/api/assign_bin/user", controllers.AssignBinToUser)
}
