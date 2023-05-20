package unit

import (
	"csbackend/global"
	"csbackend/models"

	"github.com/gofiber/fiber/v2"
)

type GetUnitResponse struct {
	Units []UnitResponse `json:"units"`
}

// @Security SessionToken
// Unit godoc
// @Summary
// @Schemes
// @Description Unit
// @Tags unit
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} unit.GetUnitResponse
// @Router /api/unit [get]
func GET(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	tx := global.DB.Begin()
	units, err := models.GetAllUnits(tx, user.RentalID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error getting units")
	}

	var unitResponses []UnitResponse
	for _, unit := range units {
		unitResponses = append(unitResponses, UnitResponse{
			ID:          unit.ID,
			RentalID:    unit.RentalID,
			Name:        unit.Name,
			HourlyPrice: unit.HourlyPrice,
		})
	}

	return c.JSON(GetUnitResponse{
		Units: unitResponses,
	})
}
