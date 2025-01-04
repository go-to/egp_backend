package usecase

import (
	"github.com/go-to/egp-backend/repository"
	"github.com/go-to/egp-backend/usecase/input"
	"github.com/go-to/egp-backend/usecase/output"
	"github.com/go-to/egp-protobuf/pb"
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
	shops, err := s.repo.GetShops()
	if err != nil {
		return &output.ShopsOutput{}, nil
	}

	var outputShops []*pb.Shop

	for _, v := range *shops {
		outputShops = append(outputShops, &pb.Shop{
			ID:       v.ID,
			ShopName: v.ShopName,
		})
	}

	return &output.ShopsOutput{
		ShopsResponse: pb.ShopsResponse{
			Shops: outputShops,
		},
	}, nil
}
