package unit

import (
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"

	"github.com/gofiber/fiber/v2"
)

type PostUnitRequest struct {
	Name        string `validate:"required" query:"name"`
	Category    string `validate:"required" query:"category"`
	HourlyPrice int    `validate:"required" query:"hourly_price"`
}

var postUnitValidator = lib.CreateValidator[PostUnitRequest]

type PostUnit struct {
	ID          uint   `json:"id"`
	RentalID    string `json:"rental_id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	HourlyPrice int    `json:"hourly_price"`
}

type PostUnitResponse struct {
	Unit PostUnit `json:"unit"`
}

// @Security SessionToken
// Unit godoc
// @Summary
// @Schemes
// @Description Unit
// @Tags api/unit
// @Accept x-www-form-urlencoded
// @Produce json
// @Param name query string true "Unit name"
// @Param hourly_price query int true "Unit hourly price"
// @Param category query string true "Unit category"
// @Success 200 {object} unit.PostUnitResponse
// @Router /api/unit [post]
func POST(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	var rawReqQuery PostUnitRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request query")
	}

	// validate query
	validationErrors := postUnitValidator(rawReqQuery)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	tx := global.DB.Begin()
	newUnit, err := models.CreateUnit(tx, rawReqQuery.Name, rawReqQuery.HourlyPrice, rawReqQuery.Category, user.RentalID)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating unit")
	}

	tx.Commit()

	return c.JSON(&PostUnitResponse{Unit: PostUnit{
		ID:          newUnit.ID,
		RentalID:    newUnit.RentalID,
		Name:        newUnit.Name,
		Category:    newUnit.Category,
		HourlyPrice: newUnit.HourlyPrice,
	}})
}
