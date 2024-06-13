package services

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/storage"
	"bitrix-statistic/utils"
	"golang.org/x/exp/slices"
	_ "net/netip"
	"regexp"
	"strconv"
	"strings"
)

type Statistic struct {
	statisticModel models.StatisticModel
	optionModel    *models.OptionModel
	searcherModel  *models.SearcherModel
	sessionModel   models.SessionModel
	cityModel      models.CityModel
	advModel       *models.AdvModel
	guestModel     *models.GuestModel
}

func NewStatistic(storage *storage.MysqlStorage) Statistic {
	return Statistic{
		statisticModel: models.NewStatisticModel(storage),
		optionModel:    models.NewOptionModel(storage),
		sessionModel:   models.NewSessionModel(storage),
		searcherModel:  models.NewSearcherModel(storage),
		cityModel:      models.NewCityModel(storage),
		advModel:       models.NewAdvModel(storage, models.NewOptionModel(storage)),
		guestModel:     models.NewGuestModel(storage),
	}
}

func (s Statistic) checkSkip(userGroups []int, remoteAddr string) (error, bool) {
	skipMode := s.optionModel.Get("SKIP_STATISTIC_WHAT")

	isSkip := false
	switch skipMode {
	case "none":
		break
	case "both":
	case "groups":
		arSkipGroups := strings.Split(",", s.optionModel.Get("SKIP_STATISTIC_GROUPS"))
		for _, group := range arSkipGroups {
			groupId, err := strconv.Atoi(group)
			if err != nil {
				return err, false
			}
			if slices.Contains(userGroups, groupId) {
				isSkip = true
			}
		}
	case "ranges":
		if skipMode == "both" && isSkip == true {
			break
		}
		isSkip = true
		var re = regexp.MustCompile(`/^.*?(\d+)\.(\d+)\.(\d+)\.(\d+)[\s-]*/`)
		arIPAAddress := re.FindStringSubmatch(remoteAddr)
		if len(re.FindStringSubmatch(remoteAddr)) > 0 {
			arSkipIPRanges := strings.Split("\n", s.optionModel.Get("SKIP_STATISTIC_IP_RANGES"))
			for _, skipRange := range arSkipIPRanges {
				var re = regexp.MustCompile(`/^.*?(\d+)\.(\d+)\.(\d+)\.(\d+)[\s-]*(\d+)\.(\d+)\.(\d+)\.(\d+)/`)
				matchSkipRange := re.FindStringSubmatch(skipRange)
				if len(matchSkipRange) > 0 {
					if utils.StrToInt(arIPAAddress[1]) >= int(skipRange[1]) &&
						utils.StrToInt(arIPAAddress[2]) >= int(skipRange[2]) &&
						utils.StrToInt(arIPAAddress[3]) >= int(skipRange[3]) &&
						utils.StrToInt(arIPAAddress[4]) >= int(skipRange[4]) &&
						utils.StrToInt(arIPAAddress[1]) <= int(skipRange[5]) &&
						utils.StrToInt(arIPAAddress[2]) <= int(skipRange[6]) &&
						utils.StrToInt(arIPAAddress[3]) <= int(skipRange[7]) &&
						utils.StrToInt(arIPAAddress[4]) <= int(skipRange[8]) {
						isSkip = true
						break
					}
				}
			}
		}
		break
	}
	return nil, isSkip
}

