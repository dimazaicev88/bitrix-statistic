package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type HitModel struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewHit(ctx context.Context, chClient driver.Conn) *HitModel {
	return &HitModel{ctx: ctx, chClient: chClient}
}

func (hm HitModel) Find(filter filters.Filter) ([]entitydb.Hit, error) {
	return nil, nil
}

func (hm HitModel) FindByUuid(uuid string) (entitydb.Hit, error) {
	var hit entitydb.Hit
	err := hm.chClient.QueryRow(hm.ctx, `select * from hit where uuid=?`, uuid).Scan(&hit)
	if err != nil {
		return entitydb.Hit{}, err
	}
	return hit, nil
}

func (hm HitModel) DeleteById(id int) {
	//hm.storage.DB().MustExec("DELETE FROM b_stat_hit WHERE ID=?", id)
}

func (hm HitModel) AddHit(hit entitydb.Hit) error {
	//_, err := hm.storage.DB().MustExec("INSERT INTO hit(`SessionId`, `DATE_HIT`, `GuestUuid`, `IsNewGuest`, `USER_ID`, `USER_AUTH`, `Url`, `Url404`, `URL_FROM`, `Ip`, `METHOD`, `COOKIES`, `UserAgent`, `StopListUuid`, `CountryId`, `CityUuid`, `SiteId`)"+
	//	"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)",
	//	hit.SessionId, hit.DateHit, hit.GuestUuid, hit.IsNewGuest, hit.UserId, hit.IsUserAuth, hit.Url, hit.Url404, hit.UrlFrom, hit.Method, hit.Cookies, hit.UserAgent, hit.StopListUuid, hit.CountryId, hit.CityUuid, hit.SiteId).LastInsertId()
	//if err != nil {
	//	return err
	//}
	return nil
}
