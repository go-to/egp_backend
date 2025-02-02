package repository

import (
	"github.com/go-to/egp_backend/model"
	"time"
)

type IShopRepository interface {
	GetShops(t *time.Time, s []int32, o []int32) (*model.ShopsResult, error)
}

type ShopRepository struct {
	model model.ShopModel
}

func NewShopRepository(m model.ShopModel) *ShopRepository {
	return &ShopRepository{model: m}
}

func (r *ShopRepository) GetShops(t *time.Time, s []int32, o []int32) (*model.ShopsResult, error) {
	return r.model.Find(t, s, o)
}
