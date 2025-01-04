package model

type Shop struct {
	ID       int64  `gorm:"primary_key"`
	ShopName string `gorm:"type:varchar(255)"`
}

var shops []Shop

type IShopModel interface {
	Find() (*[]Shop, error)
}

type ShopModel struct {
	db DB
}

func NewShopModel(db DB) *ShopModel {
	return &ShopModel{db: db}
}

func (m *ShopModel) Find() (*[]Shop, error) {
	result := m.db.Conn.Find(&shops)
	if result.Error != nil {
		return nil, result.Error
	}

	return &shops, nil
}
