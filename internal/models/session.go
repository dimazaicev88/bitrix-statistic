package models

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/session"
	"bitrix-statistic/internal/storage"
	"encoding/json"
	"time"
)

type SessionModel struct {
	storage storage.Storage
}

func NewSessionModel(storageImpl storage.Storage) SessionModel {
	return SessionModel{storage: storageImpl}
}

func (sm SessionModel) Find(filter filters.Filter) (error, []map[string]interface{}) {

	return nil, nil
}

func (sm SessionModel) AddSession(session entity.Session) error {
	_, err := sm.storage.DB().MustExec(`INSERT INTO b_stat_session (GuestId, NewGuest, USER_ID, USER_AUTH, C_EVENTS, HITS, FAVORITES, URL_FROM, URL_TO, URL_TO_404, URL_LAST,
		URL_LAST_404, UserAgent, DATE_STAT, DATE_FIRST, DATE_LAST, IP_FIRST, IP_FIRST_NUMBER, IP_LAST, IP_LAST_NUMBER, FIRST_HIT_ID, LAST_HIT_ID, PHPSESSID,
        ADV_ID, ADV_BACK, REFERER1, REFERER2, REFERER3, StopListId, CountryId, CityId, FIRST_SITE_ID, LAST_SITE_ID) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 
                                                                                                                                ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) `,
		session.GuestId, session.NewGuest, session.UserId, session.UserAuth, session.CEvents, session.Hits, session.Favorites, session.UrlFrom, session.UrlTo, session.UrlTo404, session.UrlLast, session.UrlLast404, session.UserAgent, time.Unix(session.DateStat, 0).Add(time.Hour*3),
		time.Unix(session.DateFirst, 0).Add(time.Hour*3), time.Unix(session.DateLast, 0).Add(time.Hour*3), session.IpLast, session.IpFirstNumber, session.IpLast, session.IpLastNumber, session.FirstHitId, session.LastHitId, session.PhpSessionId,
		session.AdvId, session.AdvBack, session.Referer1, session.Referer2, session.Referer3, session.StopListId, session.CountryId, session.CityId, session.FirstSiteId, session.LastSiteId).LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (sm SessionModel) DeleteById(id int) {
	sm.storage.DB().MustExec("DELETE FROM b_stat_session WHERE ID=?", id)
}

func (sm SessionModel) FindSessionByGuestMd5(guestMd5 string, sessionGcMaxLifeTime string) (session.Session, error) {
	row := sm.storage.DB().QueryRow(
		`SELECT session_data FROM session_data
                    WHERE guest_md5=? and date_last > DATE_ADD(now(), INTERVAL-? SECOND) 
                    LIMIT 1`,
		guestMd5, sessionGcMaxLifeTime,
	)
	var sessionDb string
	var sessionData session.Session
	err := row.Scan(&sessionData)
	if err != nil {
		return session.Session{}, err
	}
	if err := json.Unmarshal([]byte(sessionDb), &sessionData); err != nil {
		return session.Session{}, err
	}
	return sessionData, nil
}
