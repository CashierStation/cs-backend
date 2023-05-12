package logout

import (
	"csbackend/global"
	"net/http"
	"net/url"

	"os"

	"github.com/gofiber/fiber/v2"
)

// Logout godoc
// @securityDefinitions.basic BasicAuth
// @Summary Log user out
// @Schemes
// @Description dev: http://localhost:8080/oauth/logout
// @Description prod: https://csbackend.fly.dev/oauth/logout
// @Tags oauth
// @Accept x-www-form-urlencoded
// @Produce json
// @Router /oauth/logout [get]
func GET(c *fiber.Ctx) error {
	session, err := global.Session.Get(c)
	if err != nil {
		return err
	}

	session.Destroy()
	session.Save()

	logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	scheme := "http"
	if c.Secure() {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + c.Hostname())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	params := url.Values{}
	params.Add("returnTo", returnTo.String())
	params.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = params.Encode()

	return c.Redirect(logoutUrl.String(), http.StatusTemporaryRedirect)
}
