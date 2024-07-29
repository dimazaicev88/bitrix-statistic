package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/utils"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Adv struct {
	chClient driver.Conn
	ctx      context.Context
}

func NewAdv(
	ctx context.Context,
	chClient driver.Conn,
) *Adv {
	return &Adv{
		chClient: chClient,
		ctx:      ctx,
	}
}

// FindAdvUuidByByPage Поиск Рекламной компании по странице
func (am Adv) FindAdvUuidByByPage(page, direction string) ([]string, error) {
	strSql := `
		SELECT t_adv.uuid
		FROM adv t_adv
		INNER JOIN adv_page t_adv_page  ON (t_adv_page.adv_uuid = t_adv.uuid and t_adv_page.type=?)
 		WHERE length(t_adv_page.page) > 0 and t_adv_page.page like ?`

	var listAdvUuid []string

	rows, err := am.chClient.Query(am.ctx, strSql, direction, utils.StringConcat("%", page, "%"))
	if err != nil {
		return []string{}, err
	}

	for rows.Next() {
		var advUuid string
		if err := rows.Scan(&advUuid); err != nil {
			return []string{}, err
		}
		listAdvUuid = append(listAdvUuid, advUuid)
	}
	return listAdvUuid, nil
}

func (am Adv) FindByByDomainSearcher(host string) ([]string, error) {
	//проверяем поисковики
	sql := ` SELECT t_adv_searcher.adv_uuid
			FROM adv_searcher t_adv_searcher
					 JOIN searcher_params t_searcher_params ON t_adv_searcher.searcher_uuid = t_searcher_params.searcher_uuid
			WHERE t_searcher_params.domain like ?`

	var listAdvSearcherUuid []string
	rows, err := am.chClient.Query(am.ctx, sql, utils.StringConcat("%", host, "%"))
	if err != nil {
		return []string{}, err
	}

	for rows.Next() {
		var advUuid string
		if err := rows.Scan(&advUuid); err != nil {
			return []string{}, err
		}
		listAdvSearcherUuid = append(listAdvSearcherUuid, advUuid)
	}
	return listAdvSearcherUuid, nil
}

func (am Adv) FindByReferer(referer1, referer2 string) ([]string, error) {
	sql := `SELECT 	uuid
			FROM adv
			WHERE  referer1=? and referer2=?`

	var listUuid []string
	rows, err := am.chClient.Query(am.ctx, sql, referer1, referer2)
	if err != nil {
		return []string{}, err
	}

	for rows.Next() {
		var advUuid string
		if err = rows.Scan(&advUuid); err != nil {
			return []string{}, err
		}
		listUuid = append(listUuid, advUuid)
	}
	return listUuid, nil
}

func (am Adv) AddAdv(referer1 string, referer2 string) error {
	return am.chClient.Exec(am.ctx, `INSERT INTO adv (uuid, referer1, referer2, date_create, cost, events_view, description, priority)
		VALUES (generateUUIDv7(), ?, ?, now(),0.0,'','',1)`, referer1, referer2)
}

func (am Adv) FindByUuid(uuid string) (entitydb.Adv, error) {
	var adv entitydb.Adv
	sql := `SELECT 	* FROM adv WHERE  uuid=?`
	err := am.chClient.QueryRow(am.ctx, sql, uuid).ScanStruct(&adv)
	if err != nil {
		return entitydb.Adv{}, err
	}
	return adv, nil
}

func (am Adv) DeleteByUuid(uuid string) error {
	if err := am.chClient.Exec(am.ctx, `DELETE FROM adv WHERE uuid=?`, uuid); err != nil {
		return err
	}
	return nil
}

func (am Adv) FindRefererByListAdv(listAdv []string) (entitydb.AdvReferer, error) {
	var adv entitydb.AdvReferer
	sql := `SELECT 	referer1,referer2 FROM adv WHERE  uuid IN (?) ORDER BY priority,date_create DESC LIMIT 1`
	err := am.chClient.QueryRow(am.ctx, sql, listAdv).ScanStruct(&adv)
	if err != nil {
		return entitydb.AdvReferer{}, err
	}
	return adv, nil
}
