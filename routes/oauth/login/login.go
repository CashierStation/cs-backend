package login

import (
	"csbackend/global"

	"github.com/gofiber/fiber/v2"
)

// Login godoc
// @securityDefinitions.basic BasicAuth
// @Summary Redirect user to third party login
// @Schemes
// @Description dev: http://localhost:8080/oauth/login
// @Description prod: https://csbackend.fly.dev/oauth/login
// @Tags oauth
// @Accept x-www-form-urlencoded
// @Produce json
// @Router /oauth/login [get]
func GET(c *fiber.Ctx) error {
	state, err := global.Authenticator.GenerateRandomBase64()

	if err != nil {
		return err
	}

	// Save state in session
	sess, err := global.Session.Get(c)

	if err != nil {
		return c.SendString(err.Error())
	}

	sess.Set("state", state)
	err = sess.Save()
	if err != nil {
		return c.SendString(err.Error())
	}

	// Redirect to auth endpoint
	return c.Redirect(global.Authenticator.AuthCodeURL(state), fiber.StatusTemporaryRedirect)
}
