package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type RefererService struct {
	refererModel *models.RefererModel
}

func NewReferer(ctx context.Context, chClient driver.Conn) *RefererService {
	return &RefererService{
		refererModel: models.NewReferer(ctx, chClient),
	}
}

func (rs RefererService) Find(filter filters.Filter) ([]entitydb.RefererDB, error) {
	return rs.refererModel.Find(filter)
}
