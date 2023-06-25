package register

import (
	"csbackend/authenticator"
	"csbackend/enum"
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"
	"encoding/json"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RegisterPostRequest struct {
	Username    string `validate:"required,min=3,max=32"`
	Password    string `validate:"required,number,min=6,max=32"`
	AccessToken string `validate:"required" query:"access_token"`
}

var registerPostValidator = lib.CreateValidator[RegisterPostRequest]

type RegisterPostResponse struct {
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	Role         string `json:"role"`
	SessionToken string `json:"session_token"`
}

// Register godoc
// @Summary Register a new employee/owner account
// @Schemes
// @Description dev: http://localhost:8080/auth/register
// @Description prod: http://csbackend.sivr.tech/auth/register
// @Tags auth
// @Accept x-www-form-urlencoded
// @Param access_token query string true "Access token from Auth0"
// @Param username query string true "Username" minlength(3) maxlength(32)
// @Param password query string true "Password (Numeric)" minlength(6) maxlength(32)
// @Success 200 {object} RegisterPostResponse
// @Produce json
// @Router /auth/register [post]
func POST(c *fiber.Ctx) error {
	var rawReqQuery RegisterPostRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error parsing request query", err)
	}

	// validate query
	validationErrors := registerPostValidator(rawReqQuery)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	userinfoString, err := global.Authenticator.GetUserinfo(rawReqQuery.AccessToken)
	if err != nil {
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error getting userinfo from access token", err)
	}

	var userInfo authenticator.UserInfo
	err = json.Unmarshal([]byte(userinfoString), &userInfo)
	if err != nil {
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error unmarshalling userinfo", err)
	}

	tx := global.DB.Begin()

	// Check which rental the sender is referring to
	rentalId := userInfo.Sub
	rental, err := models.GetOrCreateRental(tx, rentalId, userInfo.Email)
	if err != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error getting rental", err)
	}

	// Check if user already exists
	_, err = models.GetEmployeeInRental(tx, rentalId, rawReqQuery.Username)
	if err == nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "User already exists in that rental", err)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error getting user", err)
	}

	// Hash password
	hashedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(rawReqQuery.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error hashing password", err)
	}

	// Check if rental already has at least one employee, which can be the owner
	hasEmployee, err := models.RentalHasEmployee(tx, rentalId)
	if err != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error checking if rental has employee", err)
	}

	selectedRole := string(enum.Karyawan)
	if !hasEmployee {
		selectedRole = string(enum.Owner)
	}

	// Find role
	role, err := models.GetRoleByName(tx, selectedRole)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).JSON("Error finding role")
	}

	// Create user
	uuid := uuid.New()
	employee, err := models.CreateEmployee(tx, uuid.String(), rawReqQuery.Username, string(hashedPasswordByte), role.ID, rental.ID)
	if err != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error creating user", err)
	}

	// Create session token as if the user logged in
	sessionToken, err := global.Authenticator.GenerateRandomHex()
	if err != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error generating session token", err)
	}

	// Upsert session
	_, err = models.UpsertSession(tx, sessionToken, employee.ID)
	if err != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error upserting session", err)
	}

	tx.Commit()

	// Return response
	res := RegisterPostResponse{
		UserID:       employee.ID,
		Username:     employee.Username,
		Role:         selectedRole,
		SessionToken: sessionToken,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
