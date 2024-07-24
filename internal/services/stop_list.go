package services

import (
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type StopListService struct {
	ctx           context.Context
	chClient      driver.Conn
	stopListModel *models.StopList
}

func NewStopList(ctx context.Context, chClient driver.Conn) *StopListService {
	return &StopListService{
		ctx:           ctx,
		chClient:      chClient,
		stopListModel: models.NewStopList(ctx),
	}
}
