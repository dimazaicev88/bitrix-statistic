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

func (rs RefererService) Add(referer string) (string, error) {
	if referer == "" {
		return "", nil
	}
	return rs.allModels.Referer.Add(referer)
}

func (rs RefererService) AddToRefererList(refererList entitydb.RefererList) error {
	return rs.allModels.Referer.AddToRefererList(refererList)
}
