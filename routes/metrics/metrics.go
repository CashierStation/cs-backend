package metrics

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// Metrics godoc
// @Summary
// @Schemes
// @Description dev: http://localhost:8080/metrics
// @Description prod: https://cs-backend-production.up.railway.app/metrics
// @Tags metrics
// @Accept x-www-form-urlencoded
// @Produce json
// @Router /metrics [get]
func Routes(app *fiber.App) {
	app.Get("/metrics", monitor.New())
}
