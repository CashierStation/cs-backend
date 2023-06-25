package login

import (
	"csbackend/authenticator"
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginPostRequest struct {
	Username    string `validate:"required,min=3,max=32"`
	Password    string `validate:"required,number,min=6,max=32"`
	AccessToken string `validate:"required" query:"access_token"`
}

var loginPostValidator = lib.CreateValidator[LoginPostRequest]

type LoginPostResponse struct {
	Username     string `json:"username"`
	SessionToken string `json:"session_token"`
}

// Login godoc
// @Summary Login a new employee/owner account
// @Schemes
// @Description dev: http://localhost:8080/auth/login
// @Description prod: http://csbackend.sivr.tech/auth/login
// @Tags auth
// @Accept x-www-form-urlencoded
// @Param access_token query string true "Access token from Auth0"
// @Param username query string true "Username" minlength(3) maxlength(32)
// @Param password query string true "Password (Numeric)" minlength(6) maxlength(32)
// @Success 200 {object} LoginPostResponse
// @Produce json
// @Router /auth/login [post]
func POST(c *fiber.Ctx) error {
	var rawReqQuery LoginPostRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error parsing request query", err)
	}

	// validate query
	validationErrors := loginPostValidator(rawReqQuery)
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
	_, err = models.GetRentalById(tx, rentalId)
	if err != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error getting rental", err)
	}

	// Check if user already exists
	employee, err := models.GetEmployeeInRental(tx, rentalId, rawReqQuery.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "User does not exist", err)
	}

	if err != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error getting employee", err)
	}

	// Match password hash
	match := bcrypt.CompareHashAndPassword([]byte(employee.PasswordHash), []byte(rawReqQuery.Password))
	if match != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Wrong password", err)
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
	res := LoginPostResponse{
		Username:     employee.Username,
		SessionToken: sessionToken,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
