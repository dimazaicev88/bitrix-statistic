package models

import (
	"bitrix-statistic/internal/entity"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type OptionModel struct {
	chClient driver.Conn
}

func NewOptionModel(ctx context.Context, chClient driver.Conn) *OptionModel {
	return &OptionModel{chClient: chClient}
}

func (o OptionModel) Add(options []entity.Option) error {
	return nil
}

func (o OptionModel) Set(key, value, desc string) error {
	return nil
}

func (o OptionModel) Get(name string) string {
	return ""
}

func (o OptionModel) GetWithDefault(s string, s2 string) string {
	return ""
}
