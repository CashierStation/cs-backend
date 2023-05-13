package callback

import (
	"csbackend/global"

	"github.com/gofiber/fiber/v2"
)

// Callback godoc
// @Summary Endpoint the user is redirected to after logging in.
// @Schemes
// @Description Callback
// @Tags oauth
// @Accept x-www-form-urlencoded
// @Produce json
// @Router /oauth/callback [get]
func GET(c *fiber.Ctx) error {
	session, err := global.Session.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	state := session.Get("state")
	if state != c.Query("state") {
		return c.SendString("Invalid state parameter")
	}

	token, err := global.Authenticator.Exchange(c.Context(), c.Query("code"))
	if err != nil {
		return c.SendString("Failed to exchange token: " + err.Error())
	}

	idToken, err := global.Authenticator.VerifyIDToken(c.Context(), token)
	if err != nil {
		return c.SendString("Failed to verify ID token: " + err.Error())
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return c.SendString("Failed to parse ID token claims: " + err.Error())
	}

	session.Set("profile", profile)
	if err := session.Save(); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendString(token.AccessToken)
}
