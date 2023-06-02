package snack

import (
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PutSnackRequest struct {
	Name     string `query:"name"`
	Category string `query:"category"`
	Price    int    `query:"price"`
}

var putSnackValidator = lib.CreateValidator[PutSnackRequest]

type PutSnackResponse struct {
	Snack SnackResponse `json:"snack"`
}

// @Security SessionToken
// Snack godoc
// @Summary
// @Schemes
// @Description Snack
// @Tags api/snack
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path int true "Snack ID"
// @Param name query string false "Snack name"
// @Param category query string false "Snack category"
// @Param price query int false "Snack price"
// @Success 200 {object} snack.PutSnackResponse
// @Router /api/snack/{id} [put]
func PUT(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	var rawReqQuery PutSnackRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request query")
	}

	// validate query
	validationErrors := putSnackValidator(rawReqQuery)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	// get snack id from path
	snackID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing snack id")
	}

	tx := global.DB.Begin()
	snack, err := models.GetSnack(tx, user.RentalID, uint(snackID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).SendString("Snack not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Error getting snack")
	}

	if rawReqQuery.Name != "" {
		snack.Name = rawReqQuery.Name
	}

	if rawReqQuery.Price != 0 {
		snack.Price = rawReqQuery.Price
	}

	if rawReqQuery.Category != "" {
		snack.Category = rawReqQuery.Category
	}

	tx.Save(&snack)
	tx.Commit()

	return c.JSON(&PutSnackResponse{Snack: SnackResponse{
		ID:       snack.ID,
		RentalID: snack.RentalID,
		Name:     snack.Name,
		Category: snack.Category,
		Price:    snack.Price,
	}})
}
