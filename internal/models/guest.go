package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"context"
	"database/sql"
	"errors"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

type Guest struct {
	ctx          context.Context
	chClient     driver.Conn
	sessionModel *Session
}

func NewGuest(ctx context.Context, chClient driver.Conn) *Guest {
	return &Guest{
		ctx:          ctx,
		chClient:     chClient,
		sessionModel: NewSession(ctx, chClient),
	}
}

func (gm Guest) Add(guest entitydb.Guest) error {
	return gm.chClient.Exec(gm.ctx,
		`INSERT INTO guest (uuid, date_add, favorites, events, sessions, hits, repair, first_session_uuid, php_session_id,
                   first_date, first_url_from, first_url_to, first_url_404, first_site_id, first_adv_uuid, first_referer1, first_referer2, 
                   first_referer3, last_session_uuid, last_date, last_user_id, last_user_auth, last_url_last, last_url_last_404, last_user_agent, 
                   last_ip, last_cookie, last_language, last_adv_uuid, last_adv_back, last_referer1, last_referer2, last_referer3, last_site_id, 
                   last_country_id, last_city_id, last_city_info, sign, version) VALUES (?,now(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,
                                                                                         ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		guest.Uuid, guest.Favorites, guest.Events, guest.Sessions, guest.Hits, guest.Repair, guest.FirstSessionUuid, guest.PhpSessionId, guest.FirstDate, guest.FirstUrlFrom,
		guest.FirstUrlTo, guest.FirstUrlTo404, guest.FirstSiteId, guest.FirstAdvUuid, guest.FirstReferer1, guest.FirstReferer2, guest.FirstReferer3, guest.LastSessionUuid,
		guest.LastDate, guest.LastUserId, guest.LastUserAuth, guest.LastUrlLast, guest.LastUrlLast404, guest.LastUserAgent, guest.LastIp, guest.LastCookie,
		guest.LastLanguage, guest.LastAdvUUid, guest.LastAdvBack, guest.LastReferer1, guest.LastReferer2, guest.LastReferer3, guest.LastSiteId, guest.LastCountryId,
		guest.LastCityId, guest.LastCityInfo, guest.Sign, guest.Version,
	)
}

func (gm Guest) Find(filter filters.Filter) ([]entitydb.Guest, error) {
	return []entitydb.Guest{}, nil
}

func (gm Guest) FindByUuid(uuid uuid.UUID) (entitydb.Guest, error) {
	var hit entitydb.Guest
	err := gm.chClient.QueryRow(gm.ctx, `select uuid, date_add, favorites, events, sessions, hits, 
       repair, first_session_uuid, php_session_id, first_date, first_url_from, first_url_to, 
       first_url_404, first_site_id, first_adv_uuid, first_referer1, first_referer2, first_referer3,
       last_session_uuid, last_date, last_user_id, last_user_auth, last_url_last, last_url_last_404, 
       last_user_agent, last_ip, last_cookie, last_language, last_adv_uuid, last_adv_back, last_referer1, last_referer2, 
       last_referer3, last_site_id, last_country_id, last_city_id, last_city_info, sign, version from guest where uuid=?`,
		uuid).ScanStruct(&hit)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return entitydb.Guest{}, err
	}
	return hit, nil
}

func (gm Guest) Update(oldValue, newValue entitydb.Guest) error {
	err := gm.chClient.Exec(gm.ctx,
		`INSERT INTO guest (uuid, date_add, favorites, events, sessions, hits, repair, first_session_uuid, php_session_id, first_date, first_url_from, 
                   first_url_to, first_url_404, first_site_id, first_adv_uuid, first_referer1, first_referer2, first_referer3, last_session_uuid, last_date, 
                   last_user_id, last_user_auth, last_url_last, last_url_last_404, last_user_agent, last_ip, last_cookie, last_language, last_adv_uuid, last_adv_back, 
                   last_referer1, last_referer2, last_referer3, last_site_id, last_country_id, last_city_id, last_city_info, sign, version)
               VALUES (?,now(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		oldValue.Uuid, oldValue.Favorites, oldValue.Events, oldValue.Sessions, oldValue.Hits, oldValue.Repair, oldValue.FirstSessionUuid, oldValue.PhpSessionId, oldValue.FirstDate,
		oldValue.FirstUrlFrom, oldValue.FirstUrlTo, oldValue.FirstUrlTo404, oldValue.FirstSiteId, oldValue.FirstAdvUuid, oldValue.FirstReferer1, oldValue.FirstReferer2,
		oldValue.FirstReferer3, oldValue.LastSessionUuid, oldValue.LastDate, oldValue.LastUserId, oldValue.LastUserAuth, oldValue.LastUrlLast, oldValue.LastUrlLast404,
		oldValue.LastUserAgent, oldValue.LastIp, oldValue.LastCookie, oldValue.LastLanguage, oldValue.LastAdvUUid, oldValue.LastAdvBack, oldValue.LastReferer1,
		oldValue.LastReferer2, oldValue.LastReferer3, oldValue.LastSiteId, oldValue.LastCountryId, oldValue.LastCityId, oldValue.LastCityInfo,
		oldValue.Sign, oldValue.Version,
	)
	if err != nil {
		return err
	}

	err = gm.chClient.Exec(gm.ctx,
		`INSERT INTO guest (uuid, date_add, favorites, events, sessions, hits, repair, first_session_uuid, php_session_id, first_date, first_url_from, first_url_to,
                   first_url_404, first_site_id, first_adv_uuid, first_referer1, first_referer2, first_referer3, last_session_uuid, last_date, last_user_id, last_user_auth, last_url_last,
                   last_url_last_404, last_user_agent, last_ip, last_cookie, last_language, last_adv_uuid, last_adv_back, last_referer1, last_referer2, last_referer3, last_site_id, last_country_id, 
                   last_city_id, last_city_info, sign, version) 
				VALUES (?,now(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		newValue.Uuid, newValue.Favorites, newValue.Events, newValue.Sessions, newValue.Hits, newValue.Repair, newValue.FirstSessionUuid, newValue.PhpSessionId, newValue.FirstDate, newValue.FirstUrlFrom,
		newValue.FirstUrlTo, newValue.FirstUrlTo404, newValue.FirstSiteId, newValue.FirstAdvUuid, newValue.FirstReferer1, newValue.FirstReferer2, newValue.FirstReferer3,
		newValue.LastSessionUuid, newValue.LastDate, newValue.LastUserId, newValue.LastUserAuth, newValue.LastUrlLast, newValue.LastUrlLast404, newValue.LastUserAgent,
		newValue.LastIp, newValue.LastCookie, newValue.LastLanguage, newValue.LastAdvUUid, newValue.LastAdvBack, newValue.LastReferer1, newValue.LastReferer2, newValue.LastReferer3,
		newValue.LastSiteId, newValue.LastCountryId, newValue.LastCityId, newValue.LastCityInfo, newValue.Sign, newValue.Version,
	)
	if err != nil {
		return err
	}

	return nil
}
