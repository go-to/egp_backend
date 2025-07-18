package model

import (
	"errors"
	"fmt"
	"github.com/go-to/egp_backend/util"
	"gorm.io/gorm"
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
var shopResult ShopDetail

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
	CountShopsTotal(year int32) (int64, error)
	FindShops(time *time.Time, userId string, year int32, keywordList []string, searchParams []int32, orderParam int32, latitude, longitude float64) (*ShopsResult, error)
	FindShop(time *time.Time, userId string, shopId int64) (*ShopDetail, error)
}

// search types
const (
	SearchTypeInCurrentSales = iota
	SearchTypeNotYet
	SearchTypeIrregularHoliday
	SearchTypeNeedsReservation
	SearchTypeBeerCocktail
)

// sort order
const (
	SortOrderNo = iota
	SortOrderDistance
)

type ShopModel struct {
	db DB
}

func NewShopModel(db DB) *ShopModel {
	return &ShopModel{db: db}
}

func (m *ShopModel) CountShopsTotal(year int32) (int64, error) {
	count := int64(0)
	if err := m.db.Conn.
		Model(&Shop{}).
		Joins("INNER JOIN events ON shops.event_id = events.id").
		Where("events.year = ?", year).
		Count(&count).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return int64(-1), err
	}

	return count, nil
}

func (m *ShopModel) FindShops(time *time.Time, userId string, year int32, keywordParams []string, searchParams []int32, orderParam int32, latitude, longitude float64) (*ShopsResult, error) {
	stDistance := fmt.Sprintf("ST_Distance(shops_location.location, 'POINT(%f %f)', false)", latitude, longitude)

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
	shopsTimeTodayCondition := "week_number = ? AND day_of_week = ? AND is_holiday = false AND start_time <= ? AND end_time >= ?"
	shopsTimeTodayConditionWithPrefix := "shops_time_day.week_number = ? AND shops_time_day.day_of_week = ? AND shops_time_day.is_holiday = false AND shops_time_day.start_time <= ? AND shops_time_day.end_time >= ?"
	shopsTimeTomorrowCondition := "week_number = ? AND day_of_week = ? AND is_holiday = false AND ? - INTERVAL '12 hour' <= '00:00:00' AND start_time <= ? AND end_time >= ?"
	shopsTimeTomorrowConditionWithPrefix := "shops_time_night.week_number = ? AND shops_time_night.day_of_week = ? AND shops_time_night.is_holiday = false AND ? - INTERVAL '12 hour' <= '00:00:00' AND shops_time_night.start_time <= ? AND shops_time_night.end_time >= ?"

	query := m.db.Conn.
		Model(&Shop{}).
		Select(fields).
		Joins("INNER JOIN events ON shops.event_id = events.id").
		Joins("INNER JOIN categories ON shops.category_id = categories.id").
		Joins("INNER JOIN shops_location ON shops.id = shops_location.shop_id").
		Joins("LEFT JOIN (SELECT shop_id, week_number, day_of_week, start_time, end_time, is_holiday FROM shops_time WHERE "+shopsTimeTodayCondition+") AS shops_time_day ON shops.id = shops_time_day.shop_id", todayWeekNum, todayDayOfWeek, nowTime, nowTime).
		Joins("LEFT JOIN (SELECT shop_id, week_number, day_of_week, start_time, end_time, is_holiday FROM shops_time WHERE "+shopsTimeTomorrowCondition+") AS shops_time_night ON shops.id = shops_time_night.shop_id", tomorrowWeekNum, tomorrowDayOfWeek, nowTime, nowTime, nowTime).
		Joins("LEFT JOIN (SELECT shop_id, number_of_times FROM stamps WHERE user_id = ? AND deleted_at IS NULL) AS stamps ON shops.id = stamps.shop_id", userId).
		Where("events.year = ?", year)

	/* 検索条件の指定があれば、検索条件を追加 */
	// 営業中の店舗で絞り込む
	if slices.Contains(searchParams, SearchTypeInCurrentSales) {
		query = query.Where("("+shopsTimeTodayConditionWithPrefix+") OR ("+shopsTimeTomorrowConditionWithPrefix+")",
			todayWeekNum, todayDayOfWeek, nowTime, nowTime,
			tomorrowWeekNum, tomorrowDayOfWeek, nowTime, nowTime, nowTime)
	}
	// スタンプ未獲得の店舗で絞り込む
	if slices.Contains(searchParams, SearchTypeNotYet) {
		query = query.Where("stamps.number_of_times IS NULL")
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
	// キーワード検索
	if len(keywordParams) > 0 {
		keywordQuery := m.db.Conn.Or("")
		for _, keyword := range keywordParams {
			keywordQuery = keywordQuery.Or("sf_translate_case(shops.no) LIKE ? || sf_translate_case(?) || ? OR sf_translate_case(shops.shop_name) LIKE ? || sf_translate_case(?) || ? OR sf_translate_case(shops.menu_name) LIKE ? || sf_translate_case(?) || ?",
				"%", keyword, "%", "%", keyword, "%", "%", keyword, "%")
		}
		query = query.Where(keywordQuery)
	}

	/* ソート */
	// No.順
	if orderParam == SortOrderNo {
		query = query.Order("CAST(shops.no AS INTEGER) ASC")
	}
	// 距離順
	if orderParam == SortOrderDistance {
		query = query.Order("distance ASC")
		query = query.Order("CAST(shops.no AS INTEGER) ASC")
	}

	// クエリ実行
	shopsResult = nil
	res := query.Scan(&shopsResult)
	if res.Error != nil {
		return nil, res.Error
	}

	return &shopsResult, nil
}

func (m *ShopModel) FindShop(time *time.Time, userId string, shopId int64) (*ShopDetail, error) {
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
		Where("shops.id = ?", shopId)

	shopResult = ShopDetail{}
	res := query.Scan(&shopResult)
	if res.Error != nil {
		return nil, res.Error
	}

	return &shopResult, nil
}
