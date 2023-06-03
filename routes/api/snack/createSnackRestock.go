package snack

import (
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

type CreateSnackRestockRequest struct {
	SnackID  uint `validate:"required" query:"snack_id"`
	Quantity uint `validate:"required" query:"quantity"`
	Price    uint `validate:"required" query:"price"`
}

var createSnackRestockValidator = lib.CreateValidator[CreateSnackRestockRequest]

type CreateSnackRestockResponse struct {
	ID         uint `json:"id"`
	SnackID    uint `json:"snack_id"`
	Quantity   uint `json:"quantity"`
	TotalPrice uint `json:"total_price"`
}

// @Security SessionToken
// Snack godoc
// @Summary
// @Schemes
// @Description Snack
// @Tags api/snack
// @Accept x-www-form-urlencoded
// @Produce json
// @Param snack_id query int true "Snack ID"
// @Param quantity query int true "Quantity"
// @Param price query int true "Total harga kulakan"
// @Success 200 {object} snack.PostSnackResponse
// @Router /api/snack/restock [post]
func CreateSnackRestock(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	var rawReqQuery CreateSnackRestockRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request query")
	}

	// validate query
	validationErrors := createSnackRestockValidator(rawReqQuery)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	tx := global.DB.Begin()

	// get snack
	snack, err := models.GetSnack(tx, user.RentalID, rawReqQuery.SnackID)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).SendString("Error getting snack, please check if snack_id is correct")
	}

	// create snack restock
	snackRestock, err := models.CreateSnackRestock(tx, user.RentalID, snack.ID, int(rawReqQuery.Quantity), int(rawReqQuery.Price))
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	tx.Commit()

	return c.JSON(&CreateSnackRestockResponse{
		ID:         snackRestock.ID,
		SnackID:    snackRestock.SnackID,
		Quantity:   uint(snackRestock.Quantity),
		TotalPrice: uint(snackRestock.Total),
	})
}
