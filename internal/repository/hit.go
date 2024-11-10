package repository

import (
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type IHitRepository interface {
	AddHit(ctx context.Context, hit models.Hit, wait bool) error
}

type Hit struct {
	chClient driver.Conn
}

func NewHit(chClient driver.Conn) *Hit {
	return &Hit{chClient: chClient}
}

func (hm Hit) AddHit(ctx context.Context, hit models.Hit, wait bool) error {
	return hm.chClient.AsyncInsert(
		ctx,
		`INSERT INTO hits (uuid, dateHit, cookies, event1, event2, guestHash,isNewGuest, ip, method, phpSessionId, 
                 referrer, siteId, url,url404, userAgent, userId)
		       VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		wait,
		hit.Uuid, hit.DateHit, hit.Cookies, hit.Event1, hit.Event2, hit.GuestHash, hit.IsNewGuest, hit.Ip, hit.Method, hit.PhpSessionId,
		hit.Referer, hit.SiteId, hit.Url, hit.Url404, hit.UserAgent, hit.UserId,
	)
}
