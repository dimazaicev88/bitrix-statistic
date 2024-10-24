package models

import (
	"bitrix-statistic/internal/entitydb"
	"context"
	"database/sql"
	"errors"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
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
		`INSERT INTO path_cache (uuid, sessionUuid, dateHit, pathId, pathPages, pathFirstPage, pathFirstPage404,
                        pathFirstPageSiteId, pathLastPage, 
                        pathLastPage404, pathPageSiteId, pathSteps, isLastPage,sign,version) 
			   VALUES (generateUUIDv7(),?,now(),?,?,?,?,?,?,?,?,?,?,?,?)`,
		pathCache.SessionUuid, pathCache.DateHit, pathCache.PathId, pathCache.PathPages, pathCache.PathFirstPage, pathCache.PathFirstPage404, pathCache.PathFirstPageSiteId,
		pathCache.PathLastPage, pathCache.PathLastPage404, pathCache.PathLastPageSiteId, pathCache.PathSteps, pathCache.IsLastPage, pathCache.Sign, pathCache.Version,
	)
}

func (pc PathCache) FindLastBySessionUuid(uuid uuid.UUID) (entitydb.PathCache, error) {
	var pathCache entitydb.PathCache
	err := pc.chClient.QueryRow(pc.ctx, `SELECT * FROM path_cache WHERE sessionUuid = ?`, uuid).ScanStruct(&pathCache)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return entitydb.PathCache{}, err
	}
	return pathCache, nil
}

func (pc PathCache) FindByReferer(uuid uuid.UUID, referer string) (entitydb.PathCache, error) {
	var pathCache entitydb.PathCache
	err := pc.chClient.QueryRow(pc.ctx, `SELECT * FROM path_cache WHERE sessionUuid = ? and pathLastPage=?`,
		uuid, referer).ScanStruct(&pathCache)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return entitydb.PathCache{}, err
	}
	return pathCache, nil
}

func (pc PathCache) FindBySession(uuid uuid.UUID) (entitydb.PathCache, error) {
	var pathCache entitydb.PathCache
	err := pc.chClient.QueryRow(pc.ctx, `SELECT * FROM path_cache WHERE sessionUuid = ? and length(pathLastPage)<0`,
		uuid).ScanStruct(&pathCache)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return entitydb.PathCache{}, err
	}
	return pathCache, nil
}

func (pc PathCache) Update(oldValue entitydb.PathCache, newValue entitydb.PathCache) error {
	err := pc.chClient.Exec(pc.ctx,
		`INSERT INTO path_cache (uuid, sessionUuid, dateHit, pathId, pathPages, pathFirstPage, pathFirstPage404,
                        pathFirstPageSiteId, pathLastPage, pathLastPage404, pathPageSiteId, pathSteps, 
                        isLastPage,sign,version) 
			   VALUES (generateUUIDv7(),?,now(),?,?,?,?,?,?,?,?,?,?,?,?)`,
		oldValue.SessionUuid, oldValue.DateHit, oldValue.PathId, oldValue.PathPages, oldValue.PathFirstPage, oldValue.PathFirstPage404, oldValue.PathFirstPageSiteId,
		oldValue.PathLastPage, oldValue.PathLastPage404, oldValue.PathLastPageSiteId, oldValue.Sign*-1, oldValue.Version,
	)
	if err != nil {
		return err
	}

	err = pc.chClient.Exec(pc.ctx,
		`INSERT INTO path_cache (uuid, sessionUuid, dateHit, pathId, pathPages, pathFirstPage, pathFirstPage404,
                        pathFirstPageSiteId, pathLastPage, pathLastPage404, pathPageSiteId, pathSteps, isLastPage,sign,version) 
			   VALUES (generateUUIDv7(),?,now(),?,?,?,?,?,?,?,?,?,?,?,?)`,
		oldValue.SessionUuid, oldValue.DateHit, oldValue.PathId, oldValue.PathPages, oldValue.PathFirstPage, oldValue.PathFirstPage404, oldValue.PathFirstPageSiteId,
		oldValue.PathLastPage, oldValue.PathLastPage404, oldValue.PathLastPageSiteId, 1, oldValue.Version+1,
	)
	if err != nil {
		return err
	}

	return nil
}
