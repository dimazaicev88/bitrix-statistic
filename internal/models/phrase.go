package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
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

func (p Phrase) Filter(filter filters.Filter) ([]entitydb.PhraseList, error) {
	return nil, nil
}
