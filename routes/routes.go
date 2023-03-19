package routes

import (
	"carapi/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)

}
