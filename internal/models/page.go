package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
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
			   WHERE date_stat = ?	and (url = ? and dir = 'Y') or (url = ? and dir ='N')`,
		dir, page, stat)
	defer rows.Close()
	if err != nil {
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

func (p *Page) FindByUuid(uuid string) (entitydb.Page, error) {
	var pageDb entitydb.Page
	err := p.chClient.QueryRow(p.ctx, `SELECT * FROM page where uuid = ?`, uuid).ScanStruct(&pageDb)
	if err != nil {
		return entitydb.Page{}, err
	}
	return pageDb, nil
}

func (p *Page) Update(oldValue entitydb.Page, newValue entitydb.Page) error {
	err := p.chClient.Exec(p.ctx,
		`INSERT INTO page (uuid, date_stat, dir, url, url_404, url_hash, site_id, counter, enter_counter, exit_counter,sign,version) 
			   VALUES (generateUUIDv7(),curdate(),?,?,?,?,?,?,?,?,?,?)`,
		oldValue.Dir, oldValue.UrlHash, oldValue.Url404, oldValue.UrlHash, oldValue.SiteId, oldValue.Counter, oldValue.EnterCounter, oldValue.ExitCounter, oldValue.Sign*-1, oldValue.Version,
	)

	if err != nil {
		return err
	}

	return p.chClient.Exec(p.ctx,
		`INSERT INTO page (uuid, date_stat, dir, url, url_404, url_hash, site_id, counter, enter_counter, exit_counter,sign,version) 
			   VALUES (generateUUIDv7(),curdate(),?,?,?,?,?,?,?,?,?,?)`,
		newValue.Dir, newValue.UrlHash, newValue.Url404, newValue.UrlHash, newValue.SiteId, newValue.Counter, newValue.EnterCounter, newValue.ExitCounter, newValue.Sign, newValue.Version+1,
	)
}
