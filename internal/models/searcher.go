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

func (s Searcher) FindSearcherByUserAgent(httpUserAgent string) (entitydb.Searcher, error) {
	var searcher entitydb.Searcher
	resultSql := `SELECT	uuid, date_cleanup, total_hits, save_statistic,
    				active, name, user_agent, diagram_default, hit_keep_days, dynamic_keep_days, phrases, phrases_hits, check_activity    
			FROM searcher
			WHERE active = 'Y' and LENGTH(user_agent)>0	and user_agent like ? 
			ORDER BY LENGTH(user_agent) desc 
			LIMIT 1`
	userAgent := utils.StringConcat("%", strings.ToUpper(httpUserAgent), "%")
	err := s.chClient.QueryRow(s.ctx, resultSql, userAgent, userAgent).ScanStruct(&searcher)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return entitydb.Searcher{}, err
	}

	return searcher, nil
}

func (s Searcher) ExistByIdAndCurrentDate(id int) ([]entitydb.SearcherDayHits, error) {
	var rows []entitydb.SearcherDayHits
	//sql := `SELECT id FROM searcher_day WHERE searcher_id='?' and date_stat=CURRENT_DATE ORDER BY id`
	//err := s.storage.().Select(&rows, sql, id)
	//if err != nil {
	//	return nil, err
	//}

	return rows, nil
}

func (s Searcher) AddHitSearcher(searcherUuid uuid.UUID, statData entityjson.StatData) error {
	err := s.chClient.Exec(s.ctx,
		`INSERT INTO searcher_hit (uuid,date_hit,searcher_uuid,url,url_404,ip,user_agent,site_id)
			   VALUES(generateUUIDv7(),NOW(),?,?,?,?,?,?)`, searcherUuid, statData.Url, statData.IsError404, statData.Ip,
		statData.UserAgent, statData.SiteId)
	if err != nil {
		return err
	}
	return nil
}

func (s Searcher) AddSearcherDayHits(searcherUuid uuid.UUID) error {
	err := s.chClient.Exec(s.ctx,
		`insert into searcher_day_hits (hit_day, searcher_uuid, total_hits)
				VALUES (curdate(), ?, 1)`, searcherUuid)
	if err != nil {
		return err
	}
	return nil
}
