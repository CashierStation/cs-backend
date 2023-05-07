package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Login godoc
// @securityDefinitions.basic BasicAuth
// @Summary Redirect user to third party login
// @Schemes
// @Description dev: http://localhost:8080/auth/login
// @Description prod: https://cs-backend-production.up.railway.app/auth/login
// @Tags auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param request query auth.Login.request true "query params"
// @Router /auth/login [get]
func (a *Auth) Login(c *gin.Context) {
	type request struct{}

	//type response struct{}

	var req request

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	state, err := generateRandomState()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Save state in session
	session := sessions.Default(c)
	session.Set("state", state)

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Redirect user to third party login
	c.Redirect(http.StatusTemporaryRedirect, a.Authenticator.AuthCodeURL(state))
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
