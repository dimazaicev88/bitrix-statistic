package repository

import (
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Hit struct {
	chClient driver.Conn
}

func NewHit(chClient driver.Conn) *Hit {
	return &Hit{chClient: chClient}
}

func (hm Hit) AddHit(ctx context.Context, hit models.Hit) error {
	return hm.chClient.AsyncInsert(
		ctx,
		`INSERT INTO hit (uuid, dateHit, cookies, event1, event2, guestHash, ip, method, phpSessionId, 
                 referrer, siteId, url, urlFrom, userAgent, userId)
		       VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?)`,
		false,
		hit.Uuid, hit.DateHit, hit.Cookies, hit.Event1, hit.Event2, hit.GuestHash, hit.Ip, hit.Method, hit.PhpSessionId,
		hit.Referer, hit.SiteId, hit.Url, hit.UrlFrom, hit.UserAgent, hit.UserId,
	)
}
