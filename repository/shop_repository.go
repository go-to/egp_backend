package repository

import (
	"github.com/go-to/egp_backend/model"
	"time"
)

type IShopRepository interface {
	GetShops(now *time.Time) (*[]model.Shop, error)
}

type ShopRepository struct {
	model model.ShopModel
}

func NewShopRepository(m model.ShopModel) *ShopRepository {
	return &ShopRepository{model: m}
}

func (s *ShopRepository) GetShops(t *time.Time) (*model.ShopsResult, error) {
	return s.model.Find(t)
}
