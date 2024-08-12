package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

// Path  Пути
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

func (p Path) Find(filter filters.Filter) {

}

func (p Path) Add(path entitydb.Path) error {
	return p.chClient.Exec(p.ctx,
		`INSERT INTO path (uuid, parent_path_id, date_stat, pages, first_page, first_page_site_id, prev_page, prev_page_hash, last_page, last_page_site_id, last_page_hash, sign, version) 
			   VALUES (generateUUIDv7(),?,?,?,?,?,?,?,?,?,?)`,
		path.ParentPathId, path.Pages, path.FirstPage, path.FirstPageSiteId, path.PrevPage, path.PrevPageHash, path.LastPage, path.LastPageSiteId,
		path.LastPageHash, path.Sign*-1, path.Version)
}

func (p Path) FindByPathId(pathId int32, dateStat string) (entitydb.Path, error) {
	var path entitydb.Path
	err := p.chClient.QueryRow(p.ctx, `SELECT * FROM path WHERE path_id = ?`, pathId).ScanStruct(&path)

	if err != nil {
		return entitydb.Path{}, err
	}
	return path, nil
}

func (p Path) Update(oldPath entitydb.Path, newPath entitydb.Path) error {
	err := p.chClient.Exec(p.ctx,
		`INSERT INTO path (uuid, parent_path_id, date_stat, pages, first_page, first_page_site_id, prev_page, prev_page_hash, last_page, last_page_site_id, last_page_hash, sign, version)
			   VALUES (generateUUIDv7(),?,curdate(),?,?,?,?,?,?,?,?,?,?)`,
		oldPath.ParentPathId, oldPath.Pages, oldPath.FirstPage, oldPath.FirstPageSiteId, oldPath.PrevPage, oldPath.PrevPageHash, oldPath.LastPage, oldPath.LastPageSiteId,
		oldPath.LastPageHash, oldPath.Sign*-1, oldPath.Version)
	if err != nil {
		return err
	}

	err = p.chClient.Exec(p.ctx,
		`INSERT INTO path (uuid, parent_path_id, date_stat, pages, first_page, first_page_site_id, prev_page, prev_page_hash, last_page, last_page_site_id, last_page_hash, sign, version) 
			   VALUES (generateUUIDv7(),?,?,?,?,?,?,?,?,?,?)`,
		newPath.ParentPathId, newPath.Pages, newPath.FirstPage, newPath.FirstPageSiteId, newPath.PrevPage, newPath.PrevPageHash, newPath.LastPage, newPath.LastPageSiteId,
		newPath.LastPageHash, 1, newPath.Version+1)

	return nil
}
