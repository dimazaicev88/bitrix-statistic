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
	otterCache      otter.Cache[string, entitydb.Session]
}

func NewStatistic() *Statistic {
	otterCache, err := otter.MustBuilder[string, entitydb.Session](15000).
		CollectStats().
		WithTTL(time.Minute * 15).
		Build()

	if err != nil {
		logrus.Fatal(err)
	}

	return &Statistic{
		otterCache: otterCache,
	}
}

func (stat *Statistic) SetHitService(hitService *HitService) {
	stat.hitService = hitService
}

func (stat *Statistic) SetAdvServices(advServices *AdvServices) {
	stat.advServices = advServices
}

func (stat *Statistic) SetGuestService(guestService *GuestService) {
	stat.guestService = guestService
}

func (stat *Statistic) SetPathService(pathService *PathService) {
	stat.pathService = pathService
}

func (stat *Statistic) SetSessionService(sessionService *SessionService) {
	stat.sessionService = sessionService
}

func (stat *Statistic) SetStatDayService(statDayService *StatDayService) {
	stat.statDayService = statDayService
}

func (stat *Statistic) SetSearcherService(searcherService *SearcherService) {
	stat.searcherService = searcherService
}

func (stat *Statistic) SetOptionService(optionService *OptionService) {
	stat.optionService = optionService
}

func (stat *Statistic) SetRefererService(refererService *RefererService) {
	stat.refererService = refererService
}

func (stat *Statistic) Add(statData entityjson.UserData) error {
	var advReferer entitydb.AdvCompany
	var sessionDb entitydb.Session
	var guestDb entitydb.Guest
	isNewGuest := true
	var hitUuid = uuid.New()
	var sessionUuid = uuid.New()
	favoriteDbValue := 0

	if statData.IsFavorite {
		favoriteDbValue = 1
	}

	isSearcher, err := stat.searcherService.IsSearcher(statData.UserAgent)
	if err != nil {
		return err
	}

	if isSearcher { //Это поисковик, не учитываем его как гостя
		if err = stat.searcherService.AddHitSearcher(statData); err != nil { //Обновляем статистику за 1 день
			return err
		}
	} else {

		//--------------------------- Adv --------------------------------------
		advReferer, err = stat.advServices.GetAdv(statData) //Получаем рекламную компанию
		if err != nil {
			return err
		}

		if advReferer == (entitydb.AdvCompany{}) { //Автоматическое создание рекламной компании
			parsedURL, err := url.Parse(statData.Url)
			if err != nil {
				return err
			}
			queryParams := parsedURL.Query()
			advDb, err := stat.advServices.AutoCreateAdv(queryParams.Get("referrer1"), queryParams.Get("referrer2"))
			if err != nil {
				return err
			}

			if advDb != (entitydb.Adv{}) {
				advReferer.AdvUuid = advDb.Uuid
				advReferer.Referer1 = advDb.Referer1
				advReferer.Referer2 = advDb.Referer2
			}
		}

		//--------------------------- Guest ------------------------------------
		guestDb, err = stat.guestService.FindByUuid(statData.GuestUuid)
		if err != nil {
			return err
		}

		//Гость не найден, добавляем гостя
		if guestDb == (entitydb.Guest{}) {
			guestDb, err = stat.guestService.Add(statData)
			if err != nil {
				return err
			}
		} else {
			isNewGuest = false
		}

		//--------------------------- Sessions ------------------------------------
		sessionDb, err = stat.sessionService.FindByPHPSessionId(statData.PHPSessionId)
		if err != nil {
			return err
		}

		//Если сессия новая, добавляем.
		if sessionDb == (entitydb.Session{}) {
			sessionDb, err = stat.sessionService.Add(sessionUuid, guestDb.Uuid, hitUuid, statData.PHPSessionId)
			if err != nil {
				return err
			}
		}

		//------------------------------- Hits ---------------------------------
		if stat.optionService.IsSaveHits() {
			if _, err = stat.hitService.Add(hitUuid, isNewGuest, sessionUuid, advReferer, statData); err != nil {
				return err
			}
		}

		//----------------------- ADV stat ----------------------------
		countNewGuests := 0
		if isNewGuest {
			countNewGuests = 1
		}
		err = stat.advServices.AddAdvStat(entitydb.AdvStat{
			AdvUuid:       advReferer.AdvUuid,
			Guests:        1,
			NewGuests:     uint32(countNewGuests),
			Favorites:     uint32(favoriteDbValue),
			Hosts:         0, // Это уникальный ip //TODO проверить что это уникальный ip
			Sessions:      1,
			Hits:          1,
			GuestsBack:    0,
			FavoritesBack: 0,
			HostsBack:     0,
			SessionsBack:  0,
			HitsBack:      0,
		})
		if err != nil {
			return err
		}

		//------------------------- Referring -------------------------
		if stat.optionService.IsSaveReferrers() {
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
		if stat.optionService.IsSavePathData() {
			err = stat.pathService.SavePath(sessionDb.Uuid, statData.SiteId, statData.Url, statData.Referer, statData.IsError404, advReferer)
			if err != nil {
				return err
			}
		}
		//------------------------------ Visits -----------------------------
		if stat.optionService.IsSaveVisits() {
			//err = stat.pathService.SaveVisits()
			//if err != nil {
			//	return err
			//}
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

func (stat *Statistic) ClearCache() {
	stat.otterCache.Close()
}
