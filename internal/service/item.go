package service

import (
	"context"
	"keeper/internal/db"
	"keeper/internal/model"
	"time"

	"gorm.io/gorm"
)

type ItemService interface {
	CreateItem(ctx context.Context, item *model.Item) (*model.Item, error)
	ListItems(ctx context.Context) ([]*model.Item, error)
	FindItemsByName(ctx context.Context, name string) ([]*model.Item, error)
	ListExpiredItems(ctx context.Context) ([]*model.Item, error)
}

type itemService struct {
	db *gorm.DB
}

func (i *itemService) CreateItem(ctx context.Context, item *model.Item) (*model.Item, error) {
	res := i.db.Create(item)
	if res.Error != nil {
		return nil, res.Error
	}
	return item, nil
}

func (i *itemService) ListItems(ctx context.Context) ([]*model.Item, error) {
	var items []*model.Item
	res := i.db.Find(&items).Order("expired_at asc")
	if res.Error != nil {
		return nil, res.Error
	}
	return items, nil
}

func (i *itemService) FindItemsByName(ctx context.Context, name string) ([]*model.Item, error) {
	var items []*model.Item
	res := i.db.Where("name LIKE ?", "%"+name+"%").Find(&items)
	if res.Error != nil {
		return nil, res.Error
	}
	return items, nil
}

func (i *itemService) ListExpiredItems(ctx context.Context) ([]*model.Item, error) {
	var items []*model.Item
	res := i.db.Where("expired_at < ?", time.Now()).Find(&items)
	if res.Error != nil {
		return nil, res.Error
	}
	return items, nil
}

func NewItemService() ItemService {
	return NewItemServiceWithDB(db.DB())
}

func NewItemServiceWithDB(db *gorm.DB) ItemService {
	return &itemService{db: db}
}
