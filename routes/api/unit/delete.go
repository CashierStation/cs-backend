package unit

import (
	"csbackend/global"
	"csbackend/lib"
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
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error parsing unit id", err)
	}

	tx := global.DB.Begin()
	unit, err := models.GetUnit(tx, uint(unitID), user.RentalID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lib.HTTPError(c, fiber.StatusNotFound, "Unit not found", err)
		}
		return lib.HTTPError(c, fiber.StatusInternalServerError, "Error getting unit", err)
	}

	tx.Delete(&unit)
	tx.Commit()

	return c.JSON(DeleteUnitResponse{
		Status: "success",
	})
}
