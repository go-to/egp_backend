package repository

import (
	"github.com/go-to/egp_backend/model"
	"time"
)

type IStampRepository interface {
	AddStamp(time *time.Time, userId string, shopId int64) (int32, error)
}

type StampRepository struct {
	model model.StampModel
}

func NewStampRepository(m model.StampModel) *StampRepository {
	return &StampRepository{model: m}
}

func (r *StampRepository) AddStamp(time *time.Time, userId string, shopId int64) (int32, error) {
	return r.model.AddStamp(time, userId, shopId)
}
