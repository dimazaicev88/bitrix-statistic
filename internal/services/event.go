package services

import (
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

func (e EventService) Delete(eventUuid string) error {
	return e.allModels.Event.Delete(eventUuid)
}
