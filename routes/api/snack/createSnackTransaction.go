package snack

import (
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

type CreateSnackTransactionRequest struct {
	UnitID   uint `validate:"required" query:"unit_id"`
	SnackID  uint `validate:"required" query:"snack_id"`
	Quantity uint `validate:"required" query:"quantity"`
}

var createSnackTransactionValidator = lib.CreateValidator[CreateSnackTransactionRequest]

type CreateSnackTransactionResponse struct {
	ID            uint `json:"id"`
	UnitID        uint `json:"unit_id"`
	UnitSessionID uint `json:"unit_session_id"`
	SnackID       uint `json:"snack_id"`
	Quantity      uint `json:"quantity"`
	TotalPrice    uint `json:"total_price"`
}

// @Security SessionToken
// Snack godoc
// @Summary
// @Schemes
// @Description Snack
// @Tags api/snack
// @Accept x-www-form-urlencoded
// @Produce json
// @Param unit_id query int true "Unit ID"
// @Param snack_id query int true "Snack ID"
// @Param quantity query int true "Quantity"
// @Success 200 {object} snack.PostSnackResponse
// @Router /api/snack/transaction [post]
func CreateSnackTransaction(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	var rawReqQuery CreateSnackTransactionRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request query")
	}

	// validate query
	validationErrors := createSnackTransactionValidator(rawReqQuery)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	tx := global.DB.Begin()

	// get unit session
	unitSession, err := models.GetLastUnitSession(tx, rawReqQuery.UnitID)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).SendString("Error getting unit session, please check if unit_id is correct")
	}

	if unitSession.FinishTime.Valid {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("Unit session already finished")
	}

	// get snack
	snack, err := models.GetSnack(tx, user.RentalID, rawReqQuery.SnackID)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).SendString("Error getting snack, please check if snack_id is correct")
	}

	// create snack transaction
	snackTransaction, err := models.CreateSnackTransaction(tx, unitSession.ID, snack.ID, int(rawReqQuery.Quantity))
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	tx.Commit()

	return c.JSON(&CreateSnackTransactionResponse{
		ID:            snackTransaction.ID,
		UnitID:        unitSession.UnitID,
		UnitSessionID: snackTransaction.UnitSessionID,
		SnackID:       snackTransaction.SnackID,
		Quantity:      uint(snackTransaction.Quantity),
		TotalPrice:    uint(snackTransaction.Total),
	})
}
