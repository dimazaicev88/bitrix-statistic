package services

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/entitydb"
	"context"
	_ "net/netip"
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

func (stat *Statistic) Add(ctx context.Context, statData dto.UserData) error {
	hash, err := stat.guestService.FindByHash(ctx, statData.GuestHash)
	if err != nil {
		return err
	}

	isNewGuest := hash == entitydb.Guest{}
	if _, err := stat.hitService.Add(ctx, statData, isNewGuest); err != nil {
		return err
	}
	return nil
}
