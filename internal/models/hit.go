package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"context"
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

func (hm Hit) FindByUuid(uuid string) (entitydb.Hit, error) {
	var hit entitydb.Hit
	err := hm.chClient.QueryRow(hm.ctx, `select * from hit where uuid=?`, uuid).Scan(&hit)
	if err != nil {
		return entitydb.Hit{}, err
	}
	return hit, nil
}

func (hm Hit) AddHit(hit entitydb.Hit) (string, error) {
	hitUuid := uuid.New().String()
	err := hm.chClient.Exec(hm.ctx, `INSERT INTO hit(uuid, session_uuid, date_hit, guest_uuid, user_id, user_auth, url, url_404, url_from,
                ip, method, cookies, user_agent, stop_list_uuid,country_id, city_uuid, site_id)
		VALUES (?,  ?,now(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)`,
		hitUuid, hit.SessionUuid, hit.GuestUuid, hit.UserId, hit.IsUserAuth, hit.Url, hit.Url404, hit.UrlFrom,
		hit.Method, hit.Cookies, hit.UserAgent, hit.StopListUuid, hit.CountryId, hit.CityUuid, hit.SiteId)
	if err != nil {
		return "", err
	}
	return hitUuid, nil
}

func (hm Hit) FindLastHitWithoutSession(guestUuid string, withoutPhpSessionId string) (entitydb.Hit, error) {
	var hit entitydb.Hit
	err := hm.chClient.QueryRow(hm.ctx, `select * from guest`).ScanStruct(&hit)
	if err != nil {
		return entitydb.Hit{}, err
	}
	return hit, nil
}
