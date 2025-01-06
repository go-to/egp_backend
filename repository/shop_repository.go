package repository

import (
	"github.com/go-to/egp_backend/model"
)

type IShopRepository interface {
	GetShops() (*[]model.Shop, error)
}

type ShopRepository struct {
	model model.ShopModel
}

func NewShopRepository(m model.ShopModel) *ShopRepository {
	return &ShopRepository{model: m}
}

func (s *ShopRepository) GetShops() (*[]model.Shop, error) {
	return s.model.Find()
}
