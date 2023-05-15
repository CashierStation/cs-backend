package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"csbackend/authenticator"
	"csbackend/routes/api/user"
	"csbackend/routes/auth"
	m "csbackend/routes/metrics"
	"csbackend/routes/oauth"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html", fiber.StatusTemporaryRedirect)
	})
	//app.Get("/example", example.GET)

	apiGroup := app.Group("/api")
	apiGroup.Use(authenticator.SessionMiddleware(db))
	apiGroup.Get("/user", user.GET)

	oauth.Routes(app)
	auth.Routes(app)
	m.Routes(app)
}
