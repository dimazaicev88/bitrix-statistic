package models

import (
	"bitrix-statistic/internal/filters"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Country struct {
	ctx      context.Context
	chClient driver.Conn
}

func (cm Country) Find(filter filters.Filter) (error, []map[string]interface{}) {
	var hits []map[string]interface{}
	//_, sql := builders.NewCountrySQLBuilder(filter).BuildSQL()
	//rows, err := cm.chClient.DB().Queryx(sql.SQL, sql.Params...)
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

func (cm Country) DeleteById(id int) {
	//cm.storage.DB().MustExec("DELETE FROM country WHERE id=?", id)
}
