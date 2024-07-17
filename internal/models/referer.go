package models

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

// RefererModel Получение данных о ссылающихся сайтах.
type RefererModel struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewRefererModel(ctx context.Context, chClient driver.Conn) *RefererModel {
	return &RefererModel{
		ctx:      ctx,
		chClient: chClient,
	}
}
