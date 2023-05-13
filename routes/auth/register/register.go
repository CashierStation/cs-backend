package register

import (
	"csbackend/authenticator"
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"
	"encoding/json"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type RegisterPostRequest struct {
	Username    string `validate:"required,min=3,max=32"`
	Password    string `validate:"required,number,min=6,max=32"`
	Role        string `validate:"omitempty"`
	AccessToken string `validate:"required" query:"access_token"`
}

var registerPostValidator = lib.CreateValidator[RegisterPostRequest]

type RegisterPostResponse struct {
	Username     string `json:"username"`
	Role         string `json:"role"`
	SessionToken string `json:"session_token"`
}

// Register godoc
// @Summary Register a new employee/owner account
// @Schemes
// @Description dev: http://localhost:8080/auth/register
// @Description prod: https://csbackend.fly.dev/auth/register
// @Tags auth
// @Accept x-www-form-urlencoded
// @Param access_token query string true "Access token from Auth0"
// @Param username query string true "Username" minlength(3) maxlength(32)
// @Param password query string true "Password (Numeric)" minlength(6) maxlength(32)
// @Param role query string true "Role" Enums(owner,karyawan)
// @Success 200 {object} RegisterPostResponse
// @Produce json
// @Router /auth/register [post]
func POST(c *fiber.Ctx) error {
	var rawReqQuery RegisterPostRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// validate query
	validationErrors := registerPostValidator(rawReqQuery)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	userinfoString, err := global.Authenticator.GetUserinfo(rawReqQuery.AccessToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var userInfo authenticator.UserInfo
	err = json.Unmarshal([]byte(userinfoString), &userInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	tx := global.DB.Begin()

	// Check which rental the sender is referring to
	rentalId := userInfo.Sub
	rental, err := models.GetOrCreateRental(tx, rentalId, userInfo.Email)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// Check if user already exists
	_, err = models.GetEmployeeByUsername(tx, rawReqQuery.Username)
	if err == nil {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("User already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// Hash password
	hashedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(rawReqQuery.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("Error hashing password")
	}

	// Find role
	role, err := models.GetRoleByName(tx, rawReqQuery.Role)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("Role does not exist")
	}

	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).JSON("Error finding role")
	}

	// Create user
	employee, err := models.CreateEmployee(tx, rawReqQuery.Username, string(hashedPasswordByte), role.ID, rental.ID)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// Create session token as if the user logged in
	sessionToken, err := global.Authenticator.GenerateRandomHex()
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	tx.Commit()

	// Return response
	res := RegisterPostResponse{
		Username:     employee.Username,
		Role:         role.Name,
		SessionToken: sessionToken,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
