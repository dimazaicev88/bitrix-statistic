package models

import (
	"bitrix-statistic/internal/builders"
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/storage"
)

type HitModel struct {
	storage *storage.MysqlStorage
}

func NewHitModel(storage *storage.MysqlStorage) HitModel {
	return HitModel{storage: storage}
}

func (hm HitModel) Find(filter filters.Filter) (error, []entity.Hit) {
	var hits []entity.Hit
	sql := builders.NewHitSQLBuilder(filter).BuildSQL()
	err := hm.storage.DB().Select(&hits, sql.SQL, sql.Params...)
	if err != nil {
		return err, nil
	}
	return nil, hits
}
