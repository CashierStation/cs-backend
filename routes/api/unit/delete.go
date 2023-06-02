package unit

import (
	"csbackend/global"
	"csbackend/models"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DeleteUnitResponse struct {
	Status string `json:"status"`
}

// @Security SessionToken
// Unit godoc
// @Summary
// @Schemes
// @Description Unit
// @Tags api/unit
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path int true "Unit ID"
// @Success 200 {object} unit.DeleteUnitResponse
// @Router /api/unit/{id} [delete]
func DELETE(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	// get unit id from path
	unitID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing unit id")
	}

	tx := global.DB.Begin()
	unit, err := models.GetUnit(tx, uint(unitID), user.RentalID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).SendString("Unit not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Error getting unit")
	}

	tx.Delete(&unit)
	tx.Commit()

	return c.JSON(DeleteUnitResponse{
		Status: "success",
	})
}
