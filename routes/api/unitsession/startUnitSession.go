package unitsession

import (
	"csbackend/enum"
	"csbackend/global"
	"csbackend/lib"
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
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error parsing unit id", err)
	}

	tx := global.DB.Begin()

	_, err = models.GetUnit(tx, uint(unitID), user.RentalID)
	if err == gorm.ErrRecordNotFound {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Unit not found", err)
	}

	if err != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error validating unit ownership", err)
	}

	// check if unit session is already started
	unitSession, err := models.GetLastUnitSession(tx, uint(unitID))
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusInternalServerError, "Error validating unit session history", err)
	}

	if err != gorm.ErrRecordNotFound && unitSession.FinishTime.Time.IsZero() {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Unit is still in use", err)
	}

	// create unit session
	unitSession, err = models.CreateUnitSession(tx, uint(unitID))
	if err != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusInternalServerError, "Error creating unit session", err)
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
