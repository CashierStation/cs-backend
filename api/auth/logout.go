package auth

import (
	"net/http"
	"net/url"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Logout godoc
// @securityDefinitions.basic BasicAuth
// @Summary Log user out
// @Schemes
// @Description Logout
// @Tags auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param request query auth.Logout.request true "query params"
// @Router /auth/logout [get]
func (a *Auth) Logout(c *gin.Context) {
	type request struct{}

	//type response struct{}

	var req request

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Clear()
	session.Save()

	logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + c.Request.Host + "/user")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	params := url.Values{}
	params.Add("returnTo", returnTo.String())
	params.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = params.Encode()

	c.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
}
