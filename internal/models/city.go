package models

import (
	"bitrix-statistic/internal/filters"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type CityModel struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewCityModel(ctx context.Context, chClient driver.Conn) *CityModel {
	return &CityModel{ctx: ctx, chClient: chClient}
}

func (cm CityModel) Find(filter filters.Filter) (error, []map[string]interface{}) {
	return nil, nil
}

func (cm CityModel) DeleteById(id int) {
	//cm.storage.DB().MustExec("DELETE FROM city WHERE id=?", id)
}

func (cm CityModel) GetCountryCode() string {
	//TODO implement
	return ""
}

func (cm CityModel) GetCityID() string {
	//TODO implement
	return ""
}
