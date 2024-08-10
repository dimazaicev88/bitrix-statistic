package models

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/filters"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

type Session struct {
	ctx      context.Context
	chClient driver.Conn
}

func NewSession(ctx context.Context, chClient driver.Conn) *Session {
	return &Session{ctx: ctx, chClient: chClient}
}

func (sm Session) Find(filter filters.Filter) (error, []map[string]interface{}) {

	return nil, nil
}

func (sm Session) Add(session entitydb.Session) (string, error) {
	sessionUuid := uuid.New().String()
	err := sm.chClient.Exec(sm.ctx,
		`INSERT INTO session (uuid, guest_uuid, new_guest, user_id, user_auth, events, hits, favorites, url_from, 
                     url_to, url_to_404, url_last, url_last_404, user_agent, date_stat, date_first, date_last, ip_first, ip_last, first_hit_uuid, 
                     last_hit_uuid, phpsessid, adv_uuid, adv_back, referer1, referer2, referer3, stop_list_uuid, country_id, first_site_uuid, last_site_uuid, city_id, sign, version) 
					VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		sessionUuid, session.GuestUuid, session.IsNewGuest, session.UserId, session.IsUserAuth, session.Events, session.Hits, session.Favorites, session.UrlFrom,
		session.FirstHitUuid, session.LastHitUuid, session.PhpSessionId, session.AdvUuid, session.AdvBack, session.Referer1, session.Referer2, session.Referer3,
		session.StopListUuid, session.CountryId, session.FirstSiteId, session.LastSiteId, session.CityId, session.Sign, session.Version,
	)
	if err != nil {
		return "", err
	}
	return sessionUuid, nil
}

func (sm Session) DeleteById(id int) error {
	err := sm.chClient.Exec(sm.ctx, "DELETE FROM session WHERE ID=?", id)
	if err != nil {
		return err
	}
	return nil
}

func (sm Session) FindSessionByGuestMd5(guestMd5 string) (entityjson.StatData, error) {
	var sessionData entityjson.StatData
	//err := sm.chClient.Exec(&sessionData,
	//	`SELECT *
	//           FROM session_data
	//           WHERE guest_md5=? and date_last > DATE_ADD(now(), INTERVAL-? SECOND)
	//           LIMIT 1`, guestMd5,
	//)
	//if err != nil {
	//	return entityjson.StatData{}, err
	//}
	return sessionData, nil
}

func (sm Session) Update(oldSession entitydb.Session, newSession entitydb.Session) error {

	err := sm.chClient.Exec(sm.ctx,
		`INSERT INTO session (uuid, guest_uuid, new_guest, user_id, user_auth, events,
                               hits, favorites, url_from, url_to, url_to_404, url_last, url_last_404, user_agent, date_stat, date_first, 
                               date_last, ip_first, ip_last, first_hit_uuid, last_hit_uuid, phpsessid, adv_uuid, adv_back,
                               referer1, referer2, referer3, stop_list_uuid, country_id, first_site_uuid, last_site_uuid, city_id,sign,version) 
								VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		oldSession.Uuid, oldSession.GuestUuid, oldSession.IsNewGuest, oldSession.UserId, oldSession.IsUserAuth, oldSession.Events, oldSession.Hits, oldSession.Favorites, oldSession.UrlFrom,
		oldSession.FirstHitUuid, oldSession.LastHitUuid, oldSession.PhpSessionId, oldSession.AdvUuid, oldSession.AdvBack, oldSession.Referer1, oldSession.Referer2, oldSession.Referer3,
		oldSession.StopListUuid, oldSession.CountryId, oldSession.FirstSiteId, oldSession.LastSiteId, oldSession.CityId, oldSession.Sign, oldSession.Version)

	if err != nil {
		return err
	}

	err = sm.chClient.Exec(sm.ctx,
		`INSERT INTO statistic.session (uuid, guest_uuid, new_guest, user_id, user_auth, events,
                               hits, favorites, url_from, url_to, url_to_404, url_last, url_last_404, user_agent, date_stat, date_first, 
                               date_last, ip_first, ip_last, first_hit_uuid, last_hit_uuid, phpsessid, adv_uuid, adv_back,
                               referer1, referer2, referer3, stop_list_uuid, country_id, first_site_uuid, last_site_uuid, city_id,sign,version) 
								VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, newSession.Uuid, newSession.GuestUuid, newSession.IsNewGuest, newSession.UserId, newSession.IsUserAuth, newSession.Events,
		newSession.Hits, newSession.Favorites, newSession.UrlFrom, newSession.FirstHitUuid, newSession.LastHitUuid, newSession.PhpSessionId, newSession.AdvUuid, newSession.AdvBack, newSession.Referer1, newSession.Referer2,
		newSession.Referer3, newSession.StopListUuid, newSession.CountryId, newSession.FirstSiteId, newSession.LastSiteId, newSession.CityId, 1, oldSession.Version+1)

	if err != nil {
		return err
	}

	return nil
}

func (sm Session) GetAttentiveness(dateStat, siteId string) {
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

func (sm Session) ExistsByPhpSession(session string) (int, error) {
	var count int
	row := sm.chClient.QueryRow(sm.ctx, `select count(uuid) as cnt from session where phpsessid=?`, session)

	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (sm Session) FindByPHPSessionId(phpSessionId string) (entitydb.Session, error) {
	var sessionDb entitydb.Session
	err := sm.chClient.QueryRow(sm.ctx, `select * from session where phpsessid=?`, phpSessionId).ScanStruct(&sessionDb)
	if err != nil {
		return entitydb.Session{}, err
	}
	return sessionDb, nil
}
