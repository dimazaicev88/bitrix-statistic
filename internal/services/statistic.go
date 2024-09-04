package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"github.com/google/uuid"
	"github.com/maypok86/otter"
	"github.com/sirupsen/logrus"
	_ "net/netip"
	"net/url"
	"time"
)

type Statistic struct {
	advServices     *AdvServices
	guestService    *GuestService
	sessionService  *SessionService
	statDayService  *StatDayService
	searcherService *SearcherService
	optionService   *OptionService
	hitService      *HitService
	refererService  *RefererService
	pathService     *PathService
	sessionCache    otter.Cache[string, entitydb.Session]
}

func NewStatistic(
	hitService *HitService,
	advServices *AdvServices,
	guestService *GuestService,
	pathService *PathService,
	sessionService *SessionService,
	statDayService *StatDayService,
	searcherService *SearcherService,
	optionService *OptionService,
	refererService *RefererService,
) *Statistic {
	otterCache, err := otter.MustBuilder[string, entitydb.Session](15000).
		CollectStats().
		WithTTL(time.Minute * 15).
		Build()

	if err != nil {
		logrus.Fatal(err)
	}

	return &Statistic{
		guestService:    guestService,
		advServices:     advServices,
		sessionService:  sessionService,
		statDayService:  statDayService,
		searcherService: searcherService,
		hitService:      hitService,
		optionService:   optionService,
		refererService:  refererService,
		sessionCache:    otterCache,
		pathService:     pathService,
	}
}

