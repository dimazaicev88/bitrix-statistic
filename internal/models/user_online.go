package models

import (
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
