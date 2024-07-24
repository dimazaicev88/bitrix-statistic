package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
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

func (p *Page) Filter(filter filters.Filter) ([]entitydb.PageDB, error) {
	return nil, nil
}

func (p *Page) DynamicList(filter filters.Filter) ([]entitydb.PageDB, error) {
	return nil, nil
}
