package models

import (
	"bitrix-statistic/internal/entitydb"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

// Path  Поисковые фразы
type Path struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewPath(ctx context.Context, chClient driver.Conn) *Path {
	return &Path{
		ctx:      ctx,
		chClient: chClient,
	}
}

func (p Path) AddCache(pathCache entitydb.PathCache) error {
	return p.chClient.Exec(p.ctx,
		`INSERT INTO path_cache (uuid, session_uuid, date_hit, path_uuid, path_pages, path_first_page, path_first_page_404, path_first_page_site_id, path_last_page, path_last_page_404, path_page_site_id, path_steps, is_last_page) 
			   VALUES (generateUUIDv7(),?,now(),?,?,?,?,?,?,?,?,?,?)`,
		pathCache.SessionUuid, pathCache.DateHit, pathCache.PathUuid, pathCache.PathPages, pathCache.PathFirstPage, pathCache.PathFirstPage404, pathCache.PathFirstPageSiteId,
		pathCache.PathLastPage, pathCache.PathLastPage404, pathCache.PathLastPageSiteId,
	)
}

func (p Path) FindLastBySessionUuid(uuid string) (entitydb.PathCache, error) {
	var pathCache entitydb.PathCache
	err := p.chClient.QueryRow(p.ctx, `SELECT * FROM path_cache WHERE session_uuid = ?`, uuid).ScanStruct(&pathCache)
	if err != nil {
		return entitydb.PathCache{}, err
	}
	return pathCache, nil
}

func (p Path) FindByReferer(uuid string, referer string) (entitydb.PathCache, error) {
	var pathCache entitydb.PathCache
	err := p.chClient.QueryRow(p.ctx, `SELECT * FROM path_cache WHERE session_uuid = ? and path_last_page=?`, uuid, referer).ScanStruct(&pathCache)
	if err != nil {
		return entitydb.PathCache{}, err
	}
	return pathCache, nil
}

func (p Path) FindBySession(uuid string) (entitydb.PathCache, error) {
	var pathCache entitydb.PathCache
	err := p.chClient.QueryRow(p.ctx, `SELECT * FROM path_cache WHERE session_uuid = ? and length(path_last_page)<0`, uuid).ScanStruct(&pathCache)
	if err != nil {
		return entitydb.PathCache{}, err
	}
	return pathCache, nil
}
