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

func AddItem(c *gin.Context) {
	item := new(model.Item)
	// 完整的参数绑定教程
	// 1. 绑定HTTP请求参数
	// 	1.1 request body到结构体：https://gin-gonic.com/zh-cn/docs/examples/bind-body-into-dirrerent-structs/
	// 	1.2 绑定uri：https://gin-gonic.com/zh-cn/docs/examples/bind-uri/
	// 	1.3 绑定query string：https://gin-gonic.com/zh-cn/docs/examples/query-and-post-form/

	// 2. 验证参数：https://gin-gonic.com/zh-cn/docs/examples/binding-and-validation/
	err := c.ShouldBindJSON(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.ParamError))
		return
	}
	c.JSON(http.StatusOK, response.EmptySuccess())
}

func JustShowError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, response.Fail(response.ExampleError))
}
