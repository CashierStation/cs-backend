package snack

import (
	"csbackend/global"
	"csbackend/models"

	"github.com/gofiber/fiber/v2"
)

type GetSnackResponse struct {
	Snacks []SnackResponse `json:"snacks"`
}

// @Security SessionToken
// Snack godoc
// @Summary
// @Schemes
// @Description Snack
// @Tags snack
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} snack.GetSnackResponse
// @Router /api/snack [get]
func GET(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	tx := global.DB.Begin()
	snacks, err := models.GetAllSnacks(tx, user.RentalID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error getting snacks")
	}

	var snackResponses []SnackResponse = []SnackResponse{}
	for _, snack := range snacks {
		snackResponses = append(snackResponses, SnackResponse{
			ID:       snack.ID,
			RentalID: snack.RentalID,
			Name:     snack.Name,
			Category: snack.Category,
			Price:    snack.Price,
		})
	}

	return c.JSON(GetSnackResponse{
		Snacks: snackResponses,
	})
}
