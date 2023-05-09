package example

import "github.com/gofiber/fiber/v2"

// Example godoc
// @Summary
// @Schemes
// @Description Example
// @Tags example
// @Accept x-www-form-urlencoded
// @Produce json
// @Router /example [get]
func GET(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
