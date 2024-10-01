package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"context"
	"database/sql"
	"errors"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Session struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewSession(ctx context.Context, chClient driver.Conn) *Session {
	return &Session{ctx: ctx, chClient: chClient}
}

func (s Session) Find(filter filters.Filter) (error, []map[string]interface{}) {

	return nil, nil
}

func (s Session) Add(session entitydb.Session) error {
	return s.chClient.Exec(s.ctx,
		`INSERT INTO session (uuid, guest_uuid, new_guest, user_id, user_auth, events, hits, favorites, url_from, url_to, url_to_404, url_last,
                     url_last_404, user_agent, date_stat, date_first, date_last, ip_first, ip_last, first_hit_uuid, last_hit_uuid, php_session_id, adv_uuid, adv_back, referer1, referer2, referer3, 
                     stop_list_uuid, country_id, first_site_id, last_site_id, city_id, sign, version) 
					VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,curdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		session.Uuid, session.GuestUuid, session.IsNewGuest, session.UserId, session.IsUserAuth, session.Events, session.Hits, session.Favorites, session.UrlFrom, session.UrlTo, session.UrlTo404, session.UrlLast, session.UrlLast404,
		session.UserAgent, session.DateFirst, session.DateLast, session.IpFirst, session.IpLast, session.FirstHitUuid, session.LastHitUuid, session.PhpSessionId, session.AdvUuid, session.AdvBack, session.Referer1, session.Referer2,
		session.Referer3, session.StopListUuid, session.CountryId, session.FirstSiteId, session.LastSiteId, session.CityId, session.Sign, session.Version,
	)
}

func (s Session) Update(oldSession entitydb.Session, newSession entitydb.Session) error {

	err := s.chClient.Exec(s.ctx,
		`INSERT INTO session (uuid, guest_uuid, new_guest, user_id, user_auth, events, hits, favorites, url_from, url_to, url_to_404, url_last,
                     url_last_404, user_agent, date_stat, date_first, date_last, ip_first, ip_last, first_hit_uuid, last_hit_uuid, php_session_id, adv_uuid, adv_back, referer1, referer2, referer3, 
                     stop_list_uuid, country_id, first_site_id, last_site_id, city_id, sign, version) 
					VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,curdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		oldSession.Uuid, oldSession.GuestUuid, oldSession.IsNewGuest, oldSession.UserId, oldSession.IsUserAuth, oldSession.Events, oldSession.Hits, oldSession.Favorites, oldSession.UrlFrom, oldSession.UrlTo, oldSession.UrlTo404,
		oldSession.UrlLast, oldSession.UrlLast404, oldSession.UserAgent, oldSession.DateFirst, oldSession.DateLast, oldSession.IpFirst, oldSession.IpLast, oldSession.FirstHitUuid, oldSession.LastHitUuid, oldSession.PhpSessionId,
		oldSession.AdvUuid, oldSession.AdvBack, oldSession.Referer1, oldSession.Referer2, oldSession.Referer3, oldSession.StopListUuid, oldSession.CountryId, oldSession.FirstSiteId, oldSession.LastSiteId, oldSession.CityId,
		oldSession.Sign, oldSession.Version,
	)

	if err != nil {
		return err
	}

	err = s.chClient.Exec(s.ctx,
		`INSERT INTO session (uuid, guest_uuid, new_guest, user_id, user_auth, events, hits, favorites, url_from, url_to, url_to_404, url_last,
                     url_last_404, user_agent, date_stat, date_first, date_last, ip_first, ip_last, first_hit_uuid, last_hit_uuid, php_session_id, adv_uuid, adv_back, referer1, referer2, referer3, 
                     stop_list_uuid, country_id, first_site_id, last_site_id, city_id, sign, version) 
					VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,curdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		newSession.Uuid, newSession.GuestUuid, newSession.IsNewGuest, newSession.UserId, newSession.IsUserAuth, newSession.Events, newSession.Hits, newSession.Favorites, newSession.UrlFrom, newSession.UrlTo, newSession.UrlTo404,
		newSession.UrlLast, newSession.UrlLast404, newSession.UserAgent, newSession.DateFirst, newSession.DateLast, newSession.IpFirst, newSession.IpLast, newSession.FirstHitUuid, newSession.LastHitUuid, newSession.PhpSessionId,
		newSession.AdvUuid, newSession.AdvBack, newSession.Referer1, newSession.Referer2, newSession.Referer3, newSession.StopListUuid, newSession.CountryId, newSession.FirstSiteId, newSession.LastSiteId, newSession.CityId,
		newSession.Sign, newSession.Version,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s Session) GetAttentiveness(dateStat, siteId string) {
	//sqlSite := ""
	//if len(siteId) > 0 {
	//	sqlSite = " and S.FIRST_SITE_ID = '" + siteId + "' " //TODO переделать
	//} else {
	//	sqlSite = ""
	//}
	//strSql := `
	//SELECT
	//sum(UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST))/count(S.ID)		AM_AVERAGE_TIME,
	//	sum(if(UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)<60,1,0))		AM_1,
	//	sum(if(UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)>=60
	//	and UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)<180,1,0))		AM_1_3,
	//	sum(if(UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)>=180
	//	and UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)<360,1,0))		AM_3_6,
	//	sum(if(UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)>=360
	//	and UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)<540,1,0))		AM_6_9,
	//	sum(if(UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)>=540
	//	and UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)<720,1,0))		AM_9_12,
	//	sum(if(UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)>=720
	//	and UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)<900,1,0))		AM_12_15,
	//	sum(if(UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)>=900
	//	and UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)<1080,1,0))		AM_15_18,
	//	sum(if(UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)>=1080
	//	and UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)<1260,1,0))		AM_18_21,
	//	sum(if(UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)>=1260
	//	and UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)<1440,1,0))		AM_21_24,
	//	sum(if(UNIX_TIMESTAMP(S.DATE_LAST)-UNIX_TIMESTAMP(S.DATE_FIRST)>=1440,1,0))	AM_24,
	//
	//	sum(S.HITS)/count(S.ID)						AH_AVERAGE_HITS,
	//	sum(if(S.HITS<=1, 1, 0))					AH_1,
	//	sum(if(S.HITS>=2 and S.HITS<=5, 1, 0))		AH_2_5,
	//	sum(if(S.HITS>=6 and S.HITS<=9, 1, 0))		AH_6_9,
	//	sum(if(S.HITS>=10 and S.HITS<=13, 1, 0))	AH_10_13,
	//	sum(if(S.HITS>=14 and S.HITS<=17, 1, 0))	AH_14_17,
	//	sum(if(S.HITS>=18 and S.HITS<=21, 1, 0))	AH_18_21,
	//	sum(if(S.HITS>=22 and S.HITS<=25, 1, 0))	AH_22_25,
	//	sum(if(S.HITS>=26 and S.HITS<=29, 1, 0))	AH_26_29,
	//	sum(if(S.HITS>=30 and S.HITS<=33, 1, 0))	AH_30_33,
	//	sum(if(S.HITS>=34, 1, 0))					AH_34
	//FROM session S
	//WHERE
	//S.DATE_STAT = cast(".$->CharToDateFunction($DATE_STAT, "SHORT")." as date)
	//$str
	//";
	//
	//$rs = $->Query($strSql, false, $err_mess.__LINE__);
	//$ar = $rs->Fetch();
	//$arKeys = array_keys($ar);
	//foreach($arKeys as $key)
	//{
	//if ($key=="AM_AVERAGE_TIME" || $key=="AH_AVERAGE_HITS")
	//{
	//$ar[$key] = (float) $ar[$key];
	//$ar[$key] = round($ar[$key],2);
	//}
	//else
	//{
	//$ar[$key] = intval($ar[$key]);
	//}
	//}
	//return $ar;
}

func (s Session) ExistsByPhpSession(session string) (bool, error) {
	var count uint8
	row := s.chClient.QueryRow(s.ctx, `select 1 as cnt from session where php_session_id=?`, session)

	err := row.Scan(&count)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return false, err
	}
	return count > 0, nil
}

func (s Session) FindByPHPSessionId(phpSessionId string) (entitydb.Session, error) {
	var sessionDb entitydb.Session
	err := s.chClient.QueryRow(s.ctx, `select * from session where php_session_id=?`, phpSessionId).ScanStruct(&sessionDb)
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		return entitydb.Session{}, err
	}
	return sessionDb, nil
}

func (s Session) FindAll(skip uint32, limit uint32) ([]entitydb.Session, error) {
	if limit > 1000 || limit < 1 {
		limit = 1000
	}
	resultSql := `select * from session limit ?, ?`
	rows, err := s.chClient.Query(s.ctx, resultSql, skip, limit)

	if err != nil {
		return nil, err
	}

	var allDbSessions []entitydb.Session
	for rows.Next() {
		var hit entitydb.Session
		err = rows.ScanStruct(&hit)
		allDbSessions = append(allDbSessions, hit)
	}

	return allDbSessions, nil
}
