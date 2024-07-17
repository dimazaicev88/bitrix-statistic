package models

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type StatDayModel struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewStatDayModel(ctx context.Context, chClient driver.Conn) *StatDayModel {
	return &StatDayModel{ctx: ctx, chClient: chClient}
}

func (cm *StatDayModel) AddDay(day string) error {
	return nil
}
