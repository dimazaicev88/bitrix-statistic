package models

import (
	"bitrix-statistic/internal/builders"
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"context"
	"database/sql"
	"errors"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

type Hit struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewHit(ctx context.Context, chClient driver.Conn) *Hit {
	return &Hit{ctx: ctx, chClient: chClient}
}

func (hm Hit) Find(filter filters.Filter) ([]entitydb.Hit, error) {
	builder := builders.NewHitSQLBuilder(filter)
	resultSql, args, err := builder.Build()
	if err != nil {
		return nil, err
	}

	rows, err := hm.chClient.Query(hm.ctx, resultSql, args...)

	if err != nil {
		return nil, err
	}

	var allDbHits = make([]entitydb.Hit, 0)
	for rows.Next() {
		var hit entitydb.Hit
		err = rows.ScanStruct(&hit)
		allDbHits = append(allDbHits, hit)
	}

	return allDbHits, nil
}

func (hm Hit) FindByUuid(uuid uuid.UUID) (entitydb.Hit, error) {
	var hit entitydb.Hit
	err := hm.chClient.QueryRow(hm.ctx, `select * from hit where uuid=?`, uuid).ScanStruct(&hit)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return hit, err
	}
	return hit, nil
}

func (hm Hit) AddHit(hit entitydb.Hit) error {
	return hm.chClient.Exec(hm.ctx,
		`INSERT INTO hit (uuid, sessionUuid, advUuid, dateHit, phpSessionId, guestUuid, isNewGuest, userId, userAuth, url, url404, urlFrom,
	            ip, method, cookies, userAgent, stopListUuid, countryId, cityUuid, siteId)
		       VALUES (?,  ?, ?, now(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?,?,?,?)`,
		hit.Uuid, hit.SessionUuid, hit.AdvUuid, hit.PhpSessionId, hit.GuestUuid, hit.IsNewGuest, hit.UserId, hit.IsUserAuth, hit.Url, hit.Url404, hit.UrlFrom, hit.Ip,
		hit.Method, hit.Cookies, hit.UserAgent, hit.StopListUuid, hit.CountryId, hit.CityUuid, hit.SiteId)
}

func (hm Hit) FindLastHitWithoutSession(guestUuid uuid.UUID, sessionId string) (entitydb.Hit, error) {
	var hit entitydb.Hit
	err := hm.chClient.QueryRow(hm.ctx, `select * from hit where guestUuid=? and phpSessionId!=? order by dateHit desc limit 1`, guestUuid, sessionId).ScanStruct(&hit)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return hit, err
	}
	return hit, nil
}
