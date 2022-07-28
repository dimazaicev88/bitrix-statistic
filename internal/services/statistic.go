package services

import (
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/storage"
	"bitrix-statistic/utils"
	"golang.org/x/exp/slices"
	"regexp"
	"strconv"
	"strings"
)

type StatisticService struct {
	statisticModel models.StatisticModel
	optionModel    models.OptionModel
}

func NewStatisticService(storage *storage.MysqlStorage) StatisticService {
	return StatisticService{
		statisticModel: models.NewStatisticModel(storage),
		optionModel:    models.NewOptionModel(storage),
	}
}

func (ss StatisticService) checkSkip(userGroups []int, remoteAddr string) (error, bool) {
	skipMode := ss.optionModel.GetOption("SKIP_STATISTIC_WHAT")

	isSkip := false
	switch skipMode {
	case "none":
		break
	case "both":
	case "groups":
		arSkipGroups := strings.Split(",", ss.optionModel.GetOption("SKIP_STATISTIC_GROUPS"))
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
			arSkipIPRanges := strings.Split("\n", ss.optionModel.GetOption("SKIP_STATISTIC_IP_RANGES"))
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
