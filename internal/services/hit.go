package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
)

type HitService struct {
	allModels *models.Models
	ctx       context.Context
}

func NewHit(ctx context.Context, allModels *models.Models) *HitService {
	return &HitService{
		ctx:       ctx,
		allModels: allModels,
	}
}

func (hs *HitService) Find(filter filters.Filter) ([]entitydb.Hit, error) {
	return hs.allModels.Hit.Find(filter)
}

func (hs *HitService) FindByUuid(uuid string) (entitydb.Hit, error) {
	return hs.allModels.Hit.FindByUuid(uuid)
}
