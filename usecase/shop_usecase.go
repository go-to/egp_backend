package usecase

import (
	"fmt"
	"github.com/go-to/egp_backend/repository"
	"github.com/go-to/egp_backend/usecase/input"
	"github.com/go-to/egp_backend/usecase/output"
	"github.com/go-to/egp_protobuf/pb"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"slices"
)

type IShopUsecase interface {
	GetShops(in *input.ShopsInput) (*output.ShopsOutput, error)
}

type ShopUsecase struct {
	config repository.IConfigRepository
	shop   repository.IShopRepository
}

func NewShopUseCase(config repository.ConfigRepository, shop repository.ShopRepository) *ShopUsecase {
	return &ShopUsecase{
		config: &config,
		shop:   &shop,
	}
}

func (u *ShopUsecase) GetShops(in *input.ShopsInput) (*output.ShopsOutput, error) {
	userId := in.ShopsRequest.GetUserId()
	searchTypes := in.ShopsRequest.GetSearchTypes()
	var searchParams []int32
	var orderParams []int32
	for _, value := range searchTypes {
		v := int32(value)
		if slices.Contains(searchParams, v) {
			continue
		}
		if key, exists := pb.SearchType_name[v]; exists {
			searchParams = append(searchParams, pb.SearchType_value[key])
		}
	}

	now, err := u.config.GetTime()
	if err != nil {
		return &output.ShopsOutput{}, err
	}

	shops, err := u.shop.GetShops(&now, userId, searchParams, orderParams)
	if err != nil {
		return &output.ShopsOutput{}, err
	}

	var outputShops []*pb.Shop
	var latLonList []string
	fmtX := message.NewPrinter(language.Japanese)

	for _, v := range *shops {
		inCurrentSales := true
		if len(v.StartTime) == 0 || len(v.EndTime) == 0 {
			inCurrentSales = false
		}
		// 緯度経度が同じ場合は、重なり防止のためにマーカーの位置をずらす
		latitude := v.Latitude
		longitude := v.Longitude
		latLon := fmt.Sprintf("%f,%f", latitude, longitude)
		if slices.Contains(latLonList, latLon) {
			latitude += 0.00002
			longitude += 0.00002
		}
		latLonList = append(latLonList, latLon)

		// 距離
		distance := fmtX.Sprintf("%dm", int(v.Distance))

		isStamped := false
		if v.NumberOfTimes > 0 {
			isStamped = true
		}

		outputShops = append(outputShops, &pb.Shop{
			ID:                         v.ID,
			EventID:                    v.EventID,
			Year:                       v.Year,
			CategoryID:                 pb.CategoryType(v.CategoryID),
			CategoryName:               v.CategoryName,
			No:                         v.No,
			ShopName:                   v.ShopName,
			MenuName:                   v.MenuName,
			MenuImageUrl:               v.MenuImageUrl,
			Phone:                      v.Phone,
			Address:                    v.Address,
			BusinessDays:               v.BusinessDays,
			RegularHoliday:             v.RegularHoliday,
			BusinessHours:              v.BusinessHours,
			ChargePrice:                v.ChargePrice,
			NormalizedChargePrice:      v.NormalizedChargePrice,
			SinglePrice:                v.SinglePrice,
			NormalizedSinglePrice:      v.NormalizedSinglePrice,
			SetPrice:                   v.SetPrice,
			NormalizedSetPrice:         v.NormalizedSetPrice,
			BeerType:                   v.BeerType,
			NeedsReservation:           v.NeedsReservation,
			NormalizedNeedsReservation: v.NormalizedNeedsReservation,
			UseHachipay:                v.UseHachipay,
			NormalizedUseHachipay:      v.NormalizedUseHachipay,
			IsOpenHoliday:              v.IsOpenHoliday,
			IsIrregularHoliday:         v.IsIrregularHoliday,
			Latitude:                   latitude,
			Longitude:                  longitude,
			Distance:                   distance,
			WeekNumber:                 v.WeekNumber,
			DayOfWeek:                  int32(v.DayOfWeek),
			StartTime:                  v.StartTime,
			EndTime:                    v.EndTime,
			IsHoliday:                  v.IsHoliday,
			InCurrentSales:             inCurrentSales,
			NumberOfTimes:              v.NumberOfTimes,
			IsStamped:                  isStamped,
		})
	}

	return &output.ShopsOutput{
		ShopsResponse: pb.ShopsResponse{
			Shops: outputShops,
		},
	}, nil
}
