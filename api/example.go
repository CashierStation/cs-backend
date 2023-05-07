package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Example struct{}

// Example godoc
// @Summary
// @Schemes
// @Description Example
// @Tags example
// @Accept x-www-form-urlencoded
// @Produce json
// @Param request query api.Example.GET.request true "query params"
// @Success 200 {object} api.Example.GET.response
// @Router /example [get]
func (l *Example) GET(c *gin.Context) {
	type request struct{}

	type response struct{}

	var req request

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response{})
}
