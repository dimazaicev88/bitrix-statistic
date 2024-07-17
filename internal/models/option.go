package models

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Option struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewOption(ctx context.Context, chClient driver.Conn) *Option {
	return &Option{ctx: ctx, chClient: chClient}
}
