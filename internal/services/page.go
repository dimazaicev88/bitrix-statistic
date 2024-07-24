package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type PageService struct {
	pageModel *models.Page
}

func NewPage(ctx context.Context, chClient driver.Conn) *PageService {
	return &PageService{pageModel: models.NewPage(ctx, chClient)}
}

func (ps PageService) Filter(filter filters.Filter) ([]entitydb.PageDB, error) {
	return ps.pageModel.Filter(filter)
}

func (ps PageService) DynamicList(filter filters.Filter) ([]entitydb.PageDB, error) {
	return ps.pageModel.DynamicList(filter)
}
