package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

// Referer Получение данных о ссылающихся сайтах.
type Referer struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewReferer(ctx context.Context, chClient driver.Conn) *Referer {
	return &Referer{
		ctx:      ctx,
		chClient: chClient,
	}
}

func (m Referer) Find(filter filters.Filter) ([]entitydb.Referer, error) {
	return nil, nil
}

func (m Referer) Add(referer string) (uuid.UUID, error) {
	refererUuid := uuid.New()
	return refererUuid, m.chClient.Exec(m.ctx, `INSERT INTO referer (uuid, site_name, sessions, hits) VALUES (?, ?, 1, 1);`, refererUuid, referer)
}

func (m Referer) AddToRefererList(refererList entitydb.RefererList) error {
	return m.chClient.Exec(m.ctx,
		`INSERT INTO referer_list (uuid, referer_uuid, date_hit, protocol, site_name, url_from, url_to, session_uuid, adv_uuid, site_id)
               VALUES (generateUUIDv7(),?,now(),?,?,?,?,?,?,?)`,
		refererList.RefererUuid, refererList.Protocol, refererList.SiteName, refererList.UrlTo, refererList.SessionUuid, refererList.SiteId)
}
