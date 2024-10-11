package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/utils"
	"context"
	"database/sql"
	"errors"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
	"time"
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
	resultSql := `
		SELECT t_adv.uuid
		FROM adv t_adv
		INNER JOIN adv_page t_adv_page  ON (t_adv_page.adv_uuid = t_adv.uuid and t_adv_page.type=?)
 		WHERE length(t_adv_page.page) > 0 and t_adv_page.page like ?`

	var listAdvUuid []string

	rows, err := am.chClient.Query(am.ctx, resultSql, direction, utils.StringConcat("%", page, "%"))
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
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
	resultSql := ` SELECT t_adv_searcher.adv_uuid
			FROM adv_searcher t_adv_searcher
					 JOIN searcher_params t_searcher_params ON t_adv_searcher.searcher_uuid = t_searcher_params.searcher_uuid
			WHERE t_searcher_params.domain like ?`

	var listAdvSearcherUuid []string
	rows, err := am.chClient.Query(am.ctx, resultSql, utils.StringConcat("%", host, "%"))
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
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
	resultSql := `SELECT uuid FROM adv WHERE  referer1=? and referer2=?`

	var listUuid []string
	rows, err := am.chClient.Query(am.ctx, resultSql, referer1, referer2)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
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

func (am Adv) AddAdv(referer1 string, referer2 string) (entitydb.Adv, error) {
	uuidAdv := uuid.New()
	timeAdd := time.Now()
	err := am.chClient.Exec(am.ctx, `INSERT INTO adv (uuid, referer1, referer2, date_create, cost, events_view, description)
	 		VALUES (?, ?, ?, ?,0.0,'','')`, uuidAdv, referer1, referer2, timeAdd)
	if err != nil {
		return entitydb.Adv{}, err
	}
	return entitydb.Adv{
		Uuid:        uuidAdv,
		Referer1:    referer1,
		Referer2:    referer2,
		DateCreated: timeAdd,
	}, nil
}

func (am Adv) AddAdvStat(advStat entitydb.AdvStat) error {
	return am.chClient.Exec(am.ctx, `INSERT INTO adv_stat (adv_uuid, guests, new_guests, favorites, hosts, sessions, hits, guests_back, favorites_back, hosts_back, sessions_back, hits_back)
	 		VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`, advStat.AdvUuid, advStat.Guests, advStat.NewGuests,
		advStat.Favorites, advStat.Hosts, advStat.Sessions, advStat.Hits, advStat.GuestsBack,
		advStat.FavoritesBack, advStat.HostsBack, advStat.SessionsBack, advStat.Hits, advStat.GuestsBack, advStat.HitsBack)
}

func (am Adv) FindByUuid(uuid uuid.UUID) (entitydb.Adv, error) {
	var adv entitydb.Adv
	resultSql := `SELECT * FROM adv WHERE  uuid=?`
	err := am.chClient.QueryRow(am.ctx, resultSql, uuid).ScanStruct(&adv)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return entitydb.Adv{}, err
	}
	return adv, nil
}

func (am Adv) DeleteByUuid(advUuid uuid.UUID) error {
	if err := am.chClient.Exec(am.ctx, `DELETE FROM adv WHERE uuid=?`, advUuid); err != nil {
		return err
	}
	return nil
}

func (am Adv) FindRefererByListAdv(listAdv []string) (entitydb.AdvCompany, error) {
	var adv entitydb.AdvCompany
	resultSql := `SELECT 	uuid as adv_uuid, referer1,referer2 FROM adv WHERE  uuid IN (?) ORDER BY priority,date_create DESC LIMIT 1`
	err := am.chClient.QueryRow(am.ctx, resultSql, listAdv).ScanStruct(&adv)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return entitydb.AdvCompany{}, err
	}
	return adv, nil
}

func (am Adv) IsExistsAdv(advUuid uuid.UUID) (bool, error) {
	rows, err := am.chClient.Query(am.ctx, `select 1 from adv where uuid=?`, advUuid)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return false, err
	}

	for rows.Next() {
		var isExists uint8
		err = rows.Scan(&isExists)
		if err != nil {
			return false, err
		}
		return isExists == 1, nil
	}

	return false, nil
}

func (am Adv) AddAdvDay(day entitydb.AdvDay) error {
	return am.chClient.Exec(am.ctx, `INSERT INTO adv_day (adv_uuid, date_stat, guests, guests_day, new_guests, favorites, hosts, hosts_day, sessions, hits, guests_back, guests_day_back, favorites_back, hosts_back, hosts_day_back, sessions_back, hits_back)
	 		VALUES (?,curdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, day.AdvUuid, day.Guests, day.GuestsDay, day.NewGuests, day.Favorites, day.Hosts, day.HostsDay, day.Sessions, day.Hits,
		day.Hits, day.GuestsBack, day.GuestsDayBack, day.FavoritesBack, day.HostsBack, day.HostsDayBack, day.SessionsBack, day.Hits)
}
