package router

import (
	"keeper/internal/service"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/health", service.Health)
	itemGroup := r.Group("/item")
	{
		itemGroup.GET("/listAll", service.ListItems)
		itemGroup.POST("/add", service.AddItem)
		itemGroup.GET("/error", service.JustShowError)
	}
}
