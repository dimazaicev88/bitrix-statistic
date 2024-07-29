package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/models"
	"context"
	_ "net/netip"
)

type Statistic struct {
	advServices     *AdvServices
	guestService    *GuestService
	sessionService  *SessionService
	statDayService  *StatDayService
	searcherService *SearcherService
	hitService      *HitService
}

func NewStatistic(ctx context.Context, allModels *models.Models) *Statistic {
	return &Statistic{
		guestService:    NewGuest(ctx, allModels),
		advServices:     NewAdv(ctx, allModels),
		sessionService:  NewSession(ctx, allModels),
		statDayService:  NewStatDay(ctx, allModels),
		searcherService: NewSearcher(ctx, allModels),
		hitService:      NewHit(ctx, allModels),
	}
}

func (stat Statistic) Add(statData entityjson.StatData) error {
	var stopListUuid string
	var guestUuid string
	var advBack string
	var advReferer entitydb.AdvReferer
	var sessionDb entitydb.Session
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
		//---------------------------Секция гостя------------------------------------

		//Гость не найден, добавляем гостя
		if existsGuest == false {
			advReferer, err = stat.advServices.GetAdv(statData.Url) //Получаем рекламную компанию
			//TODO добавить установку дефолтной рекламной компании, в случае если  не установлена рекламная компания

			if err != nil {
				return err
			}

			guestUuid, err = stat.guestService.AddGuest(statData)
			if err != nil {
				return err
			}
		}

		//---------------------------Секция сессии------------------------------------
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

		stat.hitService.Add()

		stat.statDayService.Add() //TODO доделать update
	}
	return nil
}
