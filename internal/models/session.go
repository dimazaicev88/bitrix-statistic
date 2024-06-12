package models

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/storage"
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
	_, err := sm.storage.DB().MustExec(`INSERT INTO session (GuestId, NewGuest, USER_ID, USER_AUTH, C_EVENTS, HITS, FAVORITES, URL_FROM, URL_TO, URL_TO_404, URL_LAST,
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
	sm.storage.DB().MustExec("DELETE FROM session WHERE ID=?", id)
}

func (sm SessionModel) FindSessionByGuestMd5(guestMd5 string) (entity.SessionData, error) {
	var sessionData entity.SessionData
	err := sm.storage.DB().Get(&sessionData,
		`SELECT * 
               FROM session_data
               WHERE guest_md5=? and date_last > DATE_ADD(now(), INTERVAL-? SECOND) 
               LIMIT 1`, guestMd5,
	)
	if err != nil {
		return entity.SessionData{}, err
	}
	return sessionData, nil
}

func (sm SessionModel) UpdateHits(get string, authorized string, agent string, ip string) error {
	//addr := net.ParseIP(ip)
	//ipNum := addr.To4().String()
	//
	//sm.storage.DB().MustExec("UPDATE session SET  WHERE id=? ")

	return nil
}

func (sm SessionModel) GetAttentiveness(dateStat, siteId string) {
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
	//S.DATE_STAT = cast(".$DB->CharToDateFunction($DATE_STAT, "SHORT")." as date)
	//$str
	//";
	//
	//$rs = $DB->Query($strSql, false, $err_mess.__LINE__);
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
