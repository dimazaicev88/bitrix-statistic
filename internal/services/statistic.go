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

func (stat Statistic) Add(statData entityjson.StatData) error {
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

			guestDb, err = stat.guestService.AddGuest(statData, advReferer)
			if err != nil {
				return err
			}
		} else { //Если гость уже есть
			existsGuest = true
			if err = stat.guestService.UpdateGuest(guestDb, statData, advReferer); err != nil {
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
			sessionDb, err = stat.sessionService.Add(guestUuid, hitUuid, existsGuest == true, statData, advReferer)
			if err != nil {
				return err
			}
		} else {
			err = stat.sessionService.Update(sessionDb, entitydb.Session{})
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
			//stat.pathService.SavePath()
		}

	}

	return nil
}
