package models

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/storage"
)

type SearcherDayModel struct {
	storage storage.Storage
}

func (sd SearcherDayModel) Update(id int) {
	sd.storage.DB().MustExec("UPDATE searcher_day SET date_last=NOW(),total_hits = total_hits + 1 WHERE id=?", id)
}

func (sd SearcherDayModel) Add(id int) {
	sd.storage.DB().MustExec(
		"INSERT INTO  searcher_day (date_stat,date_last,searcher_id,total_hits) VALUES (CURRENT_DATE,NOW(),?,1)",
		id,
	)
}

func (sd SearcherDayModel) ExistByIdAndCurrentDate(id int) ([]entity.SearcherDay, error) {
	var rows []entity.SearcherDay
	sql := `SELECT ID FROM b_stat_searcher_day WHERE SEARCHER_ID='?' and DATE_STAT=CURRENT_DATE ORDER BY ID`
	err := sd.storage.DB().Select(&rows, sql, id)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
