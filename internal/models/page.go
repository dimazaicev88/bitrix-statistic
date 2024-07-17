package models

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Page struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewPage(ctx context.Context, chClient driver.Conn) *Page {
	return &Page{
		ctx:      ctx,
		chClient: chClient,
	}
}
