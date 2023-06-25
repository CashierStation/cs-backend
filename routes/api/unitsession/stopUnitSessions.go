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

type stopUnitSessionsStatus struct {
	Status     enum.UnitStatus `json:"status"`
	StartTime  *time.Time      `json:"start_time"`
	FinishTime *time.Time      `json:"finish_time"`
}

type StopUnitSessionsResponse struct {
	UnitID   int                    `json:"unit_id"`
	UnitName string                 `json:"unit_name"`
	Status   stopUnitSessionsStatus `json:"status"`
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
// @Success 200 {object} unitsession.StopUnitSessionsResponse
// @Router /api/unit_session/stop/{unit_id} [put]
func StopUnitSessions(c *fiber.Ctx) error {
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

	if err == gorm.ErrRecordNotFound {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Unit not found", err)
	}

	if err != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusInternalServerError, "Error validating unit session history", err)
	}

	if !unitSession.FinishTime.Time.IsZero() {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Unit session already finished", err)
	}

	// stop unit session
	unitSession, err = models.StopUnitSession(tx, unitSession.ID)
	if err != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusInternalServerError, "Error stopping unit session", err)
	}

	tx.Commit()

	result := StartUnitSessionsResponse{
		UnitID:   int(unitSession.UnitID),
		UnitName: unitSession.Unit.Name,
		Status: startUnitSessionsStatus{
			Status:     enum.Idle, // check booking status first once its already implemented
			StartTime:  &unitSession.StartTime.Time,
			FinishTime: &unitSession.FinishTime.Time,
		},
	}

	return c.JSON(result)
}
