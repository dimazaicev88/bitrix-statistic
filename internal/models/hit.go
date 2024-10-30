package models

import (
	"bitrix-statistic/internal/entitydb"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Hit struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewHit(ctx context.Context, chClient driver.Conn) *Hit {
	return &Hit{ctx: ctx, chClient: chClient}
}

func (hm Hit) AddHit(hit entitydb.Hit) error {
	return hm.chClient.AsyncInsert(hm.ctx,
		`INSERT INTO hit (uuid, sessionUuid, advUuid, dateHit, phpSessionId, guestUuid, 
                 language, isNewGuest, userId, userAuth, url, url404, urlFrom, ip, method, cookies, userAgent, stopListUuid, countryId, cityUuid, siteId, favorites)
		       VALUES (?,  ?, ?, now(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?,?,?,?)`,
		hit.Uuid, hit.SessionUuid, hit.AdvUuid, hit.PhpSessionId, hit.GuestUuid, hit.IsNewGuest, hit.UserId, hit.IsUserAuth, hit.Url, hit.Url404, hit.UrlFrom, hit.Ip,
		hit.Method, hit.Cookies, hit.UserAgent, hit.StopListUuid, hit.CountryId, hit.CityUuid, hit.SiteId)
}
