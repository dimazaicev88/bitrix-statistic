package models

import (
	"bitrix-statistic/internal/builders"
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/storage"
)

type HitModel struct {
	storage storage.Storage
}

func NewHitModel(storageImpl storage.Storage) HitModel {
	return HitModel{storage: storageImpl}
}

func (hm HitModel) Find(filter filters.Filter) (error, []map[string]interface{}) {
	var hits []map[string]interface{}
	sql, err := builders.NewHitSQLBuilder(filter).BuildSQL()
	if err != nil {
		return err, nil
	}
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

func (hm HitModel) Find2(filter filters.BitrixHitFilter) (error, []map[string]interface{}) {
	var hits []map[string]interface{}

	return nil, hits
}

func (hm HitModel) DeleteById(id int) {
	hm.storage.DB().MustExec("DELETE FROM b_stat_hit WHERE ID=?", id)
}

func (hm HitModel) AddHit(hit entity.Hit) error {
	_, err := hm.storage.DB().MustExec("INSERT INTO b_stat_hit(`SessionId`, `DATE_HIT`, `GuestId`, `NewGuest`, `USER_ID`, `USER_AUTH`, `Url`, `Url404`, `URL_FROM`, `Ip`, `METHOD`, `COOKIES`, `UserAgent`, `StopListId`, `CountryId`, `CityId`, `SiteId`)"+
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)",
		hit.SessionId, hit.DateHit, hit.GuestId, hit.NewGuest, hit.UserId, hit.UserAuth, hit.Url, hit.Url404, hit.UrlFrom, hit.Method, hit.Cookies, hit.UserAgent, hit.StopListId, hit.CountryId, hit.CityId, hit.SiteId).LastInsertId()
	if err != nil {
		return err
	}
	return nil
}
