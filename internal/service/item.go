package service

import (
	"keeper/internal/model"
	"keeper/pkg/response"
	"net/http"
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
	items := []model.Item{item}
	c.JSON(http.StatusOK, response.Success(items))
}

func JustShowError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, response.Fail(response.ExampleError))
}
