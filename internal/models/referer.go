package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

// Referer Получение данных о ссылающихся сайтах.
type Referer struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewReferer(ctx context.Context, chClient driver.Conn) *Referer {
	return &Referer{
		ctx:      ctx,
		chClient: chClient,
	}
}

func (m Referer) Find(filter filters.Filter) ([]entitydb.Referer, error) {
	return nil, nil
}
