package user

import (
	"csbackend/models"

	"github.com/gofiber/fiber/v2"
)

// @Security SessionToken
// User godoc
// @Summary
// @Schemes
// @Description User
// @Tags user
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} user.GET.response
// @Router /api/user [get]
func GET(c *fiber.Ctx) error {
	type response struct {
		ID        string `json:"id"`
		RentalID  string `json:"rental_id"`
		Username  string `json:"username"`
		Role      string `json:"role"`
		CreatedAt string `json:"created_at"`
	}

	user := c.Locals("user").(models.Employee)

	resp := response{
		ID:        user.ID,
		RentalID:  user.RentalID,
		Username:  user.Username,
		Role:      user.Role.Name,
		CreatedAt: user.CreatedAt.String(),
	}

	return c.JSON(resp)
}
