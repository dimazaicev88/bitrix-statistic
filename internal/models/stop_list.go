package models

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type StopList struct {
	ctx context.Context
}

func NewStopList(ctx context.Context, client driver.Conn) *StopList {
	return &StopList{ctx: ctx}
}
