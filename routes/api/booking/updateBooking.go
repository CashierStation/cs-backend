package booking

import (
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UpdateBookingRequest struct {
	UnitID       uint       `validate:"omitempty,number" query:"unit_id"`
	CustomerName string     `validate:"omitempty" query:"customer_name"`
	Status       string     `validate:"omitempty,oneof=waiting accepted rejected" query:"status"`
	Time         *time.Time `validate:"omitempty" query:"time"`
}

var updateBookingValidator = lib.CreateValidator[UpdateBookingRequest]

type UpdateBookingResponse struct {
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
// @Param id path int true "Booking ID"
// @Param customer_name query string false "Customer name"
// @Param unit_id query int false "Unit ID"
// @Param time query string false "Booking time in RFC3339 format (ex: 2023-06-01T08:00:00Z)"
// @Param status query string false "Booking status" Enums(waiting,accepted,rejected)
// @Success 200 {object} booking.UpdateBookingResponse
// @Router /api/booking/{id} [put]
func UpdateBooking(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	var rawReqQuery UpdateBookingRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request query")
	}

	// validate query
	validationErrors := updateBookingValidator(rawReqQuery)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	bookingID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing unit id")
	}

	tx := global.DB.Begin()

	// validate booking exists
	booking, err := models.GetBookingWithRentalID(tx, user.RentalID, uint(bookingID))
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("Booking not found")
	}

	// validate unit belongs to rental
	if rawReqQuery.UnitID != 0 {
		_, err = models.GetUnit(tx, rawReqQuery.UnitID, user.RentalID)
		if err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusBadRequest).SendString("Unit not found")
		}
	}

	if rawReqQuery.Time != nil && rawReqQuery.Time.Before(time.Now()) {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("Booking time must be in the future")
	}

	booking, err = models.UpdateBooking(tx, booking.ID, rawReqQuery.UnitID, rawReqQuery.CustomerName, rawReqQuery.Status, rawReqQuery.Time)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating booking")
	}

	tx.Commit()

	// return booking
	return c.JSON(UpdateBookingResponse{
		ID:           booking.ID,
		CustomerName: booking.CustomerName,
		UnitID:       booking.UnitID,
		Time:         booking.Time,
		Status:       booking.Status,
	})
}
