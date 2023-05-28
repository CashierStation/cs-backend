package snack

import (
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"

	"github.com/gofiber/fiber/v2"
)

type PostSnackRequest struct {
	Name  string `validate:"required" query:"name"`
	Price int    `validate:"required" query:"price"`
}

var postSnackValidator = lib.CreateValidator[PostSnackRequest]

type SnackResponse struct {
	ID       uint   `json:"id"`
	RentalID string `json:"rental_id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
}

type PostSnackResponse struct {
	Snack SnackResponse `json:"snack"`
}

// @Security SessionToken
// Snack godoc
// @Summary
// @Schemes
// @Description Snack
// @Tags snack
// @Accept x-www-form-urlencoded
// @Produce json
// @Param name query string true "Snack name"
// @Param price query int true "Snack price"
// @Success 200 {object} snack.PostSnackResponse
// @Router /api/snack [post]
func POST(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	var rawReqQuery PostSnackRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request query")
	}

	// validate query
	validationErrors := postSnackValidator(rawReqQuery)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	tx := global.DB.Begin()
	newSnack, err := models.CreateSnack(tx, user.RentalID, rawReqQuery.Name, rawReqQuery.Price)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating snack")
	}

	tx.Commit()

	return c.JSON(&PostSnackResponse{Snack: SnackResponse{
		ID:       newSnack.ID,
		RentalID: newSnack.RentalID,
		Name:     newSnack.Name,
		Price:    newSnack.Price,
	}})
}
