package models

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type OptionModel struct {
	chClient driver.Conn
}

func NewOptionModel(ctx context.Context, chClient driver.Conn) *OptionModel {
	return &OptionModel{chClient: chClient}
}
