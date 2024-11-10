package main

import (
	"keeper/internal/db"
	"keeper/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()
	router.InitRouter(r)
	err := r.Run()
	if err != nil {
		panic("服务启动失败")
	} // 监听并在 0.0.0.0:8080 上启动服务
}
