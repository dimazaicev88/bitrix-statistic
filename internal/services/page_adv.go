package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"context"
)

type PageAdvService struct {
	allModels *models.Models
	ctx       context.Context
}

func NewPageAdvService(ctx context.Context, models *models.Models) *PageAdvService {
	return &PageAdvService{
		allModels: models,
		ctx:       ctx,
	}
}

func (pas *PageAdvService) Add(pageAdv entitydb.PageAdv) error {
	return pas.allModels.PageAdv.Add(pageAdv)
}
