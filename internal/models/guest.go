package models

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/filters"
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

func (gm GuestModel) Add(guest entity.GuestDb) {
	gm.storage.DB().MustExec(`INSERT INTO guest (
                   timestamp_x, favorites, events, sessions, hits, repair, session_id, date, url_from, url_to,
                   url_to_404, site_id, adv_id, referer1, referer2, referer3, user_id, user_auth, url, url_404, user_agent, ip,
                   cookie, language, adv_back, country_id, city_id, city_info, cookie_token) 
		VALUES (?,?,?,?,?,?,now(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		guest.Favorites, guest.Events, guest.Sessions, guest.Hits, guest.Repair, guest.SessionId, guest.Date,
		guest.UrlFrom, guest.UrlTo, guest.UrlTo404, guest.SiteId, guest.AdvId, guest.Referer1, guest.Referer2,
		guest.Referer3, guest.UserId, guest.UserAuth, guest.Url, guest.Url404, guest.UserAgent, guest.Ip,
		guest.Cookie, guest.Language, guest.AdvBack, guest.CountryId, guest.CityId, guest.CityInfo, guest.CookieToken,
	)
}

func (gm GuestModel) AddGuest(statData entity.StatData) error {
	gm.Add(entity.GuestDb{
		//CookieToken: statData.CookieToken,
		//UrlFrom:     statData.Url,
		//UrlTo:       statData.Url,
		//UrlTo404:    statData.Error404,
		//SiteId:      statData.SiteId,
		//AdvId:       0,  //TODO добавить реальные значения
		//Referer1:    "", //TODO добавить реальные значения
		//Referer2:    "", //TODO добавить реальные значения
		//Referer3:    "", //TODO добавить реальные значения
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

func (gm GuestModel) Find(filter filters.Filter) (entity.GuestDb, error) {
	return entity.GuestDb{}, nil
}
