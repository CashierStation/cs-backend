package metrics

import (
	"csbackend/global"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// Metrics godoc
// @Summary
// @Schemes
// @Description dev: http://localhost:8080/metrics
// @Description prod: https://csbackend.fly.dev/metrics
// @Tags metrics
// @Accept x-www-form-urlencoded
// @Produce json
// @Router /metrics [get]
func Routes(app *fiber.App) {
	app.Get("/metrics", monitor.New())
	app.Get("/metrics/db", DatabasePing)
}

// too lazy to create a separate file for this

// Metrics godoc
// @Summary
// @Schemes
// @Description dev: http://localhost:8080/metrics
// @Description prod: https://csbackend.fly.dev/metrics
// @Tags metrics
// @Accept x-www-form-urlencoded
// @Produce json
// @Router /metrics/db [get]
func DatabasePing(c *fiber.Ctx) error {
	results := []int64{}

	// ping 10 times with 100 ms delay
	for i := 0; i < 10; i++ {
		start := time.Now()

		tx := global.DB.Begin()
		tx.Exec("SELECT 1")
		tx.Commit()

		elapsed := time.Since(start)

		results = append(results, elapsed.Milliseconds())
		time.Sleep(100 * time.Millisecond)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": results,
	})
}
