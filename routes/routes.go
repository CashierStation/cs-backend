package routes

import (
	"github.com/gofiber/fiber/v2"

	"csbackend/routes/auth"
	"csbackend/routes/example"
	m "csbackend/routes/metrics"
	"csbackend/routes/user"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html", fiber.StatusTemporaryRedirect)
	})
	app.Get("/example", example.GET)
	app.Get("/user", user.GET)

	auth.Routes(app)
	m.Routes(app)
}
