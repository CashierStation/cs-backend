package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type User struct{}

// User godoc
// @Summary
// @Schemes
// @Description User
// @Tags user
// @Accept x-www-form-urlencoded
// @Produce json
// @Param request query api.User.GET.request true "query params"
// @Success 200 {object} api.User.GET.response
// @Router /user [get]
func (l *User) GET(c *gin.Context) {
	type request struct{}

	type response struct {
		Profile interface{} `json:"profile"`
	}

	var req request

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	profile := session.Get("profile")

	c.JSON(200, response{profile})
}
