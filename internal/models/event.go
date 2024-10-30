package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Event struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewEvent(ctx context.Context, chClient driver.Conn) *Event {
	return &Event{
		ctx:      ctx,
		chClient: chClient,
	}
}

func (e Event) Delete(uuid string) error {
	//TODO
	return nil
}

func (e Event) FindAll(skip uint32, limit uint32) ([]entitydb.Event, error) {
	return nil, nil
}

func (e Event) Find(filter filters.Filter) ([]entitydb.Event, error) {
	return nil, nil
}
