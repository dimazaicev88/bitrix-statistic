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

func NewStatistic(ctx context.Context, chClient driver.Conn) Statistic {
	return Statistic{
		guestService:    NewGuest(ctx, chClient),
		advServices:     NewAdv(ctx, chClient),
		sessionService:  NewSession(ctx, chClient),
		statDayService:  NewStatDay(ctx, chClient),
		searcherService: NewSearcherService(ctx, chClient),
	}
}

func (stat Statistic) checkSkip(userGroups []int, remoteAddr string) (error, bool) {
	//skipMode := stat.optionModel.Get("SKIP_STATISTIC_WHAT")

	isSkip := false
	//switch skipMode {
	//case "none":
	//	break
	//case "both":
	//case "groups":
	//	arSkipGroups := strings.Split(",", stat.optionModel.Get("SKIP_STATISTIC_GROUPS"))
	//	for _, group := range arSkipGroups {
	//		groupId, err := strconv.Atoi(group)
	//		if err != nil {
	//			return err, false
	//		}
	//		if slices.Contains(userGroups, groupId) {
	//			isSkip = true
	//		}
	//	}
	//case "ranges":
	//	if skipMode == "both" && isSkip == true {
	//		break
	//	}
	//	isSkip = true
	//	var re = regexp.MustCompile(`/^.*?(\d+)\.(\d+)\.(\d+)\.(\d+)[\stat-]*/`)
	//	arIPAAddress := re.FindStringSubmatch(remoteAddr)
	//	if len(re.FindStringSubmatch(remoteAddr)) > 0 {
	//		arSkipIPRanges := strings.Split("\n", stat.optionModel.Get("SKIP_STATISTIC_IP_RANGES"))
	//		for _, skipRange := range arSkipIPRanges {
	//			var re = regexp.MustCompile(`/^.*?(\d+)\.(\d+)\.(\d+)\.(\d+)[\stat-]*(\d+)\.(\d+)\.(\d+)\.(\d+)/`)
	//			matchSkipRange := re.FindStringSubmatch(skipRange)
	//			if len(matchSkipRange) > 0 {
	//				if utils.StrToInt(arIPAAddress[1]) >= int(skipRange[1]) &&
	//					utils.StrToInt(arIPAAddress[2]) >= int(skipRange[2]) &&
	//					utils.StrToInt(arIPAAddress[3]) >= int(skipRange[3]) &&
	//					utils.StrToInt(arIPAAddress[4]) >= int(skipRange[4]) &&
	//					utils.StrToInt(arIPAAddress[1]) <= int(skipRange[5]) &&
	//					utils.StrToInt(arIPAAddress[2]) <= int(skipRange[6]) &&
	//					utils.StrToInt(arIPAAddress[3]) <= int(skipRange[7]) &&
	//					utils.StrToInt(arIPAAddress[4]) <= int(skipRange[8]) {
	//					isSkip = true
	//					break
	//				}
	//			}
	//		}
	//	}
	//	break
	//}
	return nil, isSkip
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
		existsGuest, err := stat.guestService.GuestModel.ExistsGuestByHash(statData.GuestHash)
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
