package services

import (
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/storage"
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

func (ss StatisticService) checkSkip(userGroups []int, remoteAddr string) {
	skipMode := ss.optionModel.GetOption("SKIP_STATISTIC_WHAT")

	switch skipMode {
	case "none":
		break
	case "both":
	case "groups":
		arSkipGroups := strings.Split(",", ss.optionModel.GetOption("SKIP_STATISTIC_GROUPS"))
		foreach($arSkipGroups
		as $key => $value) {
		if in_array(intval($value), $arUserGroups)) {
		$GO = false;
		break;
		}
	}
		if ($skipMode == "groups")
		break
	//else
	//	continue checking
	case "ranges":
		if ($skipMode == "both" && $GO)
		break //in case group check failed
		$GO = true
		if preg_match("/^.*?(\d+)\.(\d+)\.(\d+)\.(\d+)[\s-]*/", $_SERVER["REMOTE_ADDR"], $arIPAdress)) {
	$arSkipIPRanges = explode("\n", COption::GetOptionString("statistic", "SKIP_STATISTIC_IP_RANGES"));
	foreach ($arSkipIPRanges as $key = > $value) {
	if (preg_match("/^.*?(\d+)\.(\d+)\.(\d+)\.(\d+)[\s-]*(\d+)\.(\d+)\.(\d+)\.(\d+)/", $value, $arIPRange)) {
	if (
	intval($arIPAdress[1]) >= intval($arIPRange[1]) && intval($arIPAdress[1]) <= intval($arIPRange[5]) &&
	intval($arIPAdress[2]) >= intval($arIPRange[2]) && intval($arIPAdress[2]) <= intval($arIPRange[6]) &&
	intval($arIPAdress[3]) >= intval($arIPRange[3]) && intval($arIPAdress[3]) <= intval($arIPRange[7]) &&
	intval($arIPAdress[4]) >= intval($arIPRange[4]) && intval($arIPAdress[4]) <= intval($arIPRange[8])
	) {
	$GO = false;
	break;
	}
	}
	}
	}
		break
	}
}
