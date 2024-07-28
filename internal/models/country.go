package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Country struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewCountry(ctx context.Context, chClient driver.Conn) *Country {
	return &Country{
		ctx:      ctx,
		chClient: chClient,
	}
}

func (cm Country) Find(filter filters.Filter) ([]entitydb.Country, error) {
	return nil, nil
}

func (cm Country) DeleteById(id int) {
	//cm.storage.().MustExec("DELETE FROM country WHERE id=?", id)
}
