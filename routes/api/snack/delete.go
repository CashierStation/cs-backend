package snack

import (
	"csbackend/global"
	"csbackend/models"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DeleteSnackResponse struct {
	Status string `json:"status"`
}

// @Security SessionToken
// Snack godoc
// @Summary
// @Schemes
// @Description Snack
// @Tags snack
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path int true "Snack ID"
// @Success 200 {object} snack.DeleteSnackResponse
// @Router /api/snack/{id} [delete]
func DELETE(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	// get snack id from path
	snackID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing snack id")
	}

	tx := global.DB.Begin()
	snack, err := models.GetSnack(tx, user.RentalID, uint(snackID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).SendString("Snack not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Error getting snack")
	}

	tx.Delete(&snack)
	tx.Commit()

	return c.JSON(DeleteSnackResponse{
		Status: "success",
	})
}
