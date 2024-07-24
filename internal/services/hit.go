package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type HitService struct {
	hitModel *models.HitModel
}

func NewHit(ctx context.Context, chClient driver.Conn) *HitService {
	return &HitService{
		hitModel: models.NewHit(ctx, chClient),
	}
}

func (hs *HitService) Find(filter filters.Filter) ([]entitydb.Hit, error) {
	return hs.hitModel.Find(filter)
}

func (hs *HitService) FindByUuid(uuid string) (entitydb.Hit, error) {
	return hs.hitModel.FindByUuid(uuid)
}
