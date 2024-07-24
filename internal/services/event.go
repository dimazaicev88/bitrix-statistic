package services

import (
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type EventService struct {
	modelEvent *models.Event
}

func NewEvent(ctx context.Context, chClient driver.Conn) *EventService {
	return &EventService{
		modelEvent: models.NewEvent(ctx, chClient),
	}
}
