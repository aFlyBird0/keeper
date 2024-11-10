package model

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name string `gorm:"primary_key" binding:"required"`
	// 数量
	Amount int `gorm:"not null,default:1"`
	// 量词
	Quantifier  string
	Place       string `gorm:"not null,default:''"`
	ExpiredAt   *time.Time
	Description string
}

// FillDefaults 填充默认值
func (item *Item) FillDefaults() {
	if item.Amount == 0 {
		item.Amount = 1
	}
}
