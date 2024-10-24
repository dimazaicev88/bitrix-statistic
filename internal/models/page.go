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

type Page struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewPage(ctx context.Context, chClient driver.Conn) *Page {
	return &Page{
		ctx:      ctx,
		chClient: chClient,
	}
}

func (p *Page) Filter(filter filters.Filter) ([]entitydb.Page, error) {
	return nil, nil
}

func (p *Page) DynamicList(filter filters.Filter) ([]entitydb.Page, error) {
	return nil, nil
}

func (p *Page) FindByPageAndDir(dir string, page string, stat string) ([]entitydb.Page, error) {
	rows, err := p.chClient.Query(p.ctx,
		`SELECT *
			   FROM page
			   WHERE dateStat = ?	and (url = ? and dir = 'Y') or (url = ? and dir ='N')`,
		dir, page, stat)
	defer rows.Close()
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return nil, err
	}
	var results []entitydb.Page
	for rows.Next() {
		var pageDb entitydb.Page
		if err = rows.ScanStruct(&pageDb); err != nil {
			return nil, err
		}
		results = append(results, pageDb)
	}
	return results, nil
}

func (p *Page) FindByUuid(uuid uuid.UUID) (entitydb.Page, error) {
	var pageDb entitydb.Page
	err := p.chClient.QueryRow(p.ctx, `SELECT * FROM page where uuid = ?`, uuid).ScanStruct(&pageDb)
	if err != nil {
		return entitydb.Page{}, err
	}
	return pageDb, nil
}

func (p *Page) Update(oldValue entitydb.Page, newValue entitydb.Page) error {
	err := p.chClient.Exec(p.ctx,
		`INSERT INTO page (uuid, dateStat, dir, url, url404, urlHash, siteId, counter, enterCounter, exitCounter,sign,version) 
			   VALUES (generateUUIDv7(),curdate(),?,?,?,?,?,?,?,?,?,?)`,
		oldValue.Dir, oldValue.UrlHash, oldValue.Url404, oldValue.UrlHash, oldValue.SiteId, oldValue.Counter, oldValue.EnterCounter, oldValue.ExitCounter, oldValue.Sign*-1, oldValue.Version,
	)

	if err != nil {
		return err
	}

	return p.chClient.Exec(p.ctx,
		`INSERT INTO page (uuid, dateStat, dir, url, url404, urlHash, siteId, counter, enterCounter, exitCounter,sign,version) 
			   VALUES (generateUUIDv7(),curdate(),?,?,?,?,?,?,?,?,?,?)`,
		newValue.Dir, newValue.UrlHash, newValue.Url404, newValue.UrlHash, newValue.SiteId, newValue.Counter, newValue.EnterCounter, newValue.ExitCounter, newValue.Sign, newValue.Version+1,
	)
}

func (p *Page) Add(page entitydb.Page) error {
	return p.chClient.Exec(p.ctx,
		`INSERT INTO page (uuid, dateStat, dir, url, url404, urlHash, siteId, counter, enterCounter, exitCounter,sign,version) 
			   VALUES (generateUUIDv7(),curdate(),?,?,?,?,?,?,?,?,?,?)`,
		page.Dir, page.UrlHash, page.Url404, page.UrlHash, page.SiteId, page.Counter, page.EnterCounter, page.ExitCounter, page.Sign, page.Version,
	)
}
