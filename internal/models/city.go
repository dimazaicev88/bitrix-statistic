package models

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/storage"
)

type CityModel struct {
	storage *storage.MysqlStorage
}

func NewCityModel(mysqlStorage *storage.MysqlStorage) CityModel {
	return CityModel{storage: mysqlStorage}
}

func (m CityModel) Find(filter filters.Filter) (error, []map[string]interface{}) {
	return nil, nil
}
