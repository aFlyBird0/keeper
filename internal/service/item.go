package service

import (
	"keeper/internal/model"
	"time"

	"github.com/gin-gonic/gin"
)

func ListItems(c *gin.Context) {
	item := model.Item{
		Name:        "苹果",
		Amount:      3,
		Quantifier:  "个",
		Place:       "桌子上",
		ExpiredAt:   time.Now().Add(time.Hour * 24 * 3),
		Description: "大大大大苹果",
	}
	c.JSON(200, gin.H{
		"data": []model.Item{item},
	})
}
