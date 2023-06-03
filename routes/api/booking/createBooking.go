package booking

import (
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CreateBookingRequest struct {
	CustomerName string    `validate:"required" query:"customer_name"`
	UnitID       uint      `validate:"required,number" query:"unit_id"`
	Time         time.Time `validate:"required" query:"time"`
}

var createBookingValidator = lib.CreateValidator[CreateBookingRequest]

type CreateBookingResponse struct {
	ID           uint      `json:"id"`
	CustomerName string    `json:"customer_name"`
	UnitID       uint      `json:"unit_id"`
	Time         time.Time `json:"time"`
	Status       string    `json:"status"`
}

// @Security Booking
// Booking godoc
// @Summary
// @Schemes
// @Description Submit new booking
// @Tags api/booking
// @Accept x-www-form-urlencoded
// @Produce json
// @Param customer_name query string true "Customer name"
// @Param unit_id query int true "Unit ID"
// @Param time query string true "Booking time in RFC3339 format (ex: 2023-06-01T08:00:00Z)"
// @Success 200 {object} booking.CreateBookingResponse
// @Router /api/booking [post]
func CreateBooking(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	var rawReqQuery CreateBookingRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request query")
	}

	// validate query
	validationErrors := createBookingValidator(rawReqQuery)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	tx := global.DB.Begin()

	// validate unit belongs to rental
	_, err = models.GetUnit(tx, rawReqQuery.UnitID, user.RentalID)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("Unit not found")
	}

	if rawReqQuery.Time.Before(time.Now()) {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("Booking time must be in the future")
	}

	booking, err := models.CreateBooking(tx, rawReqQuery.CustomerName, rawReqQuery.UnitID, rawReqQuery.Time)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating booking")
	}

	tx.Commit()

	// return booking
	return c.JSON(CreateBookingResponse{
		ID:           booking.ID,
		CustomerName: booking.CustomerName,
		UnitID:       booking.UnitID,
		Time:         booking.Time,
		Status:       booking.Status,
	})
}
