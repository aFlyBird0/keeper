package middleware

import (
	"keeper/internal/constant"
	"keeper/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SimpleUser(c *gin.Context) {
	// 从 Authorization Header 中提取信息
	user := c.GetHeader("Authorization")
	if user == "" {
		c.JSON(http.StatusUnauthorized, response.Fail(response.NotLogin))
		c.Abort()
	}
	c.Set(constant.UserContextKey, user)
}

func GetUser(c *gin.Context) string {
	return c.GetString(constant.UserContextKey)
}
