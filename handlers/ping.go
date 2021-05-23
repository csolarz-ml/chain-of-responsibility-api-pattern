package handlers

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.AbortWithStatusJSON(200, gin.H{
		"message": "pong",
	})
}
