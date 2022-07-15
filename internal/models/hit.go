package models

import (
	"bitrix-statistic/internal/builders"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/storage"
)

type HitModel struct {
	storage *storage.MysqlStorage
}

func NewHitModel(storage *storage.MysqlStorage) HitModel {
	return HitModel{storage: storage}
}

func (hm HitModel) Find(filter filters.Filter) (error, []map[string]interface{}) {
	var hits []map[string]interface{}
	sql := builders.NewHitSQLBuilder(filter).BuildSQL()
	rows, err := hm.storage.DB().Queryx(sql.SQL, sql.Params...)
	for rows.Next() {
		results := make(map[string]interface{})
		err = rows.MapScan(results)
		hits = append(hits, results)
		if err != nil {
			return err, nil
		}
	}
	if err != nil {
		return err, nil
	}
	return nil, hits
}
