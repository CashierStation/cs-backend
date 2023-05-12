package oauth

import (
	"csbackend/routes/oauth/callback"
	"csbackend/routes/oauth/login"
	"csbackend/routes/oauth/logout"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	oauth := app.Group("/oauth")
	oauth.Get("/login", login.GET)
	oauth.Get("/callback", callback.GET)
	oauth.Get("/logout", logout.GET)
}
