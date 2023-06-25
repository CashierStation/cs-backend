package auth

import (
	changepassword "csbackend/routes/auth/changePassword"
	"csbackend/routes/auth/login"
	"csbackend/routes/auth/register"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", register.POST)
	auth.Post("/login", login.POST)
	auth.Post("/change_password", changepassword.POST)
}
