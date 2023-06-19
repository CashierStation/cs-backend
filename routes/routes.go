package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"csbackend/authenticator"
	_ "csbackend/routes/api/employee"
	"csbackend/routes/api/user"
	"csbackend/routes/auth"
	m "csbackend/routes/metrics"
	"csbackend/routes/oauth"

	"csbackend/routes/api/booking"
	"csbackend/routes/api/snack"
	"csbackend/routes/api/unit"
	"csbackend/routes/api/unitsession"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html", fiber.StatusTemporaryRedirect)
	})
	//app.Get("/example", example.GET)

	apiGroup := app.Group("/api")
	apiGroup.Use(authenticator.SessionMiddleware(db))
	apiGroup.Get("/user", user.GET)

	// apiGroup.Get("/employee/list", employee.GetEmployeeList) temp

	apiGroup.Get("/unit", unit.GET)
	apiGroup.Post("/unit", unit.POST)
	apiGroup.Put("/unit/:id", unit.PUT)
	apiGroup.Delete("/unit/:id", unit.DELETE)

	apiGroup.Get("/snack", snack.GET)
	apiGroup.Post("/snack", snack.POST)
	apiGroup.Put("/snack/:id", snack.PUT)
	apiGroup.Delete("/snack/:id", snack.DELETE)
	apiGroup.Post("/snack/transaction", snack.CreateSnackTransaction)
	apiGroup.Post("/snack/restock", snack.CreateSnackRestock)

	// TODO: GET /snack/transaction and GET /snack/restock

	apiGroup.Get("/unit_session", unitsession.GetUnitSessions)
	apiGroup.Put("/unit_session/start/:unit_id", unitsession.StartUnitSessions)
	apiGroup.Put("/unit_session/stop/:unit_id", unitsession.StopUnitSessions)

	apiGroup.Get("/booking", booking.GetBookingList)
	apiGroup.Post("/booking", booking.CreateBooking)
	apiGroup.Put("/booking/:id", booking.UpdateBooking)

	oauth.Routes(app)
	auth.Routes(app)
	m.Routes(app)
}
