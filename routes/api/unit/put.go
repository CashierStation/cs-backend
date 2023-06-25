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

type PutUnitRequest struct {
	Name        string `query:"name"`
	Category    string `query:"category"`
	HourlyPrice int    `query:"hourly_price"`
}

var putUnitValidator = lib.CreateValidator[PutUnitRequest]

type PutUnit struct {
	ID          uint   `json:"id"`
	RentalID    string `json:"rental_id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	HourlyPrice int    `json:"hourly_price"`
}

type PutUnitResponse struct {
	Unit PutUnit `json:"unit"`
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
// @Param name query string false "Unit name"
// @Param hourly_price query int false "Unit hourly price"
// @Success 200 {object} unit.PutUnitResponse
// @Router /api/unit/{id} [put]
func PUT(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	var rawReqQuery PutUnitRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error parsing request query", err)
	}

	// validate query
	validationErrors := putUnitValidator(rawReqQuery)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

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

	if rawReqQuery.Name != "" {
		unit.Name = rawReqQuery.Name
	}

	if rawReqQuery.HourlyPrice != 0 {
		unit.HourlyPrice = rawReqQuery.HourlyPrice
	}

	if rawReqQuery.Category != "" {
		unit.Category = rawReqQuery.Category
	}

	tx.Save(&unit)
	tx.Commit()

	return c.JSON(&PutUnitResponse{Unit: PutUnit{
		ID:          unit.ID,
		RentalID:    unit.RentalID,
		Name:        unit.Name,
		Category:    unit.Category,
		HourlyPrice: unit.HourlyPrice,
	}})
}
