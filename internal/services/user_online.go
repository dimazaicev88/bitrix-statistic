package services

import (
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type UserOnlineService struct {
	ctx             context.Context
	chClient        driver.Conn
	userOnlineModel *models.UserOnline
}

func NewUserOnline(ctx context.Context, chClient driver.Conn) *UserOnlineService {
	return &UserOnlineService{
		ctx:             ctx,
		chClient:        chClient,
		userOnlineModel: models.NewUserOnline(ctx, chClient),
	}
}
