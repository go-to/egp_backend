package repository

import (
	"github.com/go-to/egp_backend/model"
	"time"
)

type IEventRepository interface {
	GetActiveEvents(time *time.Time) (*model.ActiveEvent, error)
}

type EventRepository struct {
	model model.IEventModel
}

func NewEventRepository(m model.EventModel) *EventRepository {
	return &EventRepository{&m}
}

func (r *EventRepository) GetActiveEvents(time *time.Time) (*model.ActiveEvent, error) {
	return r.model.FindActiveEvent(time)
}
