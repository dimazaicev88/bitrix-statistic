package models

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

// PathModel Поисковые фразы
type PathModel struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewPathModel(ctx context.Context, chClient driver.Conn) *PathModel {
	return &PathModel{
		ctx:      ctx,
		chClient: chClient,
	}
}
