package models

import (
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
	return nil, nil
}

func (hm Hit) FindByUuid(uuid uuid.UUID) (entitydb.Hit, error) {
	var hit entitydb.Hit
	err := hm.chClient.QueryRow(hm.ctx, `select * from hit where uuid=?`, uuid).Scan(&hit)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return hit, err
	}
	return hit, nil
}

func (hm Hit) AddHit(hit entitydb.Hit) error {
	return hm.chClient.Exec(hm.ctx,
		`INSERT INTO hit (uuid, session_uuid, adv_uuid, date_hit, php_session_id, guest_uuid, new_guest, user_id, user_auth, url, url_404, url_from,
	            ip, method, cookies, user_agent, stop_list_uuid, country_id, city_uuid, site_id)
		       VALUES (?,  ?, ?, now(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?,?,?,?)`,
		hit.Uuid, hit.SessionUuid, hit.AdvUuid, hit.PhpSessionId, hit.GuestUuid, hit.IsNewGuest, hit.UserId, hit.IsUserAuth, hit.Url, hit.Url404, hit.UrlFrom, hit.Ip,
		hit.Method, hit.Cookies, hit.UserAgent, hit.StopListUuid, hit.CountryId, hit.CityUuid, hit.SiteId)
}

func (hm Hit) FindLastHitWithoutSession(guestUuid uuid.UUID, sessionId string) (entitydb.Hit, error) {
	var hit entitydb.Hit
	err := hm.chClient.QueryRow(hm.ctx, `select * from hit where guest_uuid=? and php_session_id!=? order by date_hit desc limit 1`, guestUuid, sessionId).ScanStruct(&hit)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return hit, err
	}
	return hit, nil
}
