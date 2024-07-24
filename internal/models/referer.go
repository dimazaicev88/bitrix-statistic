package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

// RefererModel Получение данных о ссылающихся сайтах.
type RefererModel struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewReferer(ctx context.Context, chClient driver.Conn) *RefererModel {
	return &RefererModel{
		ctx:      ctx,
		chClient: chClient,
	}
}

func (m RefererModel) Find(filter filters.Filter) ([]entitydb.RefererDB, error) {
	return nil, nil
}
