package models

import "bitrix-statistic/internal/storage"

type SearcherHitModel struct {
	storage storage.Storage
}

func (shm SearcherHitModel) Add(searcherId int, uri string, error404, ip, agent, searcherHitKeepDays, siteId string) {
	shm.storage.DB().MustExec(
		`INSERT INTO searcher_hit (date_hit,searcher_id,url,url_404,ip,user_agent,hit_keep_days,site_id) 
			   VALUES(CURRENT_DATE,?,?,?,?,?,?,?)`, searcherId, uri, error404, ip, agent, searcherHitKeepDays, siteId,
	)
}
