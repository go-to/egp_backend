package usecase

import (
	"github.com/go-to/egp_backend/repository"
	"github.com/go-to/egp_backend/usecase/input"
	"github.com/go-to/egp_backend/usecase/output"
	"github.com/go-to/egp_protobuf/pb"
)

type IStampUsecase interface {
	AddStamp(in *input.AddStampInput) (*output.AddStampOutput, error)
}

type StampUsecase struct {
	config repository.IConfigRepository
	stamp  repository.IStampRepository
}

func NewStampUseCase(config repository.ConfigRepository, stamp repository.StampRepository) *StampUsecase {
	return &StampUsecase{
		config: &config,
		stamp:  &stamp,
	}
}

func (u *StampUsecase) AddStamp(in *input.AddStampInput) (*output.AddStampOutput, error) {
	userId := in.AddStampRequest.GetUserId()
	shopId := in.AddStampRequest.GetShopId()

	now, err := u.config.GetTime()
	if err != nil {
		return &output.AddStampOutput{}, err
	}

	stampNum, err := u.stamp.AddStamp(&now, userId, shopId)
	if err != nil {
		return &output.AddStampOutput{}, err
	}

	return &output.AddStampOutput{
		AddStampResponse: pb.AddStampResponse{
			NumberOfTimes: stampNum,
		},
	}, nil
}
