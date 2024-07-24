package services

import (
	"bitrix-statistic/internal/entityjson"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
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

func NewStatistic(ctx context.Context, chClient driver.Conn) *Statistic {
	return &Statistic{
		guestService:    NewGuest(ctx, chClient),
		advServices:     NewAdv(ctx, chClient),
		sessionService:  NewSession(ctx, chClient),
		statDayService:  NewStatDay(ctx, chClient),
		searcherService: NewSearcher(ctx, chClient),
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
		existsGuest, err := stat.guestService.guestModel.ExistsGuestByHash(statData.GuestHash)
		if err != nil {
			return err
		}
		//---------------------------Секция гостя------------------------------------
		var guestUuid uuid.UUID
		var stopListUuid uuid.UUID
		var advBack string
		var cityUuid string
		var countryUuid uuid.UUID

		//Гость не найден, добавляем гостя
		if existsGuest == false {
			adv, err := stat.advServices.GetAdv(statData.Url)

			if err != nil {
				return err
			}

			err = stat.guestService.AddGuest(statData, adv)
			if err != nil {
				return err
			}
		}

		//---------------------------Секция сессии------------------------------------
		//Если сессия новая, добавляем.
		if stat.sessionService.IsExistsSession(statData.PHPSessionId) == false {
			isNewGuest := existsGuest == false
			err := stat.sessionService.AddSession(advBack, cityUuid, countryUuid, stopListUuid, guestUuid, isNewGuest, statData)
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
