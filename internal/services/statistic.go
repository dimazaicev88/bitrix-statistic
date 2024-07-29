package services

import (
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/google/uuid"
	_ "net/netip"
)

type Statistic struct {
	advServices     *AdvServices
	guestService    *GuestService
	sessionService  *SessionService
	statDayService  *StatDayService
	searcherService *SearcherService
}

func NewStatistic(ctx context.Context, allModels *models.Models) *Statistic {
	return &Statistic{
		guestService:    NewGuest(ctx, allModels),
		advServices:     NewAdv(ctx, allModels),
		sessionService:  NewSession(ctx, allModels),
		statDayService:  NewStatDay(ctx, allModels),
		searcherService: NewSearcher(ctx, allModels),
	}
}

func (stat Statistic) Add(statData entityjson.StatData) error {
	isSearcher, err := stat.searcherService.IsSearcher(statData.UserAgent)
	if err != nil {
		return err
	}

	if isSearcher { //Это поисковик, не учитываем его как гостя
		if err = stat.searcherService.AddHitSearcher(statData); err != nil { //Обновляем статистику за 1 день
			return err
		}
	} else {
		existsGuest, err := stat.guestService.ExistsGuestByHash(statData.GuestHash)
		if err != nil {
			return err
		}
		//---------------------------Секция гостя------------------------------------
		var guestUuid string
		var stopListUuid string
		var advBack string
		var cityUuid string
		var countryUuid uuid.UUID

		//Гость не найден, добавляем гостя
		if existsGuest == false {
			adv, err := stat.advServices.GetAdv(statData.Url) //Получаем рекламную компанию
			//TODO добавить установку дефолтной рекламной компании, в случае если  не установлена рекламная компания

			if err != nil {
				return err
			}

			err = stat.guestService.AddGuest(statData)
			if err != nil {
				return err
			}
		}

		//---------------------------Секция сессии------------------------------------
		//Если сессия новая, добавляем.
		if stat.sessionService.IsExistsSession(statData.PHPSessionId) == false {
			isNewGuest := existsGuest == false
			err := stat.sessionService.Add(guestUuid, statData.PHPSessionId)
			if err != nil {
				return err
			}
		} else { // Обновляем имеющуюся
			err := stat.sessionService.UpdateSession(statData)
			if err != nil {
				return err
			}
		}

		stat.statDayService.Update() //TODO доделать update
	}
	return nil
}
