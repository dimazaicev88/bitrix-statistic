package models

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Phrase struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewPhrase(ctx context.Context, chClient driver.Conn) *Phrase {
	return &Phrase{
		ctx:      ctx,
		chClient: chClient,
	}
}
