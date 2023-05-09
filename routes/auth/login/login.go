package login

import (
	"crypto/rand"
	"csbackend/global"
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
)

// Login godoc
// @securityDefinitions.basic BasicAuth
// @Summary Redirect user to third party login
// @Schemes
// @Description dev: http://localhost:8080/auth/login
// @Description prod: https://csbackend.fly.dev/auth/login
// @Tags auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Router /auth/login [get]
func GET(c *fiber.Ctx) error {
	state, err := generateRandomState()

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

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
