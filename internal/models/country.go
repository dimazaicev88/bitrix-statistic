package models

import (
	"bitrix-statistic/internal/builders"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/storage"
)

type CountryModel struct {
	storage storage.MysqlStorage
}

func (cm CountryModel) Find(filter filters.Filter) (error, []map[string]interface{}) {
	var hits []map[string]interface{}
	_, sql := builders.NewCountrySQLBuilder(filter).BuildSQL()
	rows, err := cm.storage.DB().Queryx(sql.SQL, sql.Params...)
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
