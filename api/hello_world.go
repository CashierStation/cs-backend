package api

import "github.com/gin-gonic/gin"

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/ping [get]
func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
