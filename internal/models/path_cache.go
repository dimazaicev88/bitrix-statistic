package models

import (
	"bitrix-statistic/internal/entitydb"
	"context"
	"database/sql"
	"errors"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type PathCache struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewPathCache(ctx context.Context, chClient driver.Conn) *PathCache {
	return &PathCache{
		ctx:      ctx,
		chClient: chClient,
	}
}

func (pc PathCache) Add(pathCache entitydb.PathCache) error {
	return pc.chClient.Exec(pc.ctx,
		`INSERT INTO path_cache (uuid, session_uuid, date_hit, path_id, path_pages, path_first_page, path_first_page_404, path_first_page_site_id, path_last_page, 
                        path_last_page_404, path_page_site_id, path_steps, is_last_page,sign,version) 
			   VALUES (generateUUIDv7(),?,now(),?,?,?,?,?,?,?,?,?,?,?,?)`,
		pathCache.SessionUuid, pathCache.DateHit, pathCache.PathId, pathCache.PathPages, pathCache.PathFirstPage, pathCache.PathFirstPage404, pathCache.PathFirstPageSiteId,
		pathCache.PathLastPage, pathCache.PathLastPage404, pathCache.PathLastPageSiteId, pathCache.PathSteps, pathCache.IsLastPage, pathCache.Sign, pathCache.Version,
	)
}

func (pc PathCache) FindLastBySessionUuid(uuid string) (entitydb.PathCache, error) {
	var pathCache entitydb.PathCache
	err := pc.chClient.QueryRow(pc.ctx, `SELECT * FROM path_cache WHERE session_uuid = ?`, uuid).ScanStruct(&pathCache)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return entitydb.PathCache{}, err
	}
	return pathCache, nil
}

func (pc PathCache) FindByReferer(uuid string, referer string) (entitydb.PathCache, error) {
	var pathCache entitydb.PathCache
	err := pc.chClient.QueryRow(pc.ctx, `SELECT * FROM path_cache WHERE session_uuid = ? and path_last_page=?`, uuid, referer).ScanStruct(&pathCache)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return entitydb.PathCache{}, err
	}
	return pathCache, nil
}

func (pc PathCache) FindBySession(uuid string) (entitydb.PathCache, error) {
	var pathCache entitydb.PathCache
	err := pc.chClient.QueryRow(pc.ctx, `SELECT * FROM path_cache WHERE session_uuid = ? and length(path_last_page)<0`, uuid).ScanStruct(&pathCache)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return entitydb.PathCache{}, err
	}
	return pathCache, nil
}

func (pc PathCache) Update(oldValue entitydb.PathCache, newValue entitydb.PathCache) error {
	err := pc.chClient.Exec(pc.ctx,
		`INSERT INTO path_cache (uuid, session_uuid, date_hit, path_id, path_pages, path_first_page, path_first_page_404, path_first_page_site_id, path_last_page, path_last_page_404, path_page_site_id, path_steps, is_last_page,sign,version) 
			   VALUES (generateUUIDv7(),?,now(),?,?,?,?,?,?,?,?,?,?,?,?)`,
		oldValue.SessionUuid, oldValue.DateHit, oldValue.PathId, oldValue.PathPages, oldValue.PathFirstPage, oldValue.PathFirstPage404, oldValue.PathFirstPageSiteId,
		oldValue.PathLastPage, oldValue.PathLastPage404, oldValue.PathLastPageSiteId, oldValue.Sign*-1, oldValue.Version,
	)
	if err != nil {
		return err
	}

	err = pc.chClient.Exec(pc.ctx,
		`INSERT INTO path_cache (uuid, session_uuid, date_hit, path_id, path_pages, path_first_page, path_first_page_404, path_first_page_site_id, path_last_page, path_last_page_404, path_page_site_id, path_steps, is_last_page,sign,version) 
			   VALUES (generateUUIDv7(),?,now(),?,?,?,?,?,?,?,?,?,?,?,?)`,
		oldValue.SessionUuid, oldValue.DateHit, oldValue.PathId, oldValue.PathPages, oldValue.PathFirstPage, oldValue.PathFirstPage404, oldValue.PathFirstPageSiteId,
		oldValue.PathLastPage, oldValue.PathLastPage404, oldValue.PathLastPageSiteId, 1, oldValue.Version+1,
	)
	if err != nil {
		return err
	}

	return nil
}
