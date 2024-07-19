package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/utils"
	"context"
	"database/sql"
	"errors"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
	"strings"
)

type Searcher struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewSearcher(ctx context.Context, chClient driver.Conn) *Searcher {
	return &Searcher{
		ctx:      ctx,
		chClient: chClient,
	}
}

func (s Searcher) ExistStatDayForSearcher(searcherUuid uuid.UUID) bool {
	sql := `select count(uuid) from searcher_day where uuid = ?`
	row := s.chClient.QueryRow(s.ctx, sql, searcherUuid)
	var value int
	err := row.Scan(&value)
	if err != nil {
		return false
	}

	if value != 0 {
		return true
	}

	return false
}

func (s Searcher) FindSearcherByUserAgent(httpUserAgent string) (entitydb.SearcherDb, error) {
	var searcher entitydb.SearcherDb
	resultSql := `SELECT	uuid, date_cleanup, total_hits, save_statistic,
    				active, name, user_agent, diagram_default, hit_keep_days, dynamic_keep_days, phrases, phrases_hits, check_activity    
			FROM searcher
			WHERE active = 'Y' and LENGTH(user_agent)>0	and user_agent like ? 
			ORDER BY LENGTH(user_agent) desc 
			LIMIT 1`
	userAgent := utils.StringConcat("%", strings.ToUpper(httpUserAgent), "%")
	err := s.chClient.QueryRow(s.ctx, resultSql, userAgent, userAgent).ScanStruct(&searcher)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return entitydb.SearcherDb{}, err
	}

	return searcher, nil
}

func (s Searcher) UpdateSearcherDay(searcherUuid uuid.UUID) error {
	err := s.chClient.Exec(s.ctx, "UPDATE searcher_day SET date_stat=CURRENT_DATE(), date_last=NOW(),total_hits = total_hits + 1 WHERE uuid=?", searcherUuid)
	if err != nil {
		return err
	}
	return nil
}

func (s Searcher) AddSearcherDay(searcherUuid uuid.UUID) error {
	err := s.chClient.Exec(s.ctx,
		"INSERT INTO  searcher_day (uuid,date_stat,date_last,searcher_uuid,total_hits) VALUES (generateUUIDv7(),CURRENT_DATE,NOW(),?,1)",
		searcherUuid,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s Searcher) ExistByIdAndCurrentDate(id int) ([]entitydb.SearcherDayDb, error) {
	var rows []entitydb.SearcherDayDb
	//sql := `SELECT id FROM searcher_day WHERE searcher_id='?' and date_stat=CURRENT_DATE ORDER BY id`
	//err := s.storage.DB().Select(&rows, sql, id)
	//if err != nil {
	//	return nil, err
	//}

	return rows, nil
}

func (s Searcher) AddSearcherHit(searcherId uuid.UUID, searcherHitKeepDays uint32, statData entityjson.StatData) error {
	err := s.chClient.Exec(s.ctx,
		`INSERT INTO searcher_hit (uuid,date_hit,searcher_uuid,url,url_404,ip,user_agent,hit_keep_days,site_id)
			   VALUES(generateUUIDv7(),NOW(),?,?,?,?,?,?,?)`, searcherId, statData.Url, statData.IsError404, statData.Ip,
		statData.UserAgent, searcherHitKeepDays, statData.SiteId)
	if err != nil {
		return err
	}
	return nil
}
