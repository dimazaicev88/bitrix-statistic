package models

import (
	"bitrix-statistic/internal/entitydb"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type UserOnline struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewUserOnline(ctx context.Context, chClient driver.Conn) *UserOnline {
	return &UserOnline{
		ctx:      ctx,
		chClient: chClient,
	}
}

func (o UserOnline) FindAll(skip uint32, limit uint32) ([]entitydb.UserOnline, error) {
	return nil, nil
}
