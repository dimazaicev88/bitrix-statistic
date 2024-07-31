package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/models"
	"context"
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
	sessionCache    otter.Cache[string, entitydb.Session]
}

func NewStatistic(ctx context.Context, allModels *models.Models) *Statistic {
	otterCache, err := otter.MustBuilder[string, entitydb.Session](15000).
		CollectStats().
		WithTTL(time.Minute * 15).
		Build()

	if err != nil {
		logrus.Fatal(err)
	}
	return &Statistic{
		guestService:    NewGuest(ctx, allModels),
		advServices:     NewAdv(ctx, allModels),
		sessionService:  NewSession(ctx, allModels),
		statDayService:  NewStatDay(ctx, allModels),
		searcherService: NewSearcher(ctx, allModels),
		hitService:      NewHit(ctx, allModels),
		optionService:   NewOption(ctx, allModels),
		refererService:  NewReferer(ctx, allModels),
		sessionCache:    otterCache,
	}
}

func (stat Statistic) Add(statData entityjson.StatData) error {
	//var stopListUuid string
	var guestUuid string
	//var advBack string
	var advReferer entitydb.AdvReferer
	var sessionDb entitydb.Session
	var guestDb entitydb.Guest
	var existsGuest = false

	isSearcher, err := stat.searcherService.IsSearcher(statData.UserAgent)
	if err != nil {
		return err
	}

	if isSearcher { //Это поисковик, не учитываем его как гостя
		if err = stat.searcherService.AddHitSearcher(statData); err != nil { //Обновляем статистику за 1 день
			return err
		}
	} else {
		existsGuest, err = stat.guestService.ExistsGuestByHash(statData.GuestHash)
		if err != nil {
			return err
		}
		//--------------------------- Guest ------------------------------------

		//Гость не найден, добавляем гостя
		if existsGuest == false {
			advReferer, err = stat.advServices.GetAdv(statData.Url) //Получаем рекламную компанию
			//TODO добавить установку дефолтной рекламной компании, в случае если  не установлена рекламная компания
			//TODO добавить авто создание рекламной компании

			if err != nil {
				return err
			}

			guestUuid, err = stat.guestService.AddGuest(statData, advReferer)
			if err != nil {
				return err
			}
		} else { //Если гость уже есть

		}

		//--------------------------- Sessions ------------------------------------
		sessionDb, err = stat.sessionService.FindByPHPSessionId(statData.PHPSessionId)
		if err != nil {
			return err
		}

		//Если сессия новая, добавляем.
		if sessionDb == (entitydb.Session{}) {
			sessionUuid, err := stat.sessionService.Add(guestUuid, statData.PHPSessionId)
			if err != nil {
				return err
			}
			sessionDb = entitydb.Session{
				Uuid:         sessionUuid,
				GuestUuid:    guestUuid,
				PhpSessionId: statData.PHPSessionId,
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
				err = stat.refererService.AddToRefererList(entitydb.RefererList{
					Uuid:        uuid.New().String(),
					RefererId:   idReferer,
					DateHit:     time.Time{},
					Protocol:    parse.Scheme,
					SiteName:    parse.Hostname(),
					UrlFrom:     statData.Referer,
					UrlTo:       statData.Url,
					UrlTo404:    statData.IsError404,
					SessionUuid: sessionDb.Uuid,
					AdvUuid:     advReferer.AdvUuid,
					SiteId:      statData.SiteId,
				})
				if err != nil {
					return err
				}
			}

			// TODO ADD Search phrases
		}

		//------------------------------- Hits ---------------------------------
		if stat.optionService.IsSaveHits(statData.SiteId) {
			if err = stat.hitService.Add(existsGuest, sessionDb, advReferer, statData); err != nil {
				return err
			}
		}

		//------------------------------ Path data -----------------------------
		if stat.optionService.IsSavePathData(statData.SiteId) {

		}

	}

	return nil
}
