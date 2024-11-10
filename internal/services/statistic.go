package services

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/models"
	"context"
	_ "net/netip"
	"time"
)

type Statistic struct {
	guestService *GuestService
	hitService   *HitService
}

func NewStatistic(guestService *GuestService, hitService *HitService) *Statistic {
	return &Statistic{
		guestService: guestService,
		hitService:   hitService,
	}
}

func (stat *Statistic) Add(ctx context.Context, statData dto.UserData, waitAdd bool) error {
	hash, err := stat.guestService.FindByHash(ctx, statData.GuestHash)
	if err != nil {
		return err
	}

	isNewGuest := hash == models.Guest{}

	if isNewGuest {
		if err = stat.guestService.Add(ctx, models.Guest{GuestHash: statData.GuestHash, DateInsert: time.Now()}); err != nil {
			return err
		}
	}
	return stat.hitService.Add(ctx, statData, isNewGuest, waitAdd)
}
