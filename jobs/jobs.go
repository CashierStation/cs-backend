package jobs

import (
	"log"

	"github.com/bamzi/jobrunner"
	"github.com/gofiber/fiber/v2"
)

type StartJobOptions struct {
	App *fiber.App
}

// Metrics godoc
// @Summary
// @Schemes
// @Description dev: http://localhost:8080/jobs
// @Description prod: http://csbackend.sivr.tech/jobs
// @Tags jobs
// @Accept x-www-form-urlencoded
// @Produce json
// @Router /jobs [get]
func StartJob(options StartJobOptions) {
	log.Println("Starting job scheduler...")

	jobrunner.Start()
	jobrunner.Schedule("@every 30s", CreateUpdateTarif())

	if options.App != nil {
		app := options.App

		app.Get("/jobs", func(c *fiber.Ctx) error {
			return c.Render("./views/Status.html", jobrunner.StatusPage())
		})
	}
}
