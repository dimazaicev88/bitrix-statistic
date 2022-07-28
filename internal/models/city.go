package models

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/storage"
)

type CityModel struct {
	storage storage.Storage
}

func NewCityModel(storageImpl storage.Storage) CityModel {
	return CityModel{storage: storageImpl}
}

func (m CityModel) Find(filter filters.Filter) (error, []map[string]interface{}) {
	return nil, nil
}
