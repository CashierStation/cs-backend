package unitsession

import (
	"csbackend/enum"
	"csbackend/global"
	"csbackend/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type startUnitSessionsStatus struct {
	Status     enum.UnitStatus `json:"status"`
	StartTime  *time.Time      `json:"start_time"`
	FinishTime *time.Time      `json:"finish_time"`
}

type StartUnitSessionsResponse struct {
	UnitID   int                     `json:"unit_id"`
	UnitName string                  `json:"unit_name"`
	Status   startUnitSessionsStatus `json:"status"`
}

// @Security SessionToken
// Unit godoc
// @Summary
// @Schemes
// @Description Sesi pemakaian unit
// @Tags api/unit_session
// @Accept x-www-form-urlencoded
// @Produce json
// @Param unit_id path int true "Unit ID"
// @Success 200 {object} unitsession.StartUnitSessionsResponse
// @Router /api/unit_session/start/{unit_id} [put]
func StartUnitSessions(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	// get unit id from path
	unitID, err := strconv.Atoi(c.Params("unit_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing unit id")
	}

	tx := global.DB.Begin()

	_, err = models.GetUnit(tx, uint(unitID), user.RentalID)
	if err == gorm.ErrRecordNotFound {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("Unit not found")
	}

	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("Error validating unit ownership")
	}

	// check if unit session is already started
	unitSession, err := models.GetLastUnitSession(tx, uint(unitID))
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).SendString("Error validating unit session history")
	}

	if err != gorm.ErrRecordNotFound && unitSession.FinishTime.Time.IsZero() {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("Unit is still in use")
	}

	// create unit session
	unitSession, err = models.CreateUnitSession(tx, uint(unitID))
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating unit session")
	}

	tx.Commit()

	result := StartUnitSessionsResponse{
		UnitID:   int(unitSession.UnitID),
		UnitName: unitSession.Unit.Name,
		Status: startUnitSessionsStatus{
			Status:     enum.InUse,
			StartTime:  &unitSession.StartTime.Time,
			FinishTime: nil,
		},
	}

	return c.JSON(result)
}
