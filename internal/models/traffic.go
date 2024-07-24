package models

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Traffic struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewTraffic(ctx context.Context, chClient driver.Conn) *Traffic {
	return &Traffic{
		ctx:      ctx,
		chClient: chClient,
	}
}
