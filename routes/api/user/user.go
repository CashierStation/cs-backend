package user

import (
	"csbackend/global"

	"github.com/gofiber/fiber/v2"
)

// @Security ApiKeyAuth
// User godoc
// @Summary
// @Schemes
// @Description User
// @Tags user
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id_token query string true "ID Token"
// @Success 200 {object} user.GET.response
// @Router /api/user [get]
func GET(c *fiber.Ctx) error {
	type response struct {
		Aud            string `json:"aud"`
		Email          string `json:"email"`
		Email_verified bool   `json:"email_verified"`
		Exp            int    `json:"exp"`
		Family_name    string `json:"family_name"`
		Given_name     string `json:"given_name"`
		Iat            int    `json:"iat"`
		Iss            string `json:"iss"`
		Locale         string `json:"locale"`
		Name           string `json:"name"`
		Nickname       string `json:"nickname"`
		Picture        string `json:"picture"`
		Sid            string `json:"sid"`
		Sub            string `json:"sub"`
		Updated_at     string `json:"updated_at"`
	}

	id_token := c.Query("id_token")
	oidc_thing, err := global.Authenticator.VerifyRawIDToken(c.Context(), id_token)

	if err != nil {
		return c.SendString("Failed to verify ID token: " + err.Error())
	}

	var profile response
	if err := oidc_thing.Claims(&profile); err != nil {
		return c.SendString("Failed to parse ID token claims: " + err.Error())
	}

	return c.JSON(profile)

	/* sess, err := global.Session.Get(c)
	if err != nil {
		return err
	}

	profile := sess.Get("profile")

	return c.JSON(profile) */
}
