package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Callback godoc
// @Summary Endpoint the user is redirected to after logging in.
// @Schemes
// @Description Callback
// @Tags auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param request query auth.Callback.request true "query params"
// @Success 200 {object} auth.Callback.response
// @Router /auth/callback [get]
func (a *Auth) Callback(c *gin.Context) {
	type request struct{}

	type response struct{}

	var req request

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	if c.Query("state") != session.Get("state") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state parameter"})
		return
	}

	// Exchange an authorization code for a token.
	token, err := a.Exchange(c.Request.Context(), c.Query("code"))
	if err != nil {
		c.String(http.StatusUnauthorized, "Failed to exchange an authorization code for a token.")
		return
	}

	idToken, err := a.VerifyIDToken(c.Request.Context(), token)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to verify ID Token.")
		return
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	session.Set("access_token", token.AccessToken)
	session.Set("profile", profile)
	if err := session.Save(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Redirect to logged in page.
	c.Redirect(http.StatusTemporaryRedirect, "/user")

	c.JSON(200, response{})
}
