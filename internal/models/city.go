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

func (cm CityModel) Find(filter filters.Filter) (error, []map[string]interface{}) {
	return nil, nil
}

func (cm CityModel) DeleteById(id int) {
	cm.storage.DB().MustExec("DELETE FROM city WHERE id=?", id)
}

func (cm CityModel) GetCountryCode() string {
	//TODO implement
	return ""
}

func (cm CityModel) GetCityID() string {
	//TODO implement
	return ""
}
