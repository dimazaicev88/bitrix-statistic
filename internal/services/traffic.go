package services

import (
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type TrafficService struct {
	ctx          context.Context
	chClient     driver.Conn
	trafficModel *models.Traffic
}

func NewTraffic(ctx context.Context, chClient driver.Conn) *TrafficService {
	return &TrafficService{
		ctx:          ctx,
		chClient:     chClient,
		trafficModel: models.NewTraffic(ctx, chClient),
	}
}
