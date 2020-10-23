package deliveries

import "github.com/gin-gonic/gin"

func pingController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
