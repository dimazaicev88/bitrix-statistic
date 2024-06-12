package models

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/storage"
)

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

// Add TODO доделать
func (gm GuestModel) Add(guest entity.Guest) {
	gm.storage.DB().MustExec(`INSERT INTO guest (favorites,c_events,sessions,hits,repair,first_session_id,
                   	first_date,first_url_from,first_url_to,first_url_to_404,first_site_id,first_adv_id,
					first_referer1,first_referer2,first_referer3,last_session_id,last_date,last_user_id,
					last_user_auth,last_url_last,last_url_last_404,last_user_agent,last_ip,last_cookie,last_language,
					last_adv_id,last_adv_back,last_referer1,last_referer2,last_referer3,last_site_id,
					last_country_id,last_city_id,last_city_info) 
		VALUES (?,?,?,?,?,?,now(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		guest.Favorites, guest.CEvents, guest.Sessions, guest.Hits, guest.Repair, guest.FirstSessionId,
		guest.FirstDate, guest.FirstUrlFrom, guest.FirstUrlTo, guest.FirstUrlTo404, guest.FirstSiteId,
		guest.FirstAdvId, guest.FirstReferer1, guest.FirstReferer2)
}

func NewGuestModel(storage storage.Storage) *GuestModel {
	return &GuestModel{storage: storage}
}
