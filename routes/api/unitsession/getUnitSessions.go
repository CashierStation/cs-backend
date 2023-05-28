package unitsession

import (
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"
	"time"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type GetUnitSessionsRequest struct {
	UnitID uint   `validate:"omitempty,number" query:"unit_id"`
	Offset uint   `validate:"omitempty,number,min=0"`
	Limit  uint   `validate:"omitempty,number,min=1"`
	Order  string `validate:"omitempty,oneof=asc desc" query:"order"`
	SortBy string `validate:"omitempty,oneof=id unit_id start_time finish_time tarif" query:"sort_by"`
}

var getUnitSessionsValidator = lib.CreateValidator[GetUnitSessionsRequest]

type SnackTransactionResponse struct {
	ID        uint   `json:"id"`
	SnackName string `json:"snack_name"`
	Quantity  int    `json:"quantity"`
	Total     int    `json:"total"`
}

type UnitSessionResponse struct {
	ID                uint                       `json:"id"`
	UnitID            uint                       `json:"unit_id"`
	StartTime         *time.Time                 `json:"start_time"`
	FinishTime        *time.Time                 `json:"finish_time"`
	Tarif             int                        `json:"tarif"`
	SnackTransactions []SnackTransactionResponse `json:"snack_transactions"`
}

type GetUnitSessionsResponse struct {
	UnitSessions []UnitSessionResponse `json:"unit_sessions"`
	Total        uint                  `json:"total"`
}

// @Security SessionToken
// Unit godoc
// @Summary
// @Schemes
// @Description Sesi pemakaian unit
// @Tags unit_session
// @Accept x-www-form-urlencoded
// @Produce json
// @Param unit_id query int false "unit id"
// @Param offset query int false "offset" default(0)
// @Param limit query int false "limit" default(10)
// @Param order query string false "order" default(desc) Enums(asc,desc)
// @Param sort_by query string false "sort_by" default(start_time) Enums(id,unit_id,start_time,finish_time,tarif)
// @Success 200 {object} unitsession.GetUnitSessionsResponse
// @Router /api/unit_session [get]
func GetUnitSessions(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	var rawReqQuery GetUnitSessionsRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request query")
	}

	// validate query
	validationErrors := getUnitSessionsValidator(rawReqQuery)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	// get query values
	unitID := rawReqQuery.UnitID
	offset := rawReqQuery.Offset
	limit := rawReqQuery.Limit
	order := rawReqQuery.Order
	sortBy := rawReqQuery.SortBy

	// set default values
	if order == "" {
		order = "desc"
	}

	if sortBy == "" {
		sortBy = "start_time"
	}

	var result = GetUnitSessionsResponse{
		UnitSessions: []UnitSessionResponse{},
	}

	tx := global.DB.Begin()

	// validate unit existence
	if unitID != 0 {
		_, err := models.GetUnit(tx, unitID, user.RentalID)
		if err == gorm.ErrRecordNotFound {
			tx.Rollback()
			return c.Status(fiber.StatusBadRequest).SendString("Unit not found")
		}

		if err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusBadRequest).SendString("Error getting unit")
		}
	}

	unitSessions, err := models.GetUnitSessions(tx, unitID, offset, limit, order, sortBy)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).SendString("Error getting unit sessions")
	}

	tx.Commit()

	for _, unitSession := range unitSessions {
		var snackTransactions []SnackTransactionResponse = []SnackTransactionResponse{}
		for _, snackTransaction := range unitSession.SnackTransactions {
			snackTransactions = append(snackTransactions, SnackTransactionResponse{
				ID:        snackTransaction.ID,
				SnackName: snackTransaction.Snack.Name,
				Quantity:  snackTransaction.Quantity,
				Total:     snackTransaction.Total,
			})
		}

		var startTime = unitSession.StartTime.Time
		var finishTime = unitSession.FinishTime.Time

		resp := UnitSessionResponse{
			ID:                unitSession.ID,
			UnitID:            unitSession.UnitID,
			StartTime:         &startTime,
			FinishTime:        &finishTime,
			Tarif:             unitSession.Tarif,
			SnackTransactions: snackTransactions,
		}

		if resp.FinishTime.IsZero() {
			resp.FinishTime = nil
		}

		result.UnitSessions = append(result.UnitSessions, resp)
	}

	result.Total = uint(len(result.UnitSessions))

	return c.JSON(result)
}
