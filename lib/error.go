package lib

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rollbar/rollbar-go"
)

func HTTPError(c *fiber.Ctx, code int, message string, err error) error {
	var params = c.AllParams()
	var queries = c.Queries()

	rollbar.Error(errors.New(message), map[string]interface{}{
		"params": params,
		"query":  queries,
	})

	return c.Status(code).JSON(fiber.Map{
		"status":  "error",
		"message": message,
	})
}
