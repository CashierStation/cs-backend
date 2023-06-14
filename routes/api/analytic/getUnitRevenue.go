package analytic

import (
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type GetRevenueRequest struct {
	Aggregation string    `validate:"omitempty,oneof=day week month" query:"aggregation"`
	StartTime   time.Time `validate:"required" query:"start_time"`
	EndTime     time.Time `validate:"required" query:"end_time"`
}

var getRevenueValidator = lib.CreateValidator[GetRevenueRequest]

type GetRevenueResponse struct {
	*models.UnitHistoricalRevenue
}

// @Security Analytic
// Booking godoc
// @Summary
// @Schemes
// @Description Get unit revenue analytic
// @Tags api/analytic
// @Accept x-www-form-urlencoded
// @Produce json
// @Param aggregation query string false "Aggregation" Enums(day,week,month)
// @Param start_time query string true "Start Time in RFC3339 format (ex: 2023-06-01T08:00:00Z)"
// @Param end_time query string true "End Time in RFC3339 format (ex: 2023-06-01T08:00:00Z)"
// @Success 200 {object} analytic.GetRevenueResponse
// @Router /api/analytic/unit/revenue [get]
func GetRevenue(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	var req GetRevenueRequest

	// convert query to struct
	err := c.QueryParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request query")
	}

	// validate request
	validationErrors := getRevenueValidator(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	aggregation := req.Aggregation
	if aggregation == "" {
		aggregation = "day"
	}

	tx := global.DB.Begin()

	// get revenue
	revenue, err := models.GetRevenue(tx, user.RentalID, aggregation, req.StartTime, req.EndTime)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).SendString("Error getting revenue")
	}

	tx.Commit()

	return c.JSON(revenue)
}
