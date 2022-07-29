package models

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/storage"
)

type SessionModel struct {
	storage storage.Storage
}

func (sm SessionModel) Find(filter filters.Filter) (error, []map[string]interface{}) {

	return nil, nil
}

func (sm SessionModel) Add(session entity.Session) (error, []map[string]interface{}) {
	_, err := sm.storage.DB().MustExec(`INSERT INTO b_stat_session (GUEST_ID, NEW_GUEST, USER_ID, USER_AUTH, C_EVENTS, HITS, FAVORITES, URL_FROM, URL_TO, URL_TO_404, URL_LAST,
		URL_LAST_404, USER_AGENT, DATE_STAT, DATE_FIRST, DATE_LAST, IP_FIRST, IP_FIRST_NUMBER, IP_LAST, IP_LAST_NUMBER, FIRST_HIT_ID, LAST_HIT_ID, PHPSESSID,
        ADV_ID, ADV_BACK, REFERER1, REFERER2, REFERER3, STOP_LIST_ID, COUNTRY_ID, CITY_ID, FIRST_SITE_ID, LAST_SITE_ID) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 
                                                                                                                                ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) `,
		session.GuestId, session.NewGuest, session.UserId, session.UserAuth, session.CEvents, session.Hits, session.Favorites, session.UrlFrom, session.UrlTo, session.UrlTo404, session.UrlLast,
		session.AdvId, session.AdvBack, session.Referer1, session.Referer2, session.Referer3, session.StopListId, session.CountryId, session.CityId, session.FirstSiteId, session.LastSiteId).LastInsertId()
	if err != nil {
		return err, nil
	}
	return nil, nil
}

func NewSessionModel(storageImpl storage.Storage) SessionModel {
	return SessionModel{storage: storageImpl}
}
