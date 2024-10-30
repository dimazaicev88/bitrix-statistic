package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
)

type EventService struct {
	allModels *models.Models
	ctx       context.Context
}

func NewEvent(ctx context.Context, allModels *models.Models) *EventService {
	return &EventService{
		ctx:       ctx,
		allModels: allModels,
	}
}

func (e EventService) Find(filter filters.Filter) ([]entitydb.Event, error) {
	return nil, nil
}

func (e EventService) FindAll(skip uint32, limit uint32) ([]entitydb.Event, error) {
	return nil, nil
}

func (e EventService) ConvertToJSONListEvents(events []entitydb.Event) any {
	return nil
}
