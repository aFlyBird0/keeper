package handler

import "github.com/gin-gonic/gin"

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, this is Keeper!",
	})
}
