package router

import (
	"keeper/internal/handler"
	"keeper/internal/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/health", handler.Health)
	// 令牌桶限流，容量5，默认3秒填充一次令牌，每次填充2，桶里没有令牌了就会限流
	r.Use(middleware.RateLimitMiddleware(time.Second*3, 5, 2))
	itemGroup := r.Group("/item")
	itemGroup.Use(middleware.SimpleUser)
	{
		itemGroup.GET("/listAll", handler.ListItems)
		itemGroup.GET("/listExpiredItems", handler.ListExpiredItems)
		itemGroup.GET("/listByName", handler.ListItemsByName)
		itemGroup.POST("/add", handler.AddItem)
		itemGroup.GET("/error", handler.JustShowError)
	}
}
