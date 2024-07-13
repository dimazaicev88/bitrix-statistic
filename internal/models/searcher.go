package models

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/storage"
)

type SearcherModel struct {
	storage storage.Storage
}

func NewSearcherModel(storage storage.Storage) *SearcherModel {
	return &SearcherModel{
		storage: storage,
	}
}

func (sm SearcherModel) ExistById(id int) bool {
	row := sm.storage.DB().QueryRow("SELECT id FROM  searcher WHERE id =?", id)
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

func (sm SearcherModel) FindSearcherByUserAgent(httpUserAgent string) ([]entity.SearcherDb, error) {
	var rows []entity.SearcherDb
	sql := `SELECT	id, name,  save_statistic, hit_keep_days, check_activity
			FROM searcher
			WHERE ACTIVE = 'Y' and LENGTH(user_agent)>0
				and upper(?) like CONCAT('%', upper(user_agent), '%')
			ORDER BY LENGTH("user_agent") desc, ID`

	err := sm.storage.DB().Select(&rows, sql, httpUserAgent)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (sm SearcherModel) UpdateSearcherDay(id int) {
	sm.storage.DB().MustExec("UPDATE searcher_day SET date_last=NOW(),total_hits = total_hits + 1 WHERE id=?", id)
}

func (sm SearcherModel) AddSearcherDay(id int) {
	sm.storage.DB().MustExec(
		"INSERT INTO  searcher_day (date_stat,date_last,searcher_id,total_hits) VALUES (CURRENT_DATE,NOW(),?,1)",
		id,
	)
}

func (sm SearcherModel) ExistByIdAndCurrentDate(id int) ([]entity.SearcherDayDb, error) {
	var rows []entity.SearcherDayDb
	sql := `SELECT id FROM searcher_day WHERE searcher_id='?' and date_stat=CURRENT_DATE ORDER BY id`
	err := sm.storage.DB().Select(&rows, sql, id)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (sm SearcherModel) AddSearcherHit(searcherId int, uri string, error404, ip, agent, searcherHitKeepDays, siteId string) {
	sm.storage.DB().MustExec(
		`INSERT INTO searcher_hit (date_hit,searcher_id,url,url_404,ip,user_agent,hit_keep_days,site_id) 
			   VALUES(CURRENT_DATE,?,?,?,?,?,?,?)`, searcherId, uri, error404, ip, agent, searcherHitKeepDays, siteId,
	)
}
