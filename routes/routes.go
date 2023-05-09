package routes

import (
	"github.com/gofiber/fiber/v2"

	"csbackend/routes/auth"
	"csbackend/routes/example"
	"csbackend/routes/user"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/example", example.GET)
	app.Get("/user", user.GET)

	auth.Routes(app)
}
