package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
)

type PageService struct {
	allModels *models.Models
	ctx       context.Context
}

func NewPage(ctx context.Context, allModels *models.Models) *PageService {
	return &PageService{
		ctx:       ctx,
		allModels: allModels,
	}
}

func (ps PageService) Filter(filter filters.Filter) ([]entitydb.Page, error) {
	return ps.allModels.Page.Filter(filter)
}

func (ps PageService) DynamicList(filter filters.Filter) ([]entitydb.Page, error) {
	return ps.allModels.Page.DynamicList(filter)
}

func (ps PageService) FindByPageAndDir(dir, page, dateStat string) ([]entitydb.Page, error) {
	return ps.allModels.Page.FindByPageAndDir(dir, page, dateStat)
}

func (ps PageService) FindByUuid(uuid string) (entitydb.Page, error) {
	return ps.allModels.Page.FindByUuid(uuid)
}

func (ps PageService) Update(oldValue entitydb.Page, newValue entitydb.Page) error {
	return ps.allModels.Page.Update(oldValue, newValue)
}
