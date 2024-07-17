package models

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Browser struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewBrowser(ctx context.Context, chConn driver.Conn) *Browser {
	return &Browser{
		chClient: chConn,
		ctx:      ctx,
	}
}
