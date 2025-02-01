package usecase

import (
	"fmt"
	"github.com/go-to/egp_backend/repository"
	"github.com/go-to/egp_backend/usecase/input"
	"github.com/go-to/egp_backend/usecase/output"
	"github.com/go-to/egp_protobuf/pb"
	"slices"
)

type IShopUsecase interface {
	GetShops(in *input.ShopsInput) (*output.ShopsOutput, error)
}

type ShopUsecase struct {
	config repository.ConfigRepository
	shop   repository.ShopRepository
}

func NewShopUseCase(config repository.ConfigRepository, shop repository.ShopRepository) *ShopUsecase {
	return &ShopUsecase{config: config, shop: shop}
}

func (u *ShopUsecase) GetShops(in *input.ShopsInput) (*output.ShopsOutput, error) {

	now, err := u.config.GetTime()
	if err != nil {
		return &output.ShopsOutput{}, err
	}

	shops, err := u.shop.GetShops(&now)
	if err != nil {
		return &output.ShopsOutput{}, err
	}

	var outputShops []*pb.Shop
	var latLonList []string

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

		outputShops = append(outputShops, &pb.Shop{
			ID:                         v.ID,
			EventID:                    v.EventID,
			CategoryID:                 v.CategoryID,
			No:                         v.No,
			ShopName:                   v.ShopName,
			MenuName:                   v.MenuName,
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
			WeekNumber:                 v.WeekNumber,
			DayOfWeek:                  int32(v.DayOfWeek),
			StartTime:                  v.StartTime,
			EndTime:                    v.EndTime,
			IsHoliday:                  v.IsHoliday,
			InCurrentSales:             inCurrentSales,
		})
	}

	return &output.ShopsOutput{
		ShopsResponse: pb.ShopsResponse{
			Shops: outputShops,
		},
	}, nil
}
