package models

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Event struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewEvent(ctx context.Context, chClient driver.Conn) *Event {
	return &Event{
		ctx:      ctx,
		chClient: chClient,
	}
}
