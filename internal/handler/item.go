package handler

import (
	"context"
	"keeper/internal/model"
	"keeper/internal/service"
	"keeper/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListItems(c *gin.Context) {
	items, err := service.NewItemService().ListItems(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.DatabseError)
		return
	}
	c.JSON(http.StatusOK, response.Success(items))
}

func ListExpiredItems(c *gin.Context) {
	items, err := service.NewItemService().ListExpiredItems(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.DatabseError)
		return
	}
	c.JSON(http.StatusOK, response.Success(items))
}

func ListItemsByName(c *gin.Context) {
	name := c.Query("name")
	items, err := service.NewItemService().FindItemsByName(context.Background(), name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.DatabseError)
		return
	}
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

	// 填充默认值
	item.FillDefaults()

	// 存入数据库
	res, err := service.NewItemService().CreateItem(context.Background(), item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.DatabseError)
		return
	}
	// 注：可以通过item.ID获取到数据库生成的ID
	c.JSON(http.StatusOK, response.Success(res))
}

func JustShowError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, response.Fail(response.ExampleError))
}
