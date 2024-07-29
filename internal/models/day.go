package models

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Day struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewDay(ctx context.Context, chClient driver.Conn) *Day {
	return &Day{
		ctx:      ctx,
		chClient: chClient,
	}
}

func (d Day) Add() {

}
