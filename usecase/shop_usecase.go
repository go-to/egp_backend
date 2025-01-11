package usecase

import (
	"github.com/go-to/egp_backend/repository"
	"github.com/go-to/egp_backend/usecase/input"
	"github.com/go-to/egp_backend/usecase/output"
	"github.com/go-to/egp_backend/util"
	"github.com/go-to/egp_protobuf/pb"
)

type IShopUsecase interface {
	GetShops(input *input.ShopsInput) (*output.ShopsOutput, error)
}

type ShopUsecase struct {
	repo repository.ShopRepository
}

func NewShopUseCase(repo repository.ShopRepository) *ShopUsecase {
	return &ShopUsecase{repo: repo}
}

func (s *ShopUsecase) GetShops(in *input.ShopsInput) (*output.ShopsOutput, error) {

	now := util.Now()
	// FIXME デバッグ用なので最終的に消す
	//now = time.Date(2025, 4, 1, 23, 0, 0, 0, util.Location)
	//fmt.Println(now)

	shops, err := s.repo.GetShops(&now)
	if err != nil {
		return &output.ShopsOutput{}, nil
	}

	var outputShops []*pb.Shop

	for _, v := range *shops {
		inCurrentSales := true
		if len(v.StartTime) == 0 || len(v.EndTime) == 0 {
			inCurrentSales = false
		}

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
			Latitude:                   v.Latitude,
			Longitude:                  v.Longitude,
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
