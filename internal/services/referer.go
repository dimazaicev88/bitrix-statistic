package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
)

type RefererService struct {
	allModels *models.Models
	ctx       context.Context
}

func NewReferer(ctx context.Context, allModels *models.Models) *RefererService {
	return &RefererService{
		ctx:       ctx,
		allModels: allModels,
	}
}

func (rs RefererService) Find(filter filters.Filter) ([]entitydb.Referer, error) {
	return rs.allModels.Referer.Find(filter)
}
