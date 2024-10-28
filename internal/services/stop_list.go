package services

import (
	"bitrix-statistic/internal/models"
	"context"
)

type StopListService struct {
	ctx       context.Context
	allModels *models.Models
}

func NewStopList(ctx context.Context, allModels *models.Models) *StopListService {
	return &StopListService{
		ctx:       ctx,
		allModels: allModels,
	}
}

func (sl StopListService) FindAll(skip uint32, limit uint32) ([]models.StopList, error) {
	return sl.allModels.StopList.FindAll(skip, limit)
}
