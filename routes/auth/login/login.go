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
// @Description prod: https://csbackend.fly.dev/auth/login
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
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request query")
	}

	// validate query
	validationErrors := loginPostValidator(rawReqQuery)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	userinfoString, err := global.Authenticator.GetUserinfo(rawReqQuery.AccessToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error getting userinfo from access token")
	}

	var userInfo authenticator.UserInfo
	err = json.Unmarshal([]byte(userinfoString), &userInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error unmarshalling userinfo")
	}

	tx := global.DB.Begin()

	// Check which rental the sender is referring to
	rentalId := userInfo.Sub
	_, err = models.GetRentalById(tx, rentalId)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("Error getting rental")
	}

	// Check if user already exists
	employee, err := models.GetEmployeeInRental(tx, rentalId, rawReqQuery.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("User does not exist")
	}

	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("Error getting employee")
	}

	// Match password hash
	match := bcrypt.CompareHashAndPassword([]byte(employee.PasswordHash), []byte(rawReqQuery.Password))
	if match != nil {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("Wrong password")
	}

	// Create session token as if the user logged in
	sessionToken, err := global.Authenticator.GenerateRandomHex()
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("Error generating session token")
	}

	// Upsert session
	_, err = models.UpsertSession(tx, sessionToken, employee.ID)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).SendString("Error upserting session")
	}

	tx.Commit()

	// Return response
	res := LoginPostResponse{
		Username:     employee.Username,
		SessionToken: sessionToken,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
