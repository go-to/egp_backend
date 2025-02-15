package model

import (
	"fmt"
	"github.com/go-to/egp_backend/util"
	"slices"
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
	Year                       int32
	CategoryID                 int64
	CategoryName               string
	No                         int32
	ShopName                   string
	MenuName                   string
	MenuImageUrl               string
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
	Distance                   float64
	WeekNumber                 int32
	DayOfWeek                  time.Weekday
	StartTime                  string
	EndTime                    string
	IsHoliday                  bool
	InCurrentSales             bool
	NumberOfTimes              int32
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
	Find(time *time.Time, userId string, searchParams []int32, orderParams []int32) (*ShopsResult, error)
}

const (
	SearchTypeInCurrentSales = iota
	SearchTypeNotYet
	SearchTypeIrregularHoliday
	SearchTypeNeedsReservation
	SearchTypeBeerCocktail
)

type ShopModel struct {
	db DB
}

func NewShopModel(db DB) *ShopModel {
	return &ShopModel{db: db}
}

func (m *ShopModel) Find(time *time.Time, userId string, searchParams []int32, orderParams []int32) (*ShopsResult, error) {
	lat := 35.64531919787909
	lng := 139.7223368970176
	stDistance := fmt.Sprintf("ST_Distance(shops_location.location, 'POINT(%f %f)', false)", lat, lng)

	fields := `
		shops.id,
		shops.event_id,
		events.year,
		shops.category_id,
		categories.name AS category_name,
		shops.no,
		shops.shop_name,
		shops.menu_name,
		shops.menu_image_url,
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
		` + stDistance + ` AS distance,
		CASE
			WHEN shops_time_day.week_number IS NOT NULL THEN shops_time_day.week_number 
			WHEN shops_time_night.week_number IS NOT NULL THEN shops_time_night.week_number
			ELSE NULL
		END AS week_number,
		CASE
			WHEN shops_time_day.day_of_week IS NOT NULL THEN shops_time_day.day_of_week 
			WHEN shops_time_night.day_of_week IS NOT NULL THEN shops_time_night.day_of_week
			ELSE NULL
		END AS day_of_week,
		CASE
			WHEN shops_time_day.start_time IS NOT NULL THEN shops_time_day.start_time 
			WHEN shops_time_night.start_time IS NOT NULL THEN shops_time_night.start_time
			ELSE NULL
		END AS start_time,
		CASE
			WHEN shops_time_day.end_time IS NOT NULL THEN shops_time_day.end_time 
			WHEN shops_time_night.end_time IS NOT NULL THEN shops_time_night.end_time
			ELSE NULL
		END AS end_time,
		CASE
			WHEN shops_time_day.is_holiday IS NOT NULL THEN shops_time_day.is_holiday 
			WHEN shops_time_night.is_holiday IS NOT NULL THEN shops_time_night.is_holiday
			ELSE NULL
		END AS is_holiday,
		stamps.number_of_times
	`

	// 検索条件で指定する週番号、曜日、時刻の情報を取得
	todayWeekNum := util.GetWeekNumber(time)
	todayDayOfWeek := util.GetWeekDay(time)
	tomorrow := time.AddDate(0, 0, 1)
	tomorrowWeekNum := util.GetWeekNumber(&tomorrow)
	tomorrowDayOfWeek := util.GetWeekDay(&tomorrow)
	nowTime := util.GetTime(time)
	shopsTimeTodayCondition := "shops_time_day.week_number = ? AND shops_time_day.day_of_week = ? AND shops_time_day.is_holiday = false AND shops_time_day.start_time <= ? AND shops_time_day.end_time >= ?"
	shopsTimeTomorrowCondition := "shops_time_night.week_number = ? AND shops_time_night.day_of_week = ? AND shops_time_night.is_holiday = false AND ? - INTERVAL '12 hour' <= '00:00:00' AND shops_time_night.start_time <= ? AND shops_time_night.end_time >= ?"

	query := m.db.Conn.
		Model(&Shop{}).
		Select(fields).
		Joins("INNER JOIN events ON shops.event_id = events.id").
		Joins("INNER JOIN categories ON shops.category_id = categories.id").
		Joins("INNER JOIN shops_location ON shops.id = shops_location.shop_id").
		Joins("LEFT JOIN shops_time AS shops_time_day ON shops.id = shops_time_day.shop_id AND "+shopsTimeTodayCondition+"",
			todayWeekNum, todayDayOfWeek, nowTime, nowTime).
		Joins("LEFT JOIN shops_time AS shops_time_night ON shops.id = shops_time_night.shop_id AND "+shopsTimeTomorrowCondition+"",
			tomorrowWeekNum, tomorrowDayOfWeek, nowTime, nowTime, nowTime).
		Joins("LEFT JOIN stamps ON shops.id = stamps.shop_id AND stamps.user_id = ? AND stamps.deleted_at IS NULL", userId).
		Order("shops.no ASC")

	/* 検索条件の指定があれば、検索条件を追加 */
	// 営業中の店舗で絞り込む
	if slices.Contains(searchParams, SearchTypeInCurrentSales) {
		query = query.Where("("+shopsTimeTodayCondition+") OR ("+shopsTimeTomorrowCondition+")",
			todayWeekNum, todayDayOfWeek, nowTime, nowTime,
			tomorrowWeekNum, tomorrowDayOfWeek, nowTime, nowTime, nowTime)
	}
	// スタンプ未獲得の店舗で絞り込む
	if slices.Contains(searchParams, SearchTypeNotYet) {
		// TODO スタンプ機能を実装後に追加
		query = query.Where("1 = 1")
	}
	// 不定休の店舗で絞り込む
	if slices.Contains(searchParams, SearchTypeIrregularHoliday) {
		query = query.Where("shops.is_irregular_holiday = ?", true)
	}
	// 予約が必要な店舗で絞り込む
	if slices.Contains(searchParams, SearchTypeNeedsReservation) {
		query = query.Where("shops.normalized_needs_reservation = ?", true)
	}
	// ビールカクテルがある店舗で絞り込む
	if slices.Contains(searchParams, SearchTypeBeerCocktail) {
		query = query.Where("shops.category_id = ?", CATEGORY_BEER_COCKTAIL)
	}

	// クエリ実行
	shopsResult = nil
	res := query.Scan(&shopsResult)
	if res.Error != nil {
		return nil, res.Error
	}

	return &shopsResult, nil
}
