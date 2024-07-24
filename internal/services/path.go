package services

import (
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type PathService struct {
	pathModel *models.Path
}

func NewPath(ctx context.Context, chClient driver.Conn) *PathService {
	return &PathService{
		pathModel: models.NewPath(ctx, chClient),
	}
}
