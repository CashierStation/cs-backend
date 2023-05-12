package auth

import (
	"csbackend/routes/auth/register"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", register.POST)
}
