package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

type SessionService struct {
	ctx          context.Context
	chClient     driver.Conn
	sessionModel models.SessionModel
}

func NewSession(ctx context.Context, chClient driver.Conn) *SessionService {
	return &SessionService{
		ctx:      ctx,
		chClient: chClient,
	}
}

func (ss SessionService) AddSession(advBack, cityId string, countryId, stopListUuid, guestUuid uuid.UUID, isNewGuest bool, statData entityjson.StatData) error {
	//	$arFields = array(
	//		"GUEST_ID" => intval($_SESSION["SESS_GUEST_ID"]),
	//	    "NEW_GUEST" => "'" . $DB->ForSql($_SESSION["SESS_GUEST_NEW"]) . "'",
	//		"USER_ID" => intval($_SESSION["SESS_LAST_USER_ID"]),
	//	"USER_AUTH" => "'" . $DB->ForSql($IS_USER_AUTHORIZED) . "'",
	//		"URL_FROM" => "'" . $DB->ForSql($_SERVER["HTTP_REFERER"], 2000) . "'",
	//		"URL_TO" => "'" . $DB->ForSql($CURRENT_URI, 2000) . "'",
	//		"URL_TO_404" => "'" . $DB->ForSql($ERROR_404) . "'",
	//		"URL_LAST" => "'" . $DB->ForSql($CURRENT_URI, 2000) . "'",
	//		"URL_LAST_404" => "'" . $DB->ForSql($ERROR_404) . "'",
	//		"USER_AGENT" => "'" . $DB->ForSql($_SERVER["HTTP_USER_AGENT"], 500) . "'",

	//		"DATE_STAT" => $DB_now_date,
	//		"DATE_FIRST" => $DB_now,
	//		"DATE_LAST" => $DB_now,

	//		"IP_FIRST" => "'" . $DB->ForSql($_SERVER["REMOTE_ADDR"], 15) . "'",
	//		"IP_FIRST_NUMBER" => "'" . $DB->ForSql($REMOTE_ADDR_NUMBER) . "'",
	//		"IP_LAST" => "'" . $DB->ForSql($_SERVER["REMOTE_ADDR"], 15) . "'",
	//		"IP_LAST_NUMBER" => "'" . $DB->ForSql($REMOTE_ADDR_NUMBER) . "'",
	//		"PHPSESSID" => "'" . $DB->ForSql($sessionId, 255) . "'",
	//		"STOP_LIST_ID" => "'" . $DB->ForSql($STOP_LIST_ID) . "'",

	//		"COUNTRY_ID" => "'" . $DB->ForSql($_SESSION["SESS_COUNTRY_ID"], 2) . "'",
	//		"CITY_ID" => $_SESSION["SESS_CITY_ID"] > 0 ? intval($_SESSION["SESS_CITY_ID"]) : "null",
	//		"ADV_BACK" => "null",
	//		"FIRST_SITE_ID" => $sql_site,
	//		"LAST_SITE_ID" => $sql_site,
	//		"HITS" => 1,
	//);

	err := ss.sessionModel.AddSession(entitydb.SessionDb{
		GuestUuid:    guestUuid,
		IsNewGuest:   isNewGuest,
		UserId:       statData.UserId,
		IsUserAuth:   statData.IsUserAuth,
		UrlFrom:      statData.Referer,
		UrlTo:        statData.Url,
		UrlTo404:     statData.IsError404,
		UrlLast:      statData.Url,
		UrlLast404:   statData.IsError404,
		UserAgent:    statData.UserAgent,
		IpFirst:      statData.Ip,
		IpLast:       statData.Ip,
		PhpSessionId: statData.PHPSessionId,
		StopListUuid: stopListUuid,
		CountryId:    countryId,
		CityUuid:     cityId,
		AdvBack:      advBack,
		FirstSiteId:  statData.SiteId,
		LastSiteId:   statData.SiteId,
		Hits:         1,
	})

	if err != nil {
		return err
	}

	return nil
}

func (ss SessionService) IsExistsSession(phpSession string) bool {
	count, err := ss.sessionModel.ExistsByPhpSession(phpSession)
	if err != nil {
		return false
	}
	return count > 0
}

func (ss SessionService) UpdateSession(data entityjson.StatData) error {
	err := ss.sessionModel.Update(data)
	if err != nil {
		return err
	}
	return nil
}
