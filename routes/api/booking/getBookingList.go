package booking

import (
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type GetBookingListRequest struct {
	CustomerName string `validate:"omitempty" query:"customer_name"`
	UnitID       uint   `validate:"omitempty,number" query:"unit_id"`
	Status       string `validate:"omitempty,oneof=waiting accepted rejected" query:"status"`
	UnitInUse    *bool  `validate:"omitempty" query:"unit_in_use"`
	Offset       uint   `validate:"omitempty,number,min=0" query:"offset"`
	Limit        uint   `validate:"omitempty,number" query:"limit"`
}

var getBookingListValidator = lib.CreateValidator[GetBookingListRequest]

type GetBookingListResponse_01 struct {
	ID           uint      `json:"id"`
	CustomerName string    `json:"customer_name"`
	UnitID       uint      `json:"unit_id"`
	Time         time.Time `json:"time"`
	Status       string    `json:"status"`
}

type GetBookingListResponse struct {
	Bookings []GetBookingListResponse_01 `json:"bookings"`
}

// @Security Booking
// Booking godoc
// @Summary
// @Schemes
// @Description Get booking list
// @Tags api/booking
// @Accept x-www-form-urlencoded
// @Produce json
// @Param customer_name query string false "Search by customer name"
// @Param unit_id query int false "Select by Unit ID"
// @Param status query string false "Select by status" Enums(waiting,accepted,rejected)
// @Param unit_in_use query bool false "Select by unit in use"
// @Param limit query int false "Limit number of results" default(10)
// @Param offset query int false "Offset results" default(0)
// @Success 200 {object} booking.GetBookingListResponse
// @Router /api/booking [get]
func GetBookingList(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	var rawReqQuery GetBookingListRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error parsing request query", err)
	}

	// validate query
	validationErrors := getBookingListValidator(rawReqQuery)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	tx := global.DB.Begin()

	// validate unit belongs to rental
	if rawReqQuery.UnitID != 0 {
		_, err = models.GetUnit(tx, rawReqQuery.UnitID, user.RentalID)
		if err != nil {
			tx.Rollback()
			return lib.HTTPError(c, fiber.StatusBadRequest, "Unit not found", err)
		}
	}

	// get booking
	booking, err := models.GetBookingList(tx, user.RentalID, rawReqQuery.CustomerName, rawReqQuery.UnitID, rawReqQuery.Status, rawReqQuery.UnitInUse, int(rawReqQuery.Offset), int(rawReqQuery.Limit))
	if err != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusInternalServerError, "Error creating booking", err)
	}

	// commit transaction
	tx.Commit()

	// return booking
	bookings := []GetBookingListResponse_01{}
	for _, booking := range booking {
		bookings = append(bookings, GetBookingListResponse_01{
			ID:           booking.ID,
			CustomerName: booking.CustomerName,
			UnitID:       booking.UnitID,
			Time:         booking.Time,
			Status:       booking.Status,
		})
	}

	return c.JSON(GetBookingListResponse{Bookings: bookings})
}
