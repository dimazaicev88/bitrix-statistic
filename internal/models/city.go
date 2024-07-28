package models

import (
	"bitrix-statistic/internal/filters"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type City struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewCity(ctx context.Context, chClient driver.Conn) *City {
	return &City{ctx: ctx, chClient: chClient}
}

func (cm City) Find(filter filters.Filter) (error, []map[string]interface{}) {
	return nil, nil
}

func (cm City) DeleteById(id int) {
	//cm.storage.().MustExec("DELETE FROM city WHERE id=?", id)
}

func (cm City) GetCountryCode() string {
	//TODO implement
	return ""
}

func (cm City) GetCityID() string {
	//TODO implement
	return ""
}
