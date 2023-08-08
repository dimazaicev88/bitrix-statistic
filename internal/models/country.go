package models

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/storage"
)

type CountryModel struct {
	storage storage.Storage
}

func (cm CountryModel) Find(filter filters.Filter) (error, []map[string]interface{}) {
	var hits []map[string]interface{}
	//_, sql := builders.NewCountrySQLBuilder(filter).BuildSQL()
	//rows, err := cm.storage.DB().Queryx(sql.SQL, sql.Params...)
	//for rows.Next() {
	//	results := make(map[string]interface{})
	//	err = rows.MapScan(results)
	//	hits = append(hits, results)
	//	if err != nil {
	//		return err, nil
	//	}
	//}
	//if err != nil {
	//	return err, nil
	//}
	return nil, hits
}

func (cm CountryModel) DeleteById(id int) {
	cm.storage.DB().MustExec("DELETE FROM b_stat_country WHERE ID=?", id)
}
