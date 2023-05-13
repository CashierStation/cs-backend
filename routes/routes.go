package routes

import (
	"github.com/gofiber/fiber/v2"

	"csbackend/routes/auth"
	m "csbackend/routes/metrics"
	"csbackend/routes/oauth"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html", fiber.StatusTemporaryRedirect)
	})
	//app.Get("/example", example.GET)

	_ = app.Group("/api")
	//apiGroup.Get("/user", user.GET)

	oauth.Routes(app)
	auth.Routes(app)
	m.Routes(app)
}
