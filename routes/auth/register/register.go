package register

import (
	"csbackend/authenticator"
	"csbackend/global"
	"csbackend/lib"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type RegisterPostRequest struct {
	Username    string `validate:"required,min=3,max=32"`
	Password    string `validate:"required,min=8,max=32"`
	Role        string `validate:"omitempty"`
	AccessToken string `validate:"required" query:"access_token"`
}

var registerPostValidator = lib.CreateValidator[RegisterPostRequest]

// Register godoc
// @securityDefinitions.basic BasicAuth
// @Summary Redirect user to third party register
// @Schemes
// @Description dev: http://localhost:8080/auth/register
// @Description prod: https://csbackend.fly.dev/auth/register
// @Tags auth
// @Accept x-www-form-urlencoded
// @Param access_token query string true "Access token"
// @Param username query string true "Username"
// @Param password query string true "Password"
// @Param role query string true "Role"
// @Produce json
// @Router /auth/register [post]
func POST(c *fiber.Ctx) error {
	var rawReqQuery RegisterPostRequest

	// convert query to struct
	err := c.QueryParser(&rawReqQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// validate query
	errors := registerPostValidator(rawReqQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	userinfoString, err := global.Authenticator.GetUserinfo(rawReqQuery.AccessToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	var userInfo authenticator.UserInfo
	err = json.Unmarshal([]byte(userinfoString), &userInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.SendString("ok")
}
