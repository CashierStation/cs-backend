package auth

import (
	"csbackend/routes/auth/callback"
	"csbackend/routes/auth/login"
	"csbackend/routes/auth/logout"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Get("/login", login.GET)
	auth.Get("/callback", callback.GET)
	auth.Get("/logout", logout.GET)
}
