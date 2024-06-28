package models

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/storage"
	"database/sql"
	"errors"
)

type GuestModel struct {
	storage      storage.Storage
	sessionModel SessionModel
}

func NewGuestModel(storage storage.Storage) *GuestModel {
	return &GuestModel{
		storage:      storage,
		sessionModel: NewSessionModel(storage),
	}
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
func (gm GuestModel) Add(guest entity.GuestDb) {
	gm.storage.DB().MustExec(`INSERT INTO guest (favorites,c_events,sessions,hits,repair,first_session_id,
                   	first_date,first_url_from,first_url_to,first_url_to_404,first_site_id,first_adv_id,
					first_referer1,first_referer2,first_referer3,last_session_id,last_date,last_user_id,
					last_user_auth,last_url_last,last_url_last_404,last_user_agent,last_ip,last_cookie,last_language,
					last_adv_id,last_adv_back,last_referer1,last_referer2,last_referer3,last_site_id,
					last_country_id,last_city_id,last_city_info,cookie_token) 
		VALUES (?,?,?,?,?,?,now(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		guest.Favorites, guest.CEvents, guest.Sessions, guest.Hits, guest.Repair, guest.FirstSessionId,
		guest.FirstUrlFrom, guest.FirstUrlTo, guest.FirstUrlTo404, guest.FirstSiteId, guest.FirstAdvId,
		guest.FirstReferer1, guest.FirstReferer2, guest.FirstReferer3, guest.LastSessionId, guest.LastDate, guest.LastUserId,
		guest.LastUserAuth, guest.LastUrlLast, guest.LastUrlLast404, guest.LastUserAgent, guest.LastIp, guest.LastCookie, guest.LastLanguage,
		guest.LastAdvId, guest.LastAdvBack, guest.LastReferer1, guest.LastReferer2, guest.LastReferer3, guest.LastSiteId,
		guest.LastCountryId, guest.LastCityId, guest.LastCityInfo, guest.CookieToken)
}

func (gm GuestModel) AddGuest(statData entity.StatData) error {
	gm.Add(entity.GuestDb{
		CookieToken:   statData.CookieToken,
		FirstUrlFrom:  statData.Referer,
		FirstUrlTo:    statData.Url,
		FirstUrlTo404: statData.Error404,
		FirstSiteId:   statData.SiteId,
		FirstAdvId:    0,  //TODO добавить реальные значения
		FirstReferer1: "", //TODO добавить реальные значения
		FirstReferer2: "", //TODO добавить реальные значения
		FirstReferer3: "", //TODO добавить реальные значения
	})
	return nil
}

func (gm GuestModel) ExistsGuestByToken(token string) (bool, error) {
	row := gm.storage.DB().QueryRow(`
				SELECT cookie_token 
				FROM guest 				
				WHERE cookie_token=?`, token)
	var cookieToken string
	err := row.Scan(&cookieToken)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return len(cookieToken) > 0, nil
}
