package db

import (
	"keeper/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// gorm教程：https://gorm.io/zh_CN/docs/

func InitDB() {
	// 用户名:密码@tcp(地址:端口)/数据库名?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "bird:keeper123456@tcp(127.0.0.1:3306)/keeper?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	// 连接数据库
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 自动迁移表（初始化数据库表）
	err = db.AutoMigrate(&model.Item{})
	if err != nil {
		panic("failed to migrate database")
	}
}

func DB() *gorm.DB {
	if db == nil {
		panic("db is not initialized")
	}
	return db
}
