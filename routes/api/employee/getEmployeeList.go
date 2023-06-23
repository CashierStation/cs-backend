package employee

import (
	"csbackend/authenticator"
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type GetEmployeeListRequest struct {
	AccessToken string `validate:"required" query:"access_token"`
}

var getEmployeeListValidator = lib.CreateValidator[GetEmployeeListRequest]

type Employee struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type GetEmployeeListResponse struct {
	Employees []Employee `json:"employees"`
}

// @Security SessionToken
// Employee godoc
// @Summary
// @Schemes
// @Description Get list of employees from access token
// @Tags api/employee
// @Accept x-www-form-urlencoded
// @Produce json
// @Param access_token query string true "Access token from Auth0"
// @Success 200 {object} user.GET.response
// @Router /api/employee/list [get]
func GetEmployeeList(c *fiber.Ctx) error {
	var rawReqQuery GetEmployeeListRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request query")
	}

	// validate query
	validationErrors := getEmployeeListValidator(rawReqQuery)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	userinfoString, err := global.Authenticator.GetUserinfo(rawReqQuery.AccessToken)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Error getting userinfo from access token")
	}

	var userInfo authenticator.UserInfo
	err = json.Unmarshal([]byte(userinfoString), &userInfo)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Error unmarshalling userinfo")
	}

	tx := global.DB.Begin()

	employees, err := models.GetAllEmployeeInRental(tx, userInfo.Sub)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error getting employees")
	}

	tx.Commit()

	resp := GetEmployeeListResponse{}

	for _, employee := range employees {
		resp.Employees = append(resp.Employees, Employee{
			ID:        employee.ID,
			Username:  employee.Username,
			Role:      employee.Role.Name,
			CreatedAt: employee.CreatedAt.String(),
		})
	}

	return c.JSON(resp)
}
