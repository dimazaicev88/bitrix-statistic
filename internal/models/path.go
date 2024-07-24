package models

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

// Path  Поисковые фразы
type Path struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewPath(ctx context.Context, chClient driver.Conn) *Path {
	return &Path{
		ctx:      ctx,
		chClient: chClient,
	}
}
