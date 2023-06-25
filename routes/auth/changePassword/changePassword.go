package changepassword

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

type ChangePasswordRequest struct {
	Username    string `validate:"required,min=3,max=32"`
	OldPassword string `validate:"required,number,min=6,max=32" query:"old_password"`
	NewPassword string `validate:"required,number,min=6,max=32" query:"new_password"`
	AccessToken string `validate:"required" query:"access_token"`
}

var ChangePasswordValidator = lib.CreateValidator[ChangePasswordRequest]

type ChangePasswordResponse struct {
	Username string `json:"username"`
	Status   string `json:"status"`
}

// ChangePassword godoc
// @Summary Change password for employee/owner account
// @Schemes
// @Description dev: http://localhost:8080/auth/change_password
// @Description prod: http://csbackend.sivr.tech/auth/change_password
// @Tags auth
// @Accept x-www-form-urlencoded
// @Param access_token query string true "Access token from Auth0"
// @Param username query string true "Username" minlength(3) maxlength(32)
// @Param old_password query string true "Password (Numeric)" minlength(6) maxlength(32)
// @Param new_password query string true "Password (Numeric)" minlength(6) maxlength(32)
// @Success 200 {object} ChangePasswordResponse
// @Produce json
// @Router /auth/change_password [post]
func POST(c *fiber.Ctx) error {
	var rawReqQuery ChangePasswordRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error parsing request query", err)
	}

	// validate query
	validationErrors := ChangePasswordValidator(rawReqQuery)
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
	match := bcrypt.CompareHashAndPassword([]byte(employee.PasswordHash), []byte(rawReqQuery.OldPassword))
	if match != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Wrong password", err)
	}

	// Update password
	hash, err := bcrypt.GenerateFromPassword([]byte(rawReqQuery.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error hashing password", err)
	}

	employee.PasswordHash = string(hash)
	err = tx.Save(&employee).Error
	if err != nil {
		tx.Rollback()
		return lib.HTTPError(c, fiber.StatusBadRequest, "Error saving employee", err)
	}

	tx.Commit()

	// Return response
	res := ChangePasswordResponse{
		Username: employee.Username,
		Status:   "success",
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
