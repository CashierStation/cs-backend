package lib

import (
	"github.com/gofiber/fiber/v2"
)

func HTTPError(c *fiber.Ctx, code int, message string, err error) error {
	// var params = c.AllParams()
	// var queries = c.Queries()

	// rollbar.Error(errors.New(message), map[string]interface{}{
	// 	"params": params,
	// 	"query":  queries,
	// })

	return c.Status(code).JSON(fiber.Map{
		"status":  "error",
		"message": message,
	})
}
