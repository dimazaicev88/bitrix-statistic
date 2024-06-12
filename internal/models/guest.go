package models

import "bitrix-statistic/internal/storage"

type GuestModel struct {
	storage storage.Storage
}

func (gm GuestModel) FindLastById(id int) (int, string, int, int, string, error) {
	row := gm.storage.DB().QueryRow(`
				SELECT
					G.id,
					G.FAVORITES,
					G.LAST_USER_ID,
					A.ID as LAST_ADV_ID,
					if(to_days(curdate())=to_days(G.LAST_DATE), 'Y', 'N') LAST
				FROM guest G
				LEFT JOIN adv A ON A.ID = G.LAST_ADV_ID
				WHERE G.ID=?`, id)
	var guestId, lastUserId, lastAdvId int
	var favorites, last string
	err := row.Scan(&guestId, favorites, lastUserId, lastAdvId, last)
	if err != nil {
		return 0, "", 0, 0, "", err
	}
	return guestId, favorites, lastUserId, lastAdvId, last, nil
}

func NewGuestModel(storage storage.Storage) *GuestModel {
	return &GuestModel{storage: storage}
}
