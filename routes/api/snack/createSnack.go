package snack

import (
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"

	"github.com/gofiber/fiber/v2"
)

type PostSnackRequest struct {
	Name     string `validate:"required" query:"name"`
	Category string `validate:"required" query:"category"`
	Stock    int    `query:"stock"`
	Price    int    `validate:"required" query:"price"`
}

var postSnackValidator = lib.CreateValidator[PostSnackRequest]

type SnackResponse struct {
	ID       uint   `json:"id"`
	RentalID string `json:"rental_id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Stock    int    `json:"stock"`
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
// @Tags api/snack
// @Accept x-www-form-urlencoded
// @Produce json
// @Param name query string true "Snack name"
// @Param price query int true "Snack price"
// @Param category query string true "Snack category"
// @Param stock query int false "Snack stock" default(0)
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
	newSnack, err := models.CreateSnack(tx, user.RentalID, rawReqQuery.Name, rawReqQuery.Category, rawReqQuery.Price)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating snack")
	}

	tx.Commit()

	return c.JSON(&PostSnackResponse{Snack: SnackResponse{
		ID:       newSnack.ID,
		RentalID: newSnack.RentalID,
		Name:     newSnack.Name,
		Category: newSnack.Category,
		Stock:    newSnack.Stock,
		Price:    newSnack.Price,
	}})
}
