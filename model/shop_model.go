package model

import (
	"github.com/go-to/egp_backend/util"
	"time"
)

type Shop struct {
	ID                         int64
	EventID                    int64
	CategoryID                 int64
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

type ShopsLocation struct {
	ID        int64
	ShopID    int64
	Latitude  float64
	Longitude float64
	Location  string
}

type ShopsTime struct {
	ID         int64
	ShopID     int64
	WeekNumber int32
	DayOfWeek  time.Weekday
	StartTime  string
	EndTime    string
	IsHoliday  int32
}

type ShopDetail struct {
	ID                         int64
	EventID                    int64
	CategoryID                 int64
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
	Latitude                   float64
	Longitude                  float64
	Location                   string
	WeekNumber                 int32
	DayOfWeek                  time.Weekday
	StartTime                  string
	EndTime                    string
	IsHoliday                  bool
	InCurrentSales             bool
}

type ShopsResult []ShopDetail

var shopsResult ShopsResult

func (Shop) TableName() string {
	return "shops"
}

func (ShopsLocation) TableName() string {
	return "shops_location"
}

func (ShopsTime) TableName() string {
	return "shops_time"
}

type IShopModel interface {
	Find(t *time.Time) (*ShopsResult, error)
}

type ShopModel struct {
	db DB
}

func NewShopModel(db DB) *ShopModel {
	return &ShopModel{db: db}
}

func (m *ShopModel) Find(t *time.Time) (*ShopsResult, error) {
	fields := `
shops.id,
shops.event_id,
shops.category_id,
shops.no,
shops.shop_name,
shops.menu_name,
shops.phone,
shops.address,
shops.business_days,
shops.regular_holiday,
shops.business_hours,
shops.charge_price,
shops.normalized_charge_price,
shops.single_price,
shops.normalized_single_price,
shops.set_price,
shops.normalized_set_price,
shops.beer_type,
shops.needs_reservation,
shops.normalized_needs_reservation,
shops.use_hachipay,
shops.normalized_use_hachipay,
shops.is_open_holiday,
shops.is_irregular_holiday,
shops_location.latitude,
shops_location.longitude,
shops_location.location,
shops_time.week_number,
shops_time.day_of_week,
shops_time.start_time,
shops_time.end_time,
shops_time.is_holiday
`
	// 検索条件で指定する週番号、曜日、時刻の情報を取得
	todayWeekNum := util.GetWeekNumber(t)
	todayDayOfWeek := util.GetWeekDay(t)
	tomorrow := t.AddDate(0, 0, 1)
	tomorrowWeekNum := util.GetWeekNumber(&tomorrow)
	tomorrowDayOfWeek := util.GetWeekDay(&tomorrow)
	nowTime := util.GetTime(t)

	shopsTimeTodayCondition := "shops_time.week_number = ? AND shops_time.day_of_week = ? AND shops_time.is_holiday = false AND shops_time.start_time <= ? AND shops_time.end_time >= ?"
	shopsTimeTomorrowCondition := "shops_time.week_number = ? AND shops_time.day_of_week = ? AND shops_time.is_holiday = false AND ? - INTERVAL '12 hour' <= '00:00:00' AND shops_time.start_time <= ? AND shops_time.end_time >= ?"

	res := m.db.Conn.
		Model(&Shop{}).
		Select(fields).
		Joins("INNER JOIN shops_location ON shops.id = shops_location.shop_id").
		Joins("LEFT JOIN shops_time ON shops.id = shops_time.shop_id AND ("+shopsTimeTodayCondition+" OR "+shopsTimeTomorrowCondition+")",
			todayWeekNum, todayDayOfWeek, nowTime, nowTime,
			tomorrowWeekNum, tomorrowDayOfWeek, nowTime, nowTime, nowTime).
		Order("shops.no ASC").
		Scan(&shopsResult)
	if res.Error != nil {
		return nil, res.Error
	}

	return &shopsResult, nil
}
