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
	Latest bool   `validate:"omitempty" query:"latest"`
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
	GrandTotal        int                        `json:"grand_total"`
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
// @Tags api/unit_session
// @Accept x-www-form-urlencoded
// @Produce json
// @Param unit_id query int false "unit id"
// @Param offset query int false "offset" default(0)
// @Param limit query int false "limit" default(10)
// @Param order query string false "order" default(desc) Enums(asc,desc)
// @Param sort_by query string false "sort_by" default(start_time) Enums(id,unit_id,start_time,finish_time,tarif)
// @Param latest query bool false "select only latest session for each unit" default(false)
// @Success 200 {object} unitsession.GetUnitSessionsResponse
// @Router /api/unit_session [get]
func GetUnitSessions(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employee)

	var rawReqQuery GetUnitSessionsRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error parsing request query", err)
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
	latest := rawReqQuery.Latest

	// set default values
	if order == "" {
		order = "desc"
	}

	if sortBy == "" {
		sortBy = "start_time"
	}

	if limit == 0 {
		limit = 10
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
			return lib.HTTPError(c, fiber.StatusBadRequest, "Unit not found", err)
		}

		if err != nil {
			tx.Rollback()
			return lib.HTTPError(c, fiber.StatusBadRequest, "Error getting unit", err)
		}
	}

	unitSessions, err := models.GetUnitSessions(tx, unitID, offset, limit, order, sortBy, latest)
	if err != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusInternalServerError, "Error getting unit sessions", err)
	}

	tx.Commit()

	for _, unitSession := range unitSessions {
		var snackTransactions []SnackTransactionResponse = []SnackTransactionResponse{}
		var snackTotal int = 0

		for _, snackTransaction := range unitSession.SnackTransactions {
			snackTransactions = append(snackTransactions, SnackTransactionResponse{
				ID:        snackTransaction.ID,
				SnackName: snackTransaction.Snack.Name,
				Quantity:  snackTransaction.Quantity,
				Total:     snackTransaction.Total,
			})

			snackTotal += snackTransaction.Total
		}

		var startTime = unitSession.StartTime.Time
		var finishTime = unitSession.FinishTime.Time

		resp := UnitSessionResponse{
			ID:                unitSession.ID,
			UnitID:            unitSession.UnitID,
			StartTime:         &startTime,
			FinishTime:        &finishTime,
			Tarif:             unitSession.Tarif,
			GrandTotal:        unitSession.Tarif + snackTotal,
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
