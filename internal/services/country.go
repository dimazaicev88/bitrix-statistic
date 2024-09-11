package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
)

type CountryServices struct {
	allModels *models.Models
	ctx       context.Context
}

func NewCountry(ctx context.Context, allModels *models.Models) *CountryServices {
	return &CountryServices{
		ctx:       ctx,
		allModels: allModels,
	}
}

func (cs *CountryServices) Find(filter filters.Filter) ([]entitydb.Country, error) {
	return cs.allModels.Country.Find(filter)
}
