package repository

import (
	"bitrix-statistic/internal/models"
	"context"
	"github.com/uptrace/go-clickhouse/ch"
)

type Hit struct {
	chClient *ch.DB
}

func NewHit(chClient *ch.DB) *Hit {
	return &Hit{chClient: chClient}
}

func (hm Hit) AddHit(ctx context.Context, hit models.Hit) error {
	return hm.chClient.AsyncInsert(
		ctx,
		`INSERT INTO hit (uuid, dateHit, cookies, countryId, city, event1, event2, guestHash, ip, method, phpSessionId, 
                 referrer, siteId, url, urlFrom, userAgent, userId)
		       VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?)`,
		false,
		hit.Uuid, hit.DateHit, hit.Cookies, hit.CountryId, hit.CityId, hit.Event1, hit.Event2, hit.GuestHash, hit.Ip, hit.PhpSessionId,
		hit.Referer, hit.SiteId, hit.Url, hit.UrlFrom, hit.UserAgent, hit.UserId,
	)
}
