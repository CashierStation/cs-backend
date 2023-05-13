package example

import "github.com/gofiber/fiber/v2"

func GET(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
