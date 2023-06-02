package unit

import (
	"csbackend/enum"
	"csbackend/global"
	"csbackend/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type GetUnitStatus struct {
	Status     enum.UnitStatus `json:"status"`
	StartTime  *time.Time      `json:"latest_start_time"`
	FinishTime *time.Time      `json:"latest_finish_time"`
	Tarif      int             `json:"tarif"`
}

type GetUnit struct {
	ID          uint          `json:"id"`
	RentalID    string        `json:"rental_id"`
	Name        string        `json:"name"`
	Category    string        `json:"category"`
	HourlyPrice int           `json:"hourly_price"`
	Status      GetUnitStatus `json:"status"`
}

type GetUnitResponse struct {
	Units []GetUnit `json:"units"`
}

// @Security SessionToken
// Unit godoc
// @Summary
// @Schemes
// @Description Unit
// @Tags api/unit
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

	// Get all units
	var unitResponses []GetUnit = []GetUnit{}
	var unitIds = []uint{}
	for _, unit := range units {
		unitResponses = append(unitResponses, GetUnit{
			ID:          unit.ID,
			RentalID:    unit.RentalID,
			Name:        unit.Name,
			Category:    unit.Category,
			HourlyPrice: unit.HourlyPrice,
			Status:      GetUnitStatus{},
		})

		unitIds = append(unitIds, unit.ID)
	}

	// Get all unit statuses
	unitStatuses, err := models.GetLastUnitStatuses(tx, unitIds)
	if err != nil {
		println(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Error getting unit sessions history")
	}

	for _, unitStatus := range unitStatuses {
		for i, unitResponse := range unitResponses {
			if unitResponse.ID != unitStatus.UnitID {
				continue
			}

			unitResponses[i].Status = GetUnitStatus{
				Status:     unitStatus.Status,
				StartTime:  unitStatus.StartTime,
				FinishTime: unitStatus.FinishTime,
			}
		}
	}

	return c.JSON(GetUnitResponse{
		Units: unitResponses,
	})
}