func (stat Statistic) Add(statData entityjson.UserData) error {
	//var stopListUuid string

	//var advBack string
	var advReferer entitydb.AdvReferer
	var sessionDb entitydb.Session
	var guestDb entitydb.Guest
	existsGuest := false
	var guestUuid uuid.UUID
	var hitUuid uuid.UUID

	isSearcher, err := stat.searcherService.IsSearcher(statData.UserAgent)
	if err != nil {
		return err
	}

	if isSearcher { //Это поисковик, не учитываем его как гостя
		if err = stat.searcherService.AddHitSearcher(statData); err != nil { //Обновляем статистику за 1 день
			return err
		}
	} else {
		guestDb, err = stat.guestService.FindByUuid(statData.GuestUuid)
		if err != nil {
			return err
		}

		//--------------------------- Guest ------------------------------------
		sessionDb, err = stat.sessionService.FindByPHPSessionId(statData.PHPSessionId)
		if err != nil {
			return err
		}

		//Гость не найден, добавляем гостя
		if guestDb == (entitydb.Guest{}) {
			advReferer, err = stat.advServices.GetAdv(statData) //Получаем рекламную компанию
			//TODO добавить установку дефолтной рекламной компании, в случае если  не установлена рекламная компания
			//TODO добавить авто создание рекламной компании

			if err != nil {
				return err
			}

			guestDb, err = stat.guestService.Add(statData, advReferer)
			if err != nil {
				return err
			}
		} else { //Если гость уже есть
			existsGuest = true
			newGuestDb := guestDb
			newGuestDb.Sessions += 1
			if err = stat.guestService.UpdateGuest(guestDb, newGuestDb); err != nil {
				return err
			}
		}

		//------------------------------- Hits ---------------------------------
		if stat.optionService.IsSaveHits(statData.SiteId) {
			//if hitUuid, err = stat.hitService.Add(existsGuest, sessionDb, advReferer, statData); err != nil {
			//	return err
			//}
		}

		//--------------------------- Sessions ------------------------------------

		//Если сессия новая, добавляем.
		if sessionDb == (entitydb.Session{}) {
			sessionDb, err = stat.sessionService.Add(uuid.Nil, guestUuid, hitUuid, existsGuest == true, statData, advReferer)
			if err != nil {
				return err
			}
		} else {
			err = stat.sessionService.Update(sessionDb, entitydb.Session{
				UserId:     statData.UserId,
				IsUserAuth: statData.IsUserAuth,
				UserAgent:  statData.UserAgent,
				IpLast:     statData.Ip,
				Hits:       sessionDb.Hits + 1,
			})
			if err != nil {
				return err
			}
		}

		//------------------------- Referring -------------------------
		if stat.optionService.IsSaveReferrers(statData.SiteId) {
			parse, err := url.Parse(statData.Referer)
			if err != nil {
				return err
			}
			if len(statData.Referer) > 0 {
				idReferer, err := stat.refererService.Add(statData.Referer)
				if err != nil {
					return err
				}
				_, err = stat.refererService.AddToRefererList(advReferer.AdvUuid, sessionDb.Uuid, idReferer, parse, statData)
				if err != nil {
					return err
				}
			}

			// TODO ADD Search phrases
		}

		//------------------------------ Path data -----------------------------
		if stat.optionService.IsSavePathData(statData.SiteId) {
			err = stat.pathService.SavePath(sessionDb.Uuid, statData.SiteId, statData.Url, statData.Referer, statData.IsError404, advReferer)
			if err != nil {
				return err
			}
		}
		//------------------------------ Visits -----------------------------
		if stat.optionService.IsSaveVisits(statData.SiteId) {
			err = stat.pathService.SaveVisits()
			if err != nil {
				return err
			}
		}
		newSession := sessionDb
		newSession.LastHitUuid = hitUuid
		newSession.UrlLast = statData.Url
		newSession.UrlLast404 = statData.IsError404
		newSession.DateLast = time.Now()
		newSession.LastSiteId = statData.SiteId

		if err = stat.sessionService.Update(sessionDb, newSession); err != nil {
			return err
		}

		newGuestDb := guestDb
		newGuestDb.Hits += 1
		newGuestDb.LastSessionUuid = sessionDb.Uuid
		newGuestDb.LastDate = time.Now()
		newGuestDb.LastUserId = statData.UserId
		newGuestDb.LastUserAuth = statData.IsUserAuth
		newGuestDb.LastUrlLast = statData.Url
		newGuestDb.LastUrlLast404 = statData.IsError404
		newGuestDb.LastUserAgent = statData.UserAgent
		newGuestDb.LastIp = statData.Ip
		newGuestDb.LastCookie = statData.Cookies
		newGuestDb.LastLanguage = statData.Lang
		newGuestDb.LastSiteId = statData.SiteId
		newGuestDb.Favorites = statData.IsFavorite
		if err = stat.guestService.UpdateGuest(guestDb, newGuestDb); err != nil {
			return err
		}

		//TODO
		/**
		// обновляем прямые рекламные кампании
						if (intval($_SESSION["SESS_ADV_ID"])>0)
						{
							// увеличиваем счетчик хитов на прямом заходе
							$arFields = Array(
								"DATE_LAST"	=> $DB_now,
								"HITS"		=> "HITS+1"
								);
							if ($FAVORITES=="Y" && $ALLOW_ADV_FAVORITES=="Y")
							{
								// увеличиваем счетчик посетителей добавивших в избранное на прямом заходе
								$arFields["FAVORITES"] = "FAVORITES + 1";
								$favorite = 1;
							}
							$DB->Update("b_stat_adv",$arFields,"WHERE ID=".intval($_SESSION["SESS_ADV_ID"]), "File: ".__FILE__."<br>Line: ".__LINE__,false,false,false);

							// обновляем счетчик хитов по дням
							$arFields = Array("HITS" => "HITS+1", "FAVORITES" => "FAVORITES + ".intval($favorite));
							$rows = $DB->Update("b_stat_adv_day",$arFields,"WHERE ADV_ID=".intval($_SESSION["SESS_ADV_ID"])." and DATE_STAT=".$DB_now_date,"File: ".__FILE__."<br>Line: ".__LINE__,false,false,false);
							// если его нет то
							if (intval($rows)<=0)
							{
								// добавляем его
								$arFields = Array(
									"ADV_ID"		=> intval($_SESSION["SESS_ADV_ID"]),
									"DATE_STAT"		=> $DB_now_date,
									"HITS"			=> 1,
									"FAVORITES"		=> intval($favorite)
									);
								$DB->Insert("b_stat_adv_day",$arFields, "File: ".__FILE__."<br>Line: ".__LINE__);
							}
						}
						// обновляем рекламные кампании по возврату
						elseif (intval($_SESSION["SESS_LAST_ADV_ID"])>0)
						{
							// увеличиваем счетчик хитов на возврате
							$arFields = Array(
								"DATE_LAST"		=> $DB_now,
								"HITS_BACK"		=> "HITS_BACK+1"
								);
							if ($FAVORITES=="Y" && $ALLOW_ADV_FAVORITES=="Y")
							{
								// увеличиваем счетчик посетителей добавивших в избранное на возврате
								$arFields["FAVORITES_BACK"] = "FAVORITES_BACK + 1";
								$favorite = 1;
							}
							$DB->Update("b_stat_adv",$arFields,"WHERE ID=".intval($_SESSION["SESS_LAST_ADV_ID"]), "File: ".__FILE__."<br>Line: ".__LINE__,false,false,false);

							$arFields = Array("HITS_BACK" => "HITS_BACK+1", "FAVORITES_BACK" => "FAVORITES_BACK + ".intval($favorite));
							// обновляем счетчик хитов по дням
							$rows = $DB->Update("b_stat_adv_day",$arFields,"WHERE ADV_ID=".intval($_SESSION["SESS_LAST_ADV_ID"])." and DATE_STAT=".$DB_now_date,"File: ".__FILE__."<br>Line: ".__LINE__,false,false,false);
							// если его нет то
							if (intval($rows)<=0)
							{
								// добавляем его
								$arFields = Array(
									"ADV_ID" => intval($_SESSION["SESS_LAST_ADV_ID"]),
									"DATE_STAT" => $DB_now_date,
									"HITS_BACK" => 1,
									"FAVORITES_BACK" => intval($favorite),
								);
								$DB->Insert("b_stat_adv_day",$arFields, "File: ".__FILE__."<br>Line: ".__LINE__);
							}
						}

						// обрабатываем событие
						if (defined("GENERATE_EVENT") && GENERATE_EVENT=="Y")
						{
							global $event1, $event2, $event3, $goto, $money, $currency, $site_id;
							if($site_id == '')
								$site_id = false;
							CStatistics::Set_Event($event1, $event2, $event3, $goto, $money, $currency, $site_id);
						}

						// увеличиваем счетчик хитов у страны
						if ($_SESSION["SESS_COUNTRY_ID"] <> '')
						{
							CStatistics::UpdateCountry($_SESSION["SESS_COUNTRY_ID"], Array("HITS" => 1));
						}

						if($_SESSION["SESS_CITY_ID"] > 0)
						{
							CStatistics::UpdateCity($_SESSION["SESS_CITY_ID"], Array("HITS" => 1));
						}

						if (
							isset($_SESSION["SESS_FROM_SEARCHERS"])
							&& is_array($_SESSION["SESS_FROM_SEARCHERS"])
							&& !empty($_SESSION["SESS_FROM_SEARCHERS"])
						)
						{
							// обновляем счетчик хитов у поисковых фраз для поисковиков
							$arFields = Array("PHRASES_HITS" => "PHRASES_HITS+1");
							$_SESSION["SESS_FROM_SEARCHERS"] = array_unique($_SESSION["SESS_FROM_SEARCHERS"]);
							if(count($_SESSION["SESS_FROM_SEARCHERS"]) > 0)
							{
								$str = "0";
								foreach($_SESSION["SESS_FROM_SEARCHERS"] as $value)
									$str .= ", ".intval($value);
								$DB->Update("b_stat_searcher",$arFields,"WHERE ID in ($str)", "File: ".__FILE__."<br>Line: ".__LINE__,false,false,false);
							}
						}

						if (isset($_SESSION["SESS_REFERER_ID"]) && intval($_SESSION["SESS_REFERER_ID"])>0)
						{
							// обновляем ссылающиеся
							$arFields = Array("HITS"=>"HITS+1");
							$DB->Update("b_stat_referer", $arFields, "WHERE ID=".intval($_SESSION["SESS_REFERER_ID"]), "File: ".__FILE__."<br>Line: ".__LINE__,false,false,false);
						}
		*/

	}

	return nil
}
