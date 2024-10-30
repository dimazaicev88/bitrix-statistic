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
	return e.allModels.Event.Find(filter)
}

func (e EventService) FindAll(skip, limit uint32) ([]entitydb.Event, error) {
	return e.allModels.Event.FindAll(skip, limit)
}

func (e EventService) ConvertToJSONListEvents(events []entitydb.Event) any {
	return nil
}

func (e EventService) Delete(eventUuid string) error {
	return e.allModels.Event.Delete(eventUuid)
}
