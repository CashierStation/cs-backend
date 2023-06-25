package snack

import (
	"csbackend/global"
	"csbackend/lib"
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
// @Tags api/snack
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
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error parsing snack id", err)
	}

	tx := global.DB.Begin()
	snack, err := models.GetSnack(tx, user.RentalID, uint(snackID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lib.HTTPError(c, fiber.StatusNotFound, "Snack not found", err)
		}
		return lib.HTTPError(c, fiber.StatusInternalServerError, "Error getting snack", err)
	}

	tx.Delete(&snack)
	tx.Commit()

	return c.JSON(DeleteSnackResponse{
		Status: "success",
	})
}
