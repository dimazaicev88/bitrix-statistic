package models

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
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
	resultSql := `SELECT	uuid, date_cleanup, totalHits, saveStatistic,
    				active, name, userAgent, diagramDefault, hitKeepDays, dynamicKeepDays,checkActivity    
			FROM searcher
			WHERE active = 'Y' and LENGTH(userAgent)>0	and userAgent like ? 
			ORDER BY LENGTH(userAgent) desc 
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

func (s Searcher) AddHitSearcher(searcherUuid uuid.UUID, statData dto.UserData) error {
	return s.chClient.Exec(s.ctx,
		`INSERT INTO searcher_hit (uuid,dateHit,searcherUuid,url,url404,ip,userAgent,siteId)
			   VALUES(generateUUIDv7(),NOW(),?,?,?,?,?,?)`, searcherUuid, statData.Url, statData.IsError404, statData.Ip,
		statData.UserAgent, statData.SiteId)
}

func (s Searcher) AddSearcherDayHits(searcherUuid uuid.UUID) error {
	return s.chClient.Exec(s.ctx,
		`insert into searcher_day_hits (hitDay, searcherUuid, totalHits)
				VALUES (curdate(), ?, 1)`, searcherUuid)
}

func (s Searcher) FindSearcherParamsByHost(host string) (entitydb.SearcherParams, error) {
	var searcherParams entitydb.SearcherParams
	resultSql := `SELECT t_searcher_params.*
			FROM searcher t_searcher
					 JOIN searcher_params t_searcher_params ON t_searcher.uuid = t_searcher_params.searcherUuid
			WHERE t_searcher.active = 'Y'
			  and ? like t_searcher_params.domain`
	rows, err := s.chClient.Query(s.ctx, resultSql, host)

	if err != nil {
		return entitydb.SearcherParams{}, err
	}

	if rows.Next() {
		if err := rows.ScanStruct(&searcherParams); err != nil {
			return entitydb.SearcherParams{}, err
		}
	}
	return searcherParams, nil
}

func (s Searcher) AddPhraseList(list entitydb.PhraseList) error {
	return s.chClient.Exec(s.ctx,
		`INSERT INTO phrase_list (uuid, dateHit, searcherUuid, refererUuid, phrase, urlFrom, urlTo, sessionUuid, siteId)
			   VALUES(?,?,?,?,?,?,?,?,?)`, list.Uuid, list.DateHit, list.SearcherUuid, list.RefererUuid,
		list.Phrase, list.UrlFrom, list.UrlTo, list.SessionUuid, list.SiteId)
}

func (s Searcher) AddSearcherPhraseStat(searcherPhraseStat entitydb.SearcherPhraseStat) error {
	return s.chClient.Exec(s.ctx,
		`INSERT INTO searcher_phrase_stat (searcherUuid, phrases, phrasesHits)
			   VALUES(?,?,?)`, searcherPhraseStat.SearcherUuid, searcherPhraseStat.Phrases, searcherPhraseStat.PhrasesHits)
}

func (s Searcher) Find(filter filters.Filter) (any, error) {
	return nil, nil
}
