package router

import (
	"keeper/internal/service"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/health", service.Health)
}
