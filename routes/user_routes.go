package route

import (
	controllers "assignment/renie/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserSetupRoutes(app *fiber.App) {
	app.Post("/api/create/user", controllers.CreateUser)
	app.Post("/api/update/user", controllers.UpdateUser)
}
