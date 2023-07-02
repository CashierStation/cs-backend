package analytic

import (
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type GetUnitPlaytimeRequest struct {
	StartTime time.Time `validate:"required" query:"start_time"`
	EndTime   time.Time `validate:"required" query:"end_time"`
	GroupBy   string    `validate:"omitempty,oneof=unit_id unit_category" query:"group_by"`
}

var getUnitPlaytimeValidator = lib.CreateValidator[GetUnitPlaytimeRequest]

type GetUnitPlaytimeResponse struct {
	*models.HistoricalRevenue
}

// @Security SessionToken
// Booking godoc
// @Summary
// @Schemes
// @Description Get unit playtime analytic. Returns playtime in seconds
// @Tags api/analytic
// @Accept x-www-form-urlencoded
// @Produce json
// @Param group_by query string false "group by" default(unit_id) Enums(unit_id,unit_category)
// @Param start_time query string true "Start Time in RFC3339 format (ex: 2023-06-01T08:00:00Z)"
// @Param end_time query string true "End Time in RFC3339 format (ex: 2023-06-01T08:00:00Z)"
// @Success 200 {object} analytic.GetRevenueResponse
// @Router /api/analytic/unit/playtime [get]
func GetUnitPlaytime(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	var req GetUnitPlaytimeRequest

	// convert query to struct
	err := c.QueryParser(&req)
	if err != nil {
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error parsing request query", err)
	}

	// validate request
	validationErrors := getUnitPlaytimeValidator(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	groupBy := req.GroupBy
	if groupBy == "" {
		groupBy = "unit_id"
	}

	tx := global.DB.Begin()

	// get revenue
	revenue, err := models.GetUnitPlaytime(tx, user.RentalID, groupBy, req.StartTime, req.EndTime)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		log.Println(err)
		return lib.HTTPError(c, fiber.StatusInternalServerError, "Error getting revenue", err)
	}

	tx.Commit()

	return c.JSON(revenue)
}
