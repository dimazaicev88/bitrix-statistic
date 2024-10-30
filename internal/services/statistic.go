package services

import (
	"bitrix-statistic/internal/dto"
	_ "net/netip"
)

type Statistic struct {
	hitService *HitService
}

func NewStatistic(hitService *HitService) *Statistic {
	return &Statistic{
		hitService: hitService,
	}
}

func (stat *Statistic) Add(statData dto.UserData) error {
	return stat.hitService.Add(statData)
}
