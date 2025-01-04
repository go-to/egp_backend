package model

type Shop struct {
	ID                         int64
	No                         int32
	ShopName                   string
	MenuName                   string
	Phone                      string
	Address                    string
	BusinessDays               string
	RegularHoliday             string
	BusinessHours              string
	ChargePrice                string
	NormalizedChargePrice      int32
	SinglePrice                string
	NormalizedSinglePrice      int32
	SetPrice                   string
	NormalizedSetPrice         int32
	BeerType                   string
	NeedsReservation           string
	NormalizedNeedsReservation bool
	UseHachipay                string
	NormalizedUseHachipay      bool
	IsOpenHoliday              bool
	IsIrregularHoliday         bool
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