func (s Statistic) Add(statData entity.StatData) error {
	existsGuest, err := s.guestModel.ExistsGuestByToken(statData.CookieToken)
	if err != nil {
		return err
	}

	//Guest не найден
	if existsGuest == false {
		err := s.guestModel.AddGuest(statData)
		if err != nil {
			return err
		}
	}
	//	advNa := s.optionModel.Get("ADV_NA")
	//	__SetReferer("referer1", "REFERER1_SYN")
	//	__SetReferer("referer2", "REFERER2_SYN")
	//	__SetReferer("referer3", "REFERER3_SYN")
	//
	//	saveHits := s.optionModel.Get("SAVE_HITS")
	//	saveVisits := s.optionModel.Get("SAVE_VISITS")
	//	saveReferers := s.optionModel.Get("SAVE_REFERERS")
	//	savePathData := s.optionModel.Get("SAVE_PATH_DATA")
	//
	//	//$stmp = time()
	//	//$hour = date("G", $stmp)    // 0..23
	//	//$weekday = date("w", $stmp) // 0..6
	//	//if ($weekday == 0) $weekday = 7
	//	//$month = date("n", $stmp) // 1..12
	//	var FAVORITES string
	//	if stopSaveStatistic == false || stop == false {
	//		if phpSession.KeyExists("SESS_ADD_TO_FAVORITES") && phpSession.Get("SESS_ADD_TO_FAVORITES") == "Y" {
	//			FAVORITES = "Y"
	//			phpSession.Set("SESS_ADD_TO_FAVORITES", "")
	//		} else
	//		{
	//			FAVORITES = "N"
	//		}
	//
	//		//$Error404 = (defined("ERROR_404") && Error404 == "Y") ? "Y" : "N"
	//		//$DB_now = $DB- > GetNowFunction()  // save function for use in sql
	//		//$DB_now_date = $DB- > GetNowDate() // save function for use in sql
	//		//stopListId = intval($StopListId)
	//		//if ($Error404 == "Y") init_get_params($APPLICATION- > GetCurUri())
	//		//
	//		//$IsUserAuthorized = (isset($_SESSION["SESS_LAST_USER_ID"]) && intval($_SESSION["SESS_LAST_USER_ID"])>0 && is_object($USER) && $USER- > IsAuthorized()) ? "Y" : "N"
	//
	//		phpSession.Set("SESS_ADV_ID", "")
	//		phpSession.Set("SESS_SEARCHER_ID", "")
	//		phpSession.Set("SESS_SEARCHER_NAME", "")
	//		phpSession.Set("SESS_SEARCHER_CHECK_ACTIVITY", "")
	//		phpSession.Set("SESS_SEARCHER_SAVE_STATISTIC", "")
	//		phpSession.Set("SESS_SEARCHER_HIT_KEEP_DAYS", "")
	//		phpSession.Set("SESS_LAST_PROTOCOL", "")
	//		phpSession.Set("SESS_LAST_URI", "")
	//		phpSession.Set("SESS_LAST_HOST", "")
	//		phpSession.Set("SESS_LAST_PAGE", "")
	//		phpSession.Set("SESS_LAST_DIR", "")
	//		phpSession.Set("SESS_HTTP_REFERER", "")
	//		phpSession.Set("SESS_COUNTRY_ID", "")
	//		phpSession.Set("SESS_CITY_ID", "")
	//		phpSession.Set("SESS_SESSION_ID", "")
	//		phpSession.Set("SESS_REFERER_ID", "")
	//		phpSession.Set("FROM_SEARCHER_ID", "")
	//		phpSession.Set("SESS_FROM_SEARCHERS", "")
	//		phpSession.Set("SESS_REQUEST_URI_CHANGE", "")
	//		phpSession.Set("SESS_LAST_DIR_ID", "")
	//		phpSession.Set("SESS_LAST_PAGE_ID", "")
	//		phpSession.Set("SESS_GRABBER_STOP_TIME", "")
	//		phpSession.Set("SESS_GRABBER_DEFENCE_STACK", "")
	//		phpSession.Set("ACTIVITY_EXCEEDING_NOTIFIED", "")
	//
	//		// SESSION_DATA_ID will be false when there is no sessions stored
	//		// true when phpSession was not found in database
	//		// and an integer when was found and populated to $SESSION array
	//		//$SESSION_DATA_ID =
	//		s.RestoreSession(phpServer, cookieGuestId, phpSession, sessionGcMaxLifeTime)
	//
	//		// Let's check activity limit
	//		//TODO обдумать
	//		blockActibity := false
	//		//$BLOCK_ACTIVITY = CStatistics::BlockVisitorActivity()
	//
	//		// Activity under the limit
	//		if blockActibity {
	//			//Check if searcher was not deleted from searchers list
	//			if phpSession.KeyExists("SESS_SEARCHER_ID") && phpSession.GetAsInt("SESS_SEARCHER_ID") > 0 {
	//				if s.searcherModel.ExistById(phpSession.GetAsInt("SESS_SEARCHER_ID")) == false {
	//					phpSession.Delete("SESS_SEARCHER_ID")
	//				}
	//			}
	//
	//			// We did not check for searcher
	//			if !phpSession.KeyExists("SESS_SEARCHER_ID") || phpSession.Get("SESS_SEARCHER_ID") == "" {
	//				searchers, err := s.searcherModel.FindSearcherByUserAgent(userAgent)
	//				if err != nil {
	//					return
	//				}
	//				//TODO перебор цикла
	//				for _, searcher := range searchers {
	//					phpSession.Set("SESS_SEARCHER_ID", strconv.Itoa(searcher.Id))
	//					phpSession.Set("SESS_SEARCHER_NAME", searcher.Name)
	//					phpSession.Set("SESS_SEARCHER_CHECK_ACTIVITY", searcher.CheckActivity)
	//					phpSession.Set("SESS_SEARCHER_SAVE_STATISTIC", searcher.SaveStatistic)
	//					phpSession.Set("SESS_SEARCHER_HIT_KEEP_DAYS", strconv.Itoa(searcher.HitKeepDays))
	//				}
	//				phpSession.Set("SESS_SEARCHER_ID", strconv.Itoa(0))
	//			}
	//
	//		}
	//
	//		/************************************************
	//				Searcher section
	//		************************************************/
	//
	//		// searcher detected
	//		if phpSession.GetAsInt("SESS_SEARCHER_ID") > 0 {
	//			searchers, err := s.searcherModel.ExistByIdAndCurrentDate(phpSession.GetAsInt("SESS_SEARCHER_ID"))
	//			if err != nil {
	//				return
	//			}
	//			if len(searchers) > 0 {
	//				s.searcherModel.UpdateSearcherDay(phpSession.GetAsInt("SESS_SEARCHER_ID"))
	//			} else {
	//				s.searcherModel.AddSearcherDay(phpSession.GetAsInt("SESS_SEARCHER_ID"))
	//			}
	//
	//			// save indexed page if neccessary
	//			if phpSession.Get("SESS_SEARCHER_SAVE_STATISTIC") == "Y" {
	//				s.searcherModel.AddSearcherHit(phpSession.GetAsInt("SESS_SEARCHER_ID"), fullRequestUri, error404, ip, userAgent, phpSession.Get("SESS_SEARCHER_HIT_KEEP_DAYS"), siteId)
	//				//TODO ?
	//				//if error404 == "N" {
	//				//CStatistics::Set404("b_stat_searcher_hit", "ID = ".intval($id), array("URL_404" = > "Y"))
	//				//}
	//			}
	//		} else // it is not searcher
	//		{
	//
	//			/************************************************
	//					Visitor section
	//			************************************************/
	//
	//			/************************************************
	//				Variables which describe current page
	//			************************************************/
	//
	//			//$CURRENT_DIR = __GetCurrentDir()
	//			//$CURRENT_PAGE = __GetCurrentPage()
	//			//
	//			//$CURRENT_PROTOCOL = (CMain::IsHTTPS()) ? "https://": "http://" // protocol
	//			//$CURRENT_PORT = $_SERVER["SERVER_PORT"]                        // port
	//			//$CURRENT_HOST = $_SERVER["HTTP_HOST"]                          // domain
	//			//$CURRENT_PAGE = __GetFullRequestUri($CURRENT_PAGE)             // w/o parameters
	//			//$CURRENT_URI = __GetFullRequestUri()                           // with params
	//			//$CURRENT_DIR = __GetFullRequestUri($CURRENT_DIR)               // catalog
	//
	//			/************************************************
	//					Country detection
	//			************************************************/
	//
	//			if phpSession.KeyExists("SESS_COUNTRY_ID") == false || phpSession.Get("SESS_COUNTRY_ID") == "" {
	//				phpSession.Set("SESS_COUNTRY_ID", s.cityModel.GetCountryCode())
	//				phpSession.Set("SESS_CITY_ID", s.cityModel.GetCityID())
	//			}
	//
	//			/************************************************
	//					IP => number
	//			************************************************/
	//
	//			//$REMOTE_ADDR_NUMBER = ip2number($_SERVER["REMOTE_ADDR"])
	//
	//			/************************************************
	//					Advertising campaign
	//			************************************************/
	//			err := s.advModel.SetAdv(phpSession, fullRequestUri, openstat, referringSite)
	//			if err != nil {
	//				return
	//			}
	//			//CStatistics::Set_Adv()
	//
	//			/************************************************
	//					Guest ID detection
	//			************************************************/
	//			var ref1, ref2 string
	//			ref1, ref2, err = s.statisticModel.SetGuest(phpSession, siteId, referringSite, fullRequestUri, error404, cookieGuestId, cookieLastVisit, cookieAdvId)
	//			if err != nil {
	//				return
	//			}
	//
	//			// Setup default advertising campaign
	//			if advNa == "Y" && phpSession.GetAsInt("SESS_ADV_ID") <= 0 && phpSession.GetAsInt("SESS_LAST_ADV_ID") <= 0 {
	//				phpSession.Set("referer1", s.optionModel.Get("AVD_NA_REFERER1"))
	//				phpSession.Set("referer2", s.optionModel.Get("AVD_NA_REFERER2"))
	//				err := s.advModel.SetAdv(phpSession, fullRequestUri, openstat, referringSite)
	//				if err != nil {
	//					return
	//				}
	//				ref1, ref2, err = s.statisticModel.SetGuest(phpSession, siteId, referringSite, fullRequestUri, error404, cookieGuestId, cookieLastVisit, cookieAdvId)
	//				if err != nil {
	//					return
	//				}
	//			}
	//
	//			/************************************************
	//					Session section
	//			************************************************/
	//
	//			//$_SESSION["SESS_SESSION_ID"] = intval($_SESSION["SESS_SESSION_ID"] ?? 0)
	//
	//			//phpSession already exists
	//			if phpSession.GetAsInt("SESS_SESSION_ID") > 0 {
	//				SESSION_NEW := "N"
	//				// update
	//				$arFields = Array(
	//					"USER_ID" = > intval($_SESSION["SESS_LAST_USER_ID"]),
	//				"USER_AUTH" = > "'".$IS_USER_AUTHORIZED.
	//				"'",
	//					"USER_AGENT" = > "'".$DB- > ForSql($_SERVER["HTTP_USER_AGENT"], 500)."'",
	//					"DATE_LAST" = > $DB_now,
	//					"IP_LAST" = > "'".$DB- > ForSql($_SERVER["REMOTE_ADDR"], 15)."'",
	//					"IP_LAST_NUMBER" = > $REMOTE_ADDR_NUMBER,
	//					"HITS" => "HITS + 1",
	//			)
	//
	//				s.sessionModel.UpdateHits(phpSession.Get("SESS_LAST_USER_ID"), isUserAuthorized, userAgent, ip)
	//
	//				$rows = $DB- > Update("b_stat_session", $arFields, "WHERE ID='".$_SESSION["SESS_SESSION_ID"].
	//				"'", "File: ".__FILE__.
	//				"<br>Line: ".__LINE__)
	//				// was cleaned up
	//				if intval($rows)<=0)
	//				{
	//				// store as new
	//				phpSession.Set("SESS_SESSION_ID", "0")
	//				if (ADV_NA=="Y" && intval($_SESSION["SESS_ADV_ID"])<=0 && intval($_SESSION["SESS_LAST_ADV_ID"])<=0)
	//				{
	//				$_SESSION["referer1"] = COption::GetOptionString("statistic", "AVD_NA_REFERER1");
	//				$_SESSION["referer2"] = COption::GetOptionString("statistic", "AVD_NA_REFERER2");
	//				}
	//				CStatistics::Set_Adv();
	//				$arGuest = CStatistics::Set_Guest();
	//				}
	//			}
	//
	//			// it is new phpSession
	//			if phpSession.GetAsInt("SESS_SESSION_ID") <= 0 {
	//				SESSION_NEW := "Y"
	//
	//				// save phpSession data
	//				$arFields = Array(
	//					"GUEST_ID" = > intval($_SESSION["SESS_GUEST_ID"]),
	//				"NEW_GUEST" = > "'".$DB- > ForSql($_SESSION["SESS_GUEST_NEW"])."'",
	//					"USER_ID" = > intval($_SESSION["SESS_LAST_USER_ID"]),
	//				"USER_AUTH" = > "'".$DB- > ForSql($IS_USER_AUTHORIZED)."'",
	//					"URL_FROM" = > "'".$DB- > ForSql($_SERVER["HTTP_REFERER"], 2000)."'",
	//					"URL_TO" => "'".$DB- > ForSql($CURRENT_URI, 2000)."'",
	//					"URL_TO_404" = > "'".$DB- > ForSql($ERROR_404)."'",
	//					"URL_LAST" = > "'".$DB- > ForSql($CURRENT_URI, 2000)."'",
	//					"URL_LAST_404" = > "'".$DB- > ForSql($ERROR_404)."'",
	//					"USER_AGENT" = > "'".$DB- > ForSql($_SERVER["HTTP_USER_AGENT"], 500)."'",
	//					"DATE_STAT" = > $DB_now_date,
	//					"DATE_FIRST" => $DB_now,
	//					"DATE_LAST" = > $DB_now,
	//					"IP_FIRST" = > "'".$DB- > ForSql($_SERVER["REMOTE_ADDR"], 15)."'",
	//					"IP_FIRST_NUMBER" = > "'".$DB- > ForSql($REMOTE_ADDR_NUMBER)."'",
	//					"IP_LAST" = > "'".$DB- > ForSql($_SERVER["REMOTE_ADDR"], 15)."'",
	//					"IP_LAST_NUMBER" = > "'".$DB- > ForSql($REMOTE_ADDR_NUMBER)."'",
	//					"PHPSESSID" => "'".$DB- > ForSql($sessionId, 255)."'",
	//					"STOP_LIST_ID" = > "'".$DB- > ForSql($STOP_LIST_ID)."'",
	//					"COUNTRY_ID" = > "'".$DB- > ForSql($_SESSION["SESS_COUNTRY_ID"], 2)."'",
	//					"CITY_ID" = > $_SESSION["SESS_CITY_ID"] > 0? intval($_SESSION["SESS_CITY_ID"]): "null",
	//					"ADV_BACK" => "null",
	//					"FIRST_SITE_ID" = > $sql_site,
	//					"LAST_SITE_ID" = > $sql_site,
	//					"HITS" = > 1,
	//			)
	//
	//				// campaign?
	//				if phpSession.GetAsInt("SESS_ADV_ID") > 0 {
	//					$arFields["ADV_ID"] = intval($_SESSION["SESS_ADV_ID"])
	//					$arFields["ADV_BACK"] = "'N'"
	//					$arFields["REFERER1"] = "'".$DB- > ForSql($_SESSION["referer1"], 255)."'"
	//					$arFields["REFERER2"] = "'".$DB- > ForSql($_SESSION["referer2"], 255)."'"
	//					$arFields["REFERER3"] = "'".$DB- > ForSql($_SESSION["referer3"], 255)."'"
	//				} else if phpSession.GetAsInt("SESS_LAST_ADV_ID") > 0 { // comeback?
	//
	//					$arFields["ADV_ID"] = intval($_SESSION["SESS_LAST_ADV_ID"])
	//					$arFields["ADV_BACK"] = "'Y'"
	//					$arFields["REFERER1"] = "'".$DB- > ForSql($arGuest["last_referer1"], 255)."'"
	//					$arFields["REFERER2"] = "'".$DB- > ForSql($arGuest["last_referer2"], 255)."'"
	//				}
	//
	//				// look for the same IP?
	//				day_host_counter := 1
	//				day_host_counter_site := $SITE_ID < > ''? 1: 0
	//				strSql := `
	//				SELECT
	//				S.FIRST_SITE_ID
	//				FROM
	//				b_stat_session
	//				S
	//				WHERE
	//				S.IP_FIRST_NUMBER = ".$REMOTE_ADDR_NUMBER."
	//				AND
	//				S.DATE_STAT = ".$DB_now_date."
	//				`
	//				$e = $DB- > Query($strSql, false, "File: ".__FILE__.
	//				"<br>Line: ".__LINE__)
	//				while($er = $e- > Fetch())
	//				{
	//				$day_host_counter = 0;
	//				if ($SITE_ID==$er["FIRST_SITE_ID"])
	//				{
	//				$day_host_counter_site = 0;
	//				break;
	//				}
	//				}
	//
	//				$_SESSION["SESS_SESSION_ID"] = intval($DB- > Insert("b_stat_session", $arFields, "File: ".__FILE__.
	//				"<br>Line: ".__LINE__))
	//
	//				if ($ERROR_404 == "N")
	//				{
	//				CStatistics::Set404("b_stat_session", "ID = ".$_SESSION["SESS_SESSION_ID"], array("URL_TO_404" = > "Y", "URL_LAST_404" = > "Y"));
	//				}
	//
	//				$day_guest_counter = 0
	//				$new_guest_counter = 0
	//				// new guest
	//				if ($_SESSION["SESS_GUEST_NEW"] == "Y")
	//				{
	//				// update day statistic
	//				$day_guest_counter = 1;
	//				$new_guest_counter = 1;
	//				}
	//				else // guest was here
	//				{
	//				// first hit for today
	//				if ($_SESSION["SESS_LAST"]!="Y")
	//				{
	//				// update day statistic
	//				$day_guest_counter = 1;
	//				$_SESSION["SESS_LAST"] = "Y";
	//				}
	//				}
	//
	//				// update day counter
	//				$arFields = Array(
	//					"SESSIONS" = > 1,
	//					"C_HOSTS" = > intval($day_host_counter),
	//				"GUESTS" = > intval($day_guest_counter),
	//				"NEW_GUESTS" = > intval($new_guest_counter),
	//				"SESSION" => 1,
	//					"HOST" = > intval($day_host_counter),
	//				"GUEST" = > intval($day_guest_counter),
	//				"NEW_GUEST" = > intval($new_guest_counter),
	//			)
	//				// when current day is already exists
	//				// we have to update it
	//				$rows = CTraffic::IncParam($arFields)
	//				if ($rows != = false && $rows <= 0)
	//				// otherwise
	//				{
	//				// add new one
	//				CStatistics::SetNewDay(
	//				1, // HOSTS
	//				0, // TOTAL_HOSTS (now ignored)
	//				1,                          // SESSIONS
	//				0,                          // HITS
	//				intval($new_guest_counter), // NEW_GUESTS
	//				1                // GUESTS
	//				);
	//
	//				// and update it
	//				CTraffic::IncParam(
	//				array(
	//				"SESSION" = > 1,
	//				"HOST" = > 1,
	//				"GUEST" = > 1,
	//				"NEW_GUEST" = > intval($new_guest_counter),
	//				)
	//				);
	//				}
	//
	//				// site is not defined
	//				if ($SITE_ID < > '')
	//				{
	//				// ��������� ������� "�� ����" ��� �������� �����
	//				$arFields = Array(
	//				"SESSIONS" = > 1,
	//				"C_HOSTS" = > intval($day_host_counter_site),
	//				"SESSION"    = > 1,
	//				"HOST" = > intval($day_host_counter_site),
	//				);
	//				// ������� �������� �������� ��� �������� ���
	//				$rows = CTraffic::IncParam(array(), $arFields, $SITE_ID);
	//				// ���� �������� ��� ��� ����� � ���� ��� ��� ��
	//				if ($rows!= =false && intval($rows)<=0)
	//				{
	//				// ��������� ���
	//				CStatistics::SetNewDayForSite(
	//				$SITE_ID,
	//				1,    // HOSTS
	//				0,    // TOTAL_HOSTS  (now ignored)
	//				1     // SESSIONS
	//				);
	//
	//				// ������� �������� �������� ��� �������� ���
	//				CTraffic::IncParam(
	//				array(),
	//				array(
	//				"SESSION" = > 1,
	//				"HOST" = > 1,
	//				),
	//				$SITE_ID
	//				);
	//				}
	//				}
	//
	//				// ���� ������ ���������� ��
	//				if ($_SESSION["SESS_COUNTRY_ID"] < > '')
	//				{
	//				$arFields = Array(
	//				"SESSIONS" = > 1,
	//				"NEW_GUESTS" = > $new_guest_counter,
	//				);
	//				CStatistics::UpdateCountry($_SESSION["SESS_COUNTRY_ID"], $arFields);
	//				}
	//
	//				if ($_SESSION["SESS_CITY_ID"] > 0)
	//				{
	//				$arFields = Array(
	//				"SESSIONS" = > 1,
	//				"NEW_GUESTS" = > $new_guest_counter,
	//				);
	//				CStatistics::UpdateCity($_SESSION["SESS_CITY_ID"], $arFields);
	//				}
	//
	//				// ��������� �����
	//				$arFields = Array(
	//					"SESSIONS" = > "SESSIONS + 1",
	//					"LAST_SESSION_ID" = > $_SESSION["SESS_SESSION_ID"],
	//					"LAST_USER_AGENT" = > "'".$DB- > ForSql($_SERVER["HTTP_USER_AGENT"], 500)."'",
	//					"LAST_COUNTRY_ID" = > "'".$DB- > ForSql($_SESSION["SESS_COUNTRY_ID"], 2)."'",
	//					"LAST_CITY_ID" = > $_SESSION["SESS_CITY_ID"] > 0? intval($_SESSION["SESS_CITY_ID"]): "null",
	//			)
	//				//
	//				if ($obCity)
	//				{
	//				$arFields["LAST_CITY_INFO"] = "'".$obCity->ForSQL()."'";
	//				}
	//				// ���� ��� ������ ����� �� ��������� �������� ��
	//				if intval($_SESSION["SESS_ADV_ID"])>0)
	//				{
	//				// ��������� ��������� �������� ���������� ������ �����
	//				$arFields["LAST_ADV_ID"] = intval($_SESSION["SESS_ADV_ID"]);
	//				$arFields["LAST_ADV_BACK"] = "'N'";
	//				$arFields["LAST_REFERER1"] = "'".$DB->ForSql($_SESSION["referer1"], 255)."'";
	//				$arFields["LAST_REFERER2"] = "'".$DB->ForSql($_SESSION["referer2"], 255)."'";
	//				$arFields["LAST_REFERER3"] = "'".$DB->ForSql($_SESSION["referer3"], 255)."'";
	//				}
	//				elseif(intval($_SESSION["SESS_LAST_ADV_ID"])>0) // ����� ���� ��� ������� ��
	//				{
	//				// ������� ���� �������� �� ��������� ������ �����
	//				$arFields["LAST_ADV_BACK"] = "'Y'";
	//				$arFields["LAST_REFERER1"] = "'".$DB->ForSql($arGuest["last_referer1"], 255)."'";
	//				$arFields["LAST_REFERER2"] = "'".$DB->ForSql($arGuest["last_referer2"], 255)."'";
	//				}
	//
	//				if ($_SESSION["SESS_GUEST_NEW"] == "Y")
	//				$arFields["FIRST_SESSION_ID"] = $_SESSION["SESS_SESSION_ID"]
	//				$rows = $DB- > Update("b_stat_guest", $arFields, "WHERE ID=".intval($_SESSION["SESS_GUEST_ID"]), "File: ".__FILE__.
	//				"<br>Line: ".__LINE__, false, false, false)
	//
	//				// ��������� ��������� ��������
	//				if intval($_SESSION["SESS_ADV_ID"])>0 || intval($_SESSION["SESS_LAST_ADV_ID"])>0)
	//				{
	//				CStatistics::Update_Adv();
	//				}
	//
	//				/************************************************
	//						Referring sites
	//				************************************************/
	//				if (
	//				$SAVE_REFERERS != "N"
	//				&& __GetReferringSite($PROT, $SN, $SN_WithoutPort, $PAGE_FROM)
	//				&& $SN < > ''
	//				&& $SN != $_SERVER["HTTP_HOST"]
	//				)
	//				{
	//				$REFERER_LIST_ID = CStatistics::GetRefererListID($PROT, $SN, $PAGE_FROM, $CURRENT_URI, $ERROR_404, $sql_site);
	//
	//				/************************************************
	//						Search phrases
	//				************************************************/
	//
	//				if (mb_substr($SN, 0, 4) == "www.")
	//				$sql = "('".$DB->ForSql(mb_substr($SN, 4), 255)."' like P.DOMAIN or '".$DB->ForSql($SN, 255)."' like P.DOMAIN)";
	//				else
	//				$sql = "'".$DB->ForSql($SN, 255)."' like P.DOMAIN";
	//				$strSql = "
	//				SELECT
	//				S.ID,
	//				S.NAME,
	//				P.DOMAIN,
	//				P.VARIABLE,
	//				P.CHAR_SET
	//				FROM
	//				b_stat_searcher S,
	//				b_stat_searcher_params P
	//				WHERE
	//				S.ACTIVE= 'Y'
	//				and    P.SEARCHER_ID = S.ID
	//				and    ".$sql."
	//				";
	//				$q = $DB->Query($strSql, false, "File: ".__FILE__."<br>Line: ".__LINE__);
	//				if ($qr = $q->Fetch())
	//				{
	//				$_SESSION["FROM_SEARCHER_ID"] = $qr["ID"];
	//				$FROM_SEARCHER_NAME = $qr["NAME"];
	//				$FROM_SEARCHER_PHRASE = "";
	//				if ($qr["VARIABLE"] <> '')
	//				{
	//				$page = mb_substr($PAGE_FROM, mb_strpos($PAGE_FROM, "?") + 1);
	//				$bIsUTF8 = is_utf8_url($page);
	//				parse_str($page, $arr);
	//				$arrVar = explode(",", $qr["VARIABLE"]);
	//				foreach ($arrVar as $var )
	//				{
	//				$var = trim($var );
	//				$phrase = $arr[$var ];
	//
	//				if($bIsUTF8)
	//				{
	//				$phrase_temp = trim($APPLICATION->ConvertCharset($phrase, "utf-8", LANG_CHARSET));
	//				if ($phrase_temp <> '')
	//				{
	//				$phrase = $phrase_temp;
	//				}
	//				}
	//				elseif($qr["CHAR_SET"] <> '')
	//				{
	//				$phrase_temp = trim($APPLICATION->ConvertCharset($phrase, $qr["CHAR_SET"], LANG_CHARSET));
	//				if ($phrase_temp <> '')
	//				{
	//				$phrase = $phrase_temp;
	//				}
	//				}
	//
	//				$phrase = trim($phrase);
	//				if ($phrase <> '')
	//				{
	//				$FROM_SEARCHER_PHRASE.= ($FROM_SEARCHER_PHRASE <> '')? " / ".$phrase: $phrase;
	//				}
	//				}
	//				}
	//				//echo "FROM_SEARCHER_PHRASE = ".$FROM_SEARCHER_PHRASE."<br>\n";
	//				// ���� �������� ��������� �����, �� ������� �� � ����
	//				if ($FROM_SEARCHER_PHRASE <> '')
	//				{
	//				$arFields = Array(
	//				"DATE_HIT" = > $DB_now,
	//				"SEARCHER_ID" = > intval($_SESSION["FROM_SEARCHER_ID"]),
	//				"REFERER_ID" = > $REFERER_LIST_ID,
	//				"PHRASE" = > "'".$DB->ForSql($FROM_SEARCHER_PHRASE, 255)."'",
	//				"URL_FROM" => "'".$DB->ForSql($PROT.$SN.$PAGE_FROM, 2000)."'",
	//				"URL_TO" = > "'".$DB->ForSql($CURRENT_URI, 2000)."'",
	//				"URL_TO_404" = > "'".$ERROR_404."'",
	//				"SESSION_ID" = > $_SESSION["SESS_SESSION_ID"],
	//				"SITE_ID" = > $sql_site,
	//				);
	//				$id = $DB->Insert("b_stat_phrase_list", $arFields, "File: ".__FILE__."<br>Line: ".__LINE__);
	//				if ($ERROR_404=="N")
	//				{
	//				CStatistics::Set404("b_stat_phrase_list", "ID = ".intval($id), array("URL_TO_404" = > "Y"));
	//				}
	//
	//				// �������� ��������� ����� � ������
	//				$_SESSION["SESS_SEARCH_PHRASE"] = $FROM_SEARCHER_PHRASE;
	//
	//				// �������� ������� ���� � ��������� �������
	//				$_SESSION["SESS_FROM_SEARCHERS"][] = $_SESSION["FROM_SEARCHER_ID"];
	//				$arFields = Array("PHRASES" = > "PHRASES + 1");
	//				$rows = $DB->Update("b_stat_searcher", $arFields, "WHERE ID=".intval($_SESSION["FROM_SEARCHER_ID"]), "File: ".__FILE__."<br>Line: ".__LINE__, false,false, false);
	//
	//				}
	//				}
	//				}
	//			}
	//
	//			/************************************************
	//						Hits
	//			************************************************/
	//
	//			if ($_SESSION["SESS_SESSION_ID"] > 0)
	//			{
	//			if ($SAVE_HITS!="N")
	//			{
	//			// ��������� ���
	//			$arFields = Array(
	//			"SESSION_ID" = > $_SESSION["SESS_SESSION_ID"],
	//			"DATE_HIT" = > $DB_now,
	//			"GUEST_ID" = > intval($_SESSION["SESS_GUEST_ID"]),
	//			"NEW_GUEST" = > "'".$DB->ForSql($_SESSION["SESS_GUEST_NEW"])."'",
	//			"USER_ID" = > intval($_SESSION["SESS_LAST_USER_ID"]),
	//			"USER_AUTH" = > "'".$IS_USER_AUTHORIZED."'",
	//			"URL" = > "'".$DB->ForSql($CURRENT_URI, 2000)."'",
	//			"URL_404" = > "'".$ERROR_404."'",
	//			"URL_FROM" = > "'".$DB->ForSql(isset($_SERVER["HTTP_REFERER"])? $_SERVER["HTTP_REFERER"]: "", 2000)."'",
	//			"IP" = > "'".$DB->ForSql($_SERVER["REMOTE_ADDR"], 15)."'",
	//			"METHOD" => "'".$DB->ForSql($_SERVER["REQUEST_METHOD"], 10)."'",
	//			"COOKIES" = > "'".$DB->ForSql(GetCookieString(), 2000)."'",
	//			"USER_AGENT" = > "'".$DB->ForSql($_SERVER["HTTP_USER_AGENT"], 500)."'",
	//			"STOP_LIST_ID" = > "'".$STOP_LIST_ID."'",
	//			"COUNTRY_ID" = > "'".$DB->ForSql($_SESSION["SESS_COUNTRY_ID"], 2)."'",
	//			"CITY_ID" = > $_SESSION["SESS_CITY_ID"] > 0? intval($_SESSION["SESS_CITY_ID"]): "null",
	//			"SITE_ID" = > $sql_site,
	//			);
	//			self::$HIT_ID = intval($DB->Insert("b_stat_hit", $arFields, "File: ".__FILE__."<br>Line: ".__LINE__));
	//			if ($ERROR_404=="N")
	//			{
	//			CStatistics::Set404("b_stat_hit", "ID = ".self::$HIT_ID, array("URL_404" = > "Y"));
	//			}
	//			}
	//
	//			// ���� ����� �� ������ ���� ������� � �������� � �� ����� ��� �� �������� ��
	//			$favorites_counter = 0;
	//			if ($FAVORITES=="Y" && $_SESSION["SESS_GUEST_FAVORITES"]=="N")
	//			{
	//			$ALLOW_ADV_FAVORITES = "Y";
	//			$_SESSION["SESS_GUEST_FAVORITES"] = "Y";
	//			$favorites_counter = 1;
	//			}
	//			// ��������� ������� "�� ����"
	//			$arFields = Array(
	//			"HITS" = > 1,
	//			"FAVORITES" = > $favorites_counter,
	//			"HIT" = > 1,
	//			"FAVORITE" = > $favorites_counter,
	//			);
	//			// ���� ������� ���� ���� � ���� ��
	//			// ������� �������� �������� ��� �������� ���
	//			$rows = CTraffic::IncParam($arFields);
	//			if ($rows!= = false && intval($rows)<=0)
	//			{
	//			// ���� ������� ���� �� ��������� � ���� ��
	//			// ��������� ���
	//			$new_guest_counter = ($_SESSION["SESS_GUEST_NEW"]=="Y") ? 1: 0;
	//			CStatistics::SetNewDay(
	//			1, // HOSTS
	//			0, // TOTAL_HOSTS (now ignored)
	//			1, // SESSIONS
	//			1,                  // HITS
	//			$new_guest_counter, // NEW_GUESTS
	//			1,                  // GUESTS
	//			$favorites_counter  // FAVORITES
	//			);
	//
	//			// ������� �������� �������� ��� �������� ���
	//			CTraffic::IncParam(
	//			array(
	//			"SESSION" = > 1,
	//			"HIT" = > 1,
	//			"HOST" = > 1,
	//			"GUEST" = > 1,
	//			"NEW_GUEST" = > $new_guest_counter,
	//			"FAVORITE" = > $favorites_counter
	//			)
	//			);
	//			}
	//
	//			// ���� ���� ��������� ��
	//			if ($SITE_ID <> '')
	//			{
	//			// ��������� ������� "�� ����"
	//			$arFields = Array(
	//			"HITS" = > 1,
	//			"HIT" = > 1,
	//			);
	//			// ���� ������� ���� ����� ��������� � ���� ��
	//			// ������� �������� �������� ��� �������� ���
	//			$rows = CTraffic::IncParam(array(), $arFields, $SITE_ID);
	//			if ($rows!= = false && intval($rows)<=0)
	//			{
	//			// ���� ������� ���� ����� �� ��������� � ���� ��
	//			// ��������� ���
	//			CStatistics::SetNewDayForSite(
	//			$SITE_ID,
	//			1,            // HOSTS
	//			0,            // TOTAL_HOSTS (now ignored)
	//			1,            // SESSIONS
	//			1             // HITS
	//			);
	//
	//			// ������� �������� �������� ��� �������� ���
	//			CTraffic::IncParam(
	//			array(),
	//			array(
	//			"SESSION"    = > 1,
	//			"HIT" = > 1,
	//			"HOST"        = > 1,
	//			),
	//			$SITE_ID
	//			);
	//			}
	//			}
	//
	//			/************************************************
	//							���� �� �����
	//			************************************************/
	//
	//			if ($SAVE_PATH_DATA!="N")
	//			CStatistics::SavePathData($SITE_ID, $CURRENT_PAGE, $ERROR_404);
	//
	//			/************************************************
	//						��������� �������� � �������
	//			************************************************/
	//
	//			if ($SAVE_VISITS!="N")
	//			CStatistics::SaveVisits($sql_site, $SESSION_NEW, $CURRENT_DIR, $CURRENT_PAGE, $ERROR_404);
	//
	//			// ��������� ������
	//			$arFields = Array(
	//			//"HITS"			=> "HITS + 1",
	//			"LAST_HIT_ID" => self::$HIT_ID,
	//			"URL_LAST" = > "'".$DB->ForSql($CURRENT_URI, 2000)."'",
	//			"URL_LAST_404" = > "'".$ERROR_404."'",
	//			"DATE_LAST" = > $DB_now,
	//			"LAST_SITE_ID" = > $sql_site
	//			);
	//			if ($SESSION_NEW=="Y") $arFields["FIRST_HIT_ID"] = self::$HIT_ID;
	//			if ($FAVORITES=="Y") $arFields["FAVORITES"] = "'Y'";
	//			$DB->Update("b_stat_session", $arFields, "WHERE ID=".$_SESSION["SESS_SESSION_ID"], "File: ".__FILE__."<br>Line: ".__LINE__, false, false, false);
	//			if ($ERROR_404=="N")
	//			{
	//			CStatistics::Set404("b_stat_session", "ID = ".$_SESSION["SESS_SESSION_ID"], array("URL_LAST_404" = > "Y"));
	//			}
	//
	//			// ��������� �����
	//			$arFields = Array(
	//			"HITS" = > "HITS + 1",
	//			"LAST_SESSION_ID" = > $_SESSION["SESS_SESSION_ID"],
	//			"LAST_DATE" = > $DB_now,
	//			"LAST_USER_ID" = > intval($_SESSION["SESS_LAST_USER_ID"]),
	//			"LAST_USER_AUTH" => "'".$IS_USER_AUTHORIZED."'",
	//			"LAST_URL_LAST" => "'".$DB->ForSql($CURRENT_URI, 2000)."'",
	//			"LAST_URL_LAST_404" = > "'".$ERROR_404."'",
	//			"LAST_USER_AGENT" = > "'".$DB->ForSql($_SERVER["HTTP_USER_AGENT"], 500)."'",
	//			"LAST_IP"        = > "'".$DB->ForSql($_SERVER["REMOTE_ADDR"], 15)."'",
	//			"LAST_COOKIE" = > "'".$DB->ForSql(GetCookieString(), 2000)."'",
	//			"LAST_LANGUAGE" = > "'".$DB->ForSql($_SERVER["HTTP_ACCEPT_LANGUAGE"], 255)."'",
	//			"LAST_SITE_ID" = > $sql_site
	//			);
	//			if ($FAVORITES=="Y") $arFields["FAVORITES"] = "'Y'";
	//			$DB->Update("b_stat_guest", $arFields, "WHERE ID=".intval($_SESSION["SESS_GUEST_ID"]), "File: ".__FILE__."<br>Line: ".__LINE__, false, false,false);
	//			if ($ERROR_404=="N")
	//			{
	//			CStatistics::Set404("b_stat_guest", "ID = ".intval($_SESSION["SESS_GUEST_ID"]), array("LAST_URL_LAST_404" = > "Y"));
	//			}
	//
	//			// ��������� ������ ��������� ��������
	//			if (intval($_SESSION["SESS_ADV_ID"])>0)
	//			{
	//			// ����������� ������� ����� �� ������ ������
	//			$arFields = Array(
	//			"DATE_LAST" = > $DB_now,
	//			"HITS" = > "HITS+1"
	//			);
	//			if ($FAVORITES=="Y" && $ALLOW_ADV_FAVORITES=="Y")
	//			{
	//			// ����������� ������� ����������� ���������� � ��������� �� ������ ������
	//			$arFields["FAVORITES"] = "FAVORITES + 1";
	//			$favorite = 1;
	//			}
	//			$DB->Update("b_stat_adv", $arFields, "WHERE ID=".intval($_SESSION["SESS_ADV_ID"]), "File: ".__FILE__."<br>Line: ".__LINE__, false, false, false);
	//
	//			// ��������� ������� ����� �� ����
	//			$arFields = Array("HITS" = > "HITS+1", "FAVORITES" => "FAVORITES + ".intval($favorite));
	//			$rows = $DB->Update("b_stat_adv_day", $arFields, "WHERE ADV_ID=".intval($_SESSION["SESS_ADV_ID"])." and DATE_STAT=".$DB_now_date,"File: ".__FILE__."<br>Line: ".__LINE__, false,false, false);
	//			// ���� ��� ��� ��
	//			if (intval($rows)<=0)
	//			{
	//			// ��������� ���
	//			$arFields = Array(
	//			"ADV_ID" = > intval($_SESSION["SESS_ADV_ID"]),
	//			"DATE_STAT" = > $DB_now_date,
	//			"HITS" = > 1,
	//			"FAVORITES" => intval($favorite)
	//			);
	//			$DB->Insert("b_stat_adv_day", $arFields, "File: ".__FILE__."<br>Line: ".__LINE__);
	//			}
	//			}
	//			// ��������� ��������� �������� �� ��������
	//			elseif (intval($_SESSION["SESS_LAST_ADV_ID"])>0)
	//			{
	//			// ����������� ������� ����� �� ��������
	//			$arFields = Array(
	//			"DATE_LAST" = > $DB_now,
	//			"HITS_BACK" = > "HITS_BACK+1"
	//			);
	//			if ($FAVORITES=="Y" && $ALLOW_ADV_FAVORITES=="Y")
	//			{
	//			// ����������� ������� ����������� ���������� � ��������� �� ��������
	//			$arFields["FAVORITES_BACK"] = "FAVORITES_BACK + 1";
	//			$favorite = 1;
	//			}
	//			$DB->Update("b_stat_adv", $arFields, "WHERE ID=".intval($_SESSION["SESS_LAST_ADV_ID"]), "File: ".__FILE__."<br>Line: ".__LINE__, false, false, false);
	//
	//			$arFields = Array("HITS_BACK" = > "HITS_BACK+1", "FAVORITES_BACK" => "FAVORITES_BACK + ".intval($favorite));
	//			// ��������� ������� ����� �� ����
	//			$rows = $DB->Update("b_stat_adv_day", $arFields, "WHERE ADV_ID=".intval($_SESSION["SESS_LAST_ADV_ID"])." and DATE_STAT=".$DB_now_date,"File: ".__FILE__."<br>Line: ".__LINE__, false,false, false);
	//			// ���� ��� ��� ��
	//			if (intval($rows)<=0)
	//			{
	//			// ��������� ���
	//			$arFields = Array(
	//			"ADV_ID" = > intval($_SESSION["SESS_LAST_ADV_ID"]),
	//			"DATE_STAT" = > $DB_now_date,
	//			"HITS_BACK" = > 1,
	//			"FAVORITES_BACK" => intval($favorite),
	//			);
	//			$DB->Insert("b_stat_adv_day", $arFields, "File: ".__FILE__."<br>Line: ".__LINE__);
	//			}
	//			}
	//
	//			// ������������ �������
	//			if (defined("GENERATE_EVENT") && GENERATE_EVENT=="Y")
	//			{
	//			global $event1, $event2, $event3, $goto, $money, $currency, $site_id;
	//			if ($site_id == '')
	//			$site_id = false;
	//			CStatistics::Set_Event($event1, $event2, $event3, $goto, $money, $currency, $site_id);
	//			}
	//
	//			// ����������� ������� ����� � ������
	//			if ($_SESSION["SESS_COUNTRY_ID"] <> '')
	//			{
	//			CStatistics::UpdateCountry($_SESSION["SESS_COUNTRY_ID"], Array("HITS" = > 1));
	//			}
	//
	//			if ($_SESSION["SESS_CITY_ID"] > 0)
	//			{
	//			CStatistics::UpdateCity($_SESSION["SESS_CITY_ID"], Array("HITS" = > 1));
	//			}
	//
	//			if (
	//			isset($_SESSION["SESS_FROM_SEARCHERS"])
	//			&& is_array($_SESSION["SESS_FROM_SEARCHERS"])
	//			&& !empty($_SESSION["SESS_FROM_SEARCHERS"])
	//			)
	//			{
	//			// ��������� ������� ����� � ��������� ���� ��� �����������
	//			$arFields = Array("PHRASES_HITS" = > "PHRASES_HITS+1");
	//			$_SESSION["SESS_FROM_SEARCHERS"] = array_unique($_SESSION["SESS_FROM_SEARCHERS"]);
	//			if (count($_SESSION["SESS_FROM_SEARCHERS"]) > 0)
	//			{
	//			$str = "0";
	//			foreach($_SESSION["SESS_FROM_SEARCHERS"] as $value)
	//			$str.= ", ".intval($value);
	//			$DB->Update("b_stat_searcher", $arFields,"WHERE ID in ($str)", "File: ".__FILE__."<br>Line: ".__LINE__,false, false, false);
	//			}
	//			}
	//
	//			if (isset($_SESSION["SESS_REFERER_ID"]) && intval($_SESSION["SESS_REFERER_ID"])>0)
	//			{
	//			// ��������� �����������
	//			$arFields = Array("HITS" = >"HITS+1");
	//			$DB->Update("b_stat_referer", $arFields, "WHERE ID=".intval($_SESSION["SESS_REFERER_ID"]), "File: ".__FILE__."<br>Line: ".__LINE__, false, false, false);
	//			}
	//			}
	//
	//			/*******************************************************
	//				���������� �������� ��������� ���������� ��������
	//			*******************************************************/
	//
	//			$_SESSION["SESS_HTTP_REFERER"] = $_SESSION["SESS_LAST_URI"] ?? ''
	//			$_SESSION["SESS_LAST_PROTOCOL"] = $CURRENT_PROTOCOL
	//			$_SESSION["SESS_LAST_PORT"] = $CURRENT_PORT
	//			$_SESSION["SESS_LAST_HOST"] = $CURRENT_HOST
	//			$_SESSION["SESS_LAST_URI"] = $CURRENT_URI
	//			$_SESSION["SESS_LAST_PAGE"] = $CURRENT_PAGE
	//			$_SESSION["SESS_LAST_DIR"] = $CURRENT_DIR
	//		}
	//	} else // if (!$BLOCK_ACTIVITY)
	//	{
	//		/************************************************
	//			��������� ���������� ������ ����������
	//		*************************************************/
	//
	//		$fname = $_SERVER["DOCUMENT_ROOT"].BX_PERSONAL_ROOT.
	//		"/activity_limit.php"
	//		if file_exists($fname))
	//		{
	//		include($fname);
	//		}
	//		else
	//		{
	//		CHTTP::SetStatus("503 Service Unavailable");
	//		die();
	//		}
	//	}
	//
	//	/************************************************************
	//		������������ �������� ����� �� ��������������
	//		������ �/��� �� ����������� ����
	//	*************************************************************/
	//
	//	// ���� �� ������ select �� ������� b_stat_session_data ��
	//	if ($SESSION_DATA_ID)
	//	{
	//	$arrSTAT_SESSION = stat_session_register(true);
	//	$sess_data_for_db = ($DB->type == "ORACLE") ? "'".$DB->ForSql(serialize($arrSTAT_SESSION), 2000)."'":  "'".$DB->ForSql(serialize($arrSTAT_SESSION))."'";
	//	// ���� � ���������� ����� select'� ���� ������� ������ ��
	//	if ((intval($SESSION_DATA_ID) > 0) && ($SESSION_DATA_ID != = true))
	//	{
	//	// ��������� ��
	//	$arFields = array(
	//	"DATE_LAST" = > $DB_now,
	//	"GUEST_MD5" = > "'".get_guest_md5()."'",
	//	"SESS_SESSION_ID" = > intval($_SESSION["SESS_SESSION_ID"]),
	//	"SESSION_DATA" = > $sess_data_for_db
	//	);
	//	$DB->Update("b_stat_session_data", $arFields, "WHERE ID = ".intval($SESSION_DATA_ID), "File: ".__FILE__."<br>Line: ".__LINE__, false, false, false);
	//	}
	//	else
	//	{
	//	// ����� ��������� ��� ������
	//	$arFields = array(
	//	"DATE_FIRST" = > $DB_now,
	//	"DATE_LAST" = > $DB_now,
	//	"GUEST_MD5" = > "'".get_guest_md5()."'",
	//	"SESS_SESSION_ID" => intval($_SESSION["SESS_SESSION_ID"]),
	//	"SESSION_DATA" = > $sess_data_for_db
	//	);
	//	$DB->Insert("b_stat_session_data", $arFields, "File: ".__FILE__."<br>Line: ".__LINE__);
	//	}
	//	}
	//}
	return nil
}

//func (s Statistic) RestoreSession(phpServer entity.PhpServer, cookieGuestId int, session *session.Session, sessionGcMaxLifeTime string) {
//	// if there is no session ID
//	if session.KeyExists("SESS_SESSION_ID") == false || session.GetAsInt("SESS_SESSION_ID") <= 0 {
//		if session.Get("SAVE_SESSION_DATA") == "Y" {
//			// try to use cookie
//			if cookieGuestId <= 0 {
//				// restore session data from table session_data
//				md5, err := s.GetGuestMd5(phpServer)
//				if err != nil {
//					return
//				}
//				sess, err := s.sessionModel.FindSessionByGuestMd5(md5, sessionGcMaxLifeTime)
//				session.SetAll(&sess)
//			}
//		}
//	}
//}

//func (s Statistic) GetGuestMd5(phpServer entity.PhpServer) (string, error) {
//	var strBuilder strings.Builder
//	strBuilder.WriteString(phpServer.HttpUserAgent)
//	strBuilder.WriteString(phpServer.RemoteAddr)
//	strBuilder.WriteString(phpServer.HttpXForwardedFor)
//	sum, err := checksum.MD5sum(strBuilder.String())
//	return sum, err
//}
