package services

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	_ "net/netip"
)

type Statistic struct {
	statisticModel *models.StatisticModel
	optionModel    *models.OptionModel
	searcherModel  *models.SearcherModel
	sessionModel   *models.SessionModel
	cityModel      *models.CityModel
	advServices    *AdvServices
	guestService   *GuestService
	sessionService *SessionService
}

func NewStatistic(ctx context.Context, chClient driver.Conn) Statistic {
	return Statistic{
		statisticModel: models.NewStatisticModel(ctx, chClient),
		optionModel:    models.NewOptionModel(ctx, chClient),
		sessionModel:   models.NewSessionModel(ctx, chClient),
		searcherModel:  models.NewSearcherModel(ctx, chClient),
		cityModel:      models.NewCityModel(ctx, chClient),
		guestService:   NewGuestService(ctx, chClient),
		advServices:    NewAdvServices(ctx, chClient),
		sessionService: NewSessionService(ctx, chClient),
	}
}

func (s Statistic) checkSkip(userGroups []int, remoteAddr string) (error, bool) {
	//skipMode := s.optionModel.Get("SKIP_STATISTIC_WHAT")

	isSkip := false
	//switch skipMode {
	//case "none":
	//	break
	//case "both":
	//case "groups":
	//	arSkipGroups := strings.Split(",", s.optionModel.Get("SKIP_STATISTIC_GROUPS"))
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
	//	var re = regexp.MustCompile(`/^.*?(\d+)\.(\d+)\.(\d+)\.(\d+)[\s-]*/`)
	//	arIPAAddress := re.FindStringSubmatch(remoteAddr)
	//	if len(re.FindStringSubmatch(remoteAddr)) > 0 {
	//		arSkipIPRanges := strings.Split("\n", s.optionModel.Get("SKIP_STATISTIC_IP_RANGES"))
	//		for _, skipRange := range arSkipIPRanges {
	//			var re = regexp.MustCompile(`/^.*?(\d+)\.(\d+)\.(\d+)\.(\d+)[\s-]*(\d+)\.(\d+)\.(\d+)\.(\d+)/`)
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

func (s Statistic) Add(statData entity.StatData) error {
	existsGuest, err := s.guestService.GuestModel.ExistsGuestByHash(statData.GuestHash)
	if err != nil {
		return err
	}

	//---------------------------Секция гостя------------------------------------
	var guestUuid string
	var stopListUuid string
	var advBack string
	var cityUuid string
	var countryUuid string

	//Гость не найден, добавляем гостя
	if existsGuest == false {
		adv, err := s.advServices.GetAdv(statData.Url)

		if err != nil {
			return err
		}

		err = s.guestService.AddGuest(statData, adv)
		if err != nil {
			return err
		}
	}

	//---------------------------Секция сессии------------------------------------
	//Если сессия новая, добавляем.
	if s.sessionService.IsExistsSession(statData.PHPSessionId) == false {
		isNewGuest := existsGuest == false
		err := s.sessionService.AddSession(advBack, cityUuid, countryUuid, stopListUuid, guestUuid, isNewGuest, statData)
		if err != nil {
			return err
		}
	} else { // Обновляем имеющуюся
		s.sessionService.UpdateSession(statData.PHPSessionId, statData)

	}
	return nil
}
