package user

import (
	"csbackend/global"

	"github.com/gofiber/fiber/v2"
)

// User godoc
// @Summary
// @Schemes
// @Description User
// @Tags user
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} user.GET.response
// @Router /user [get]
func GET(c *fiber.Ctx) error {
	type response struct {
		Profile interface{} `json:"profile"`
	}

	sess, err := global.Session.Get(c)
	if err != nil {
		return err
	}

	profile := sess.Get("profile")

	return c.JSON(response{profile})
}
