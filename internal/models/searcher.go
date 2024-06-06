package models

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/storage"
)

type Searcher struct {
	storage storage.Storage
}

func NewSearcherModel(storage storage.Storage) *Searcher {
	return &Searcher{
		storage: storage,
	}
}

func (s Searcher) ExistById(id int) bool {
	row := s.storage.DB().QueryRow("SELECT id FROM  searcher WHERE id =?", id)
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

func (s Searcher) FindByUserAgent(httpUserAgent string) ([]entity.Searcher, error) {
	var rows []entity.Searcher
	sql := `SELECT
	id, name,  SAVE_STATISTIC, HIT_KEEP_DAYS, CHECK_ACTIVITY
	FROM
	searcher
	WHERE
	ACTIVE = 'Y'
	and LENGTH(USER_AGENT)>0
	and upper(?) like CONCAT("'%'", "upper(USER_AGENT)", "'%'")
	ORDER BY LENGTH("USER_AGENT") desc, ID`

	err := s.storage.DB().Select(&rows, sql, httpUserAgent)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (s Searcher) ExistByIdAndCurrentDate(id int) ([]entity.Searcher, error) {
	var rows []entity.Searcher
	sql := `SELECT ID FROM b_stat_searcher_day WHERE SEARCHER_ID='?' and DATE_STAT=CURRENT_DATE ORDER BY ID`
	err := s.storage.DB().Select(&rows, sql, id)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
