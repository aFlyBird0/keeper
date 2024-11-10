package router

import (
	"keeper/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/health", handler.Health)
	itemGroup := r.Group("/item")
	{
		itemGroup.GET("/listAll", handler.ListItems)
		itemGroup.GET("/listExpiredItems", handler.ListExpiredItems)
		itemGroup.GET("/listByName", handler.ListItemsByName)
		itemGroup.POST("/add", handler.AddItem)
		itemGroup.GET("/error", handler.JustShowError)
	}
}
