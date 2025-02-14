package repository

import (
	"github.com/go-to/egp_backend/model"
	"time"
)

type IShopRepository interface {
	GetShops(time *time.Time, userId string, searchParams []int32, orderParams []int32) (*model.ShopsResult, error)
}

type ShopRepository struct {
	model model.ShopModel
}

func NewShopRepository(m model.ShopModel) *ShopRepository {
	return &ShopRepository{model: m}
}

func (r *ShopRepository) GetShops(time *time.Time, userId string, searchParams []int32, orderParams []int32) (*model.ShopsResult, error) {
	return r.model.Find(time, userId, searchParams, orderParams)
}
