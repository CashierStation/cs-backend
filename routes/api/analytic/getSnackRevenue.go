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

type GetSnackRevenueRequest struct {
	Aggregation string    `validate:"omitempty,oneof=day week month" query:"aggregation"`
	StartTime   time.Time `validate:"required" query:"start_time"`
	EndTime     time.Time `validate:"required" query:"end_time"`
}

var getSnackRevenueValidator = lib.CreateValidator[GetSnackRevenueRequest]

type GetSnackRevenueResponse struct {
	*models.HistoricalRevenue
}

// @Security Analytic
// Booking godoc
// @Summary
// @Schemes
// @Description Get snack revenue analytic
// @Tags api/analytic
// @Accept x-www-form-urlencoded
// @Produce json
// @Param aggregation query string false "Aggregation" Enums(day,week,month)
// @Param start_time query string true "Start Time in RFC3339 format (ex: 2023-06-01T08:00:00Z)"
// @Param end_time query string true "End Time in RFC3339 format (ex: 2023-06-01T08:00:00Z)"
// @Success 200 {object} analytic.GetRevenueResponse
// @Router /api/analytic/snack/revenue [get]
func GetSnackRevenue(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	var req GetSnackRevenueRequest

	// convert query to struct
	err := c.QueryParser(&req)
	if err != nil {
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error parsing request query", err)
	}

	// validate request
	validationErrors := getSnackRevenueValidator(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	aggregation := req.Aggregation
	if aggregation == "" {
		aggregation = "day"
	}

	tx := global.DB.Begin()

	// get revenue
	revenue, err := models.GetSnackRevenue(tx, user.RentalID, aggregation, req.StartTime, req.EndTime)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		log.Println(err)
		return lib.HTTPError(c, fiber.StatusInternalServerError, "Error getting revenue", err)
	}

	tx.Commit()

	return c.JSON(revenue)
}
