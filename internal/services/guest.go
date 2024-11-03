package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/models"
	"context"
	"errors"
)

type GuestService struct {
	guestModel *models.Guest
}

func NewGuest(guestModel *models.Guest) *GuestService {
	return &GuestService{
		guestModel: guestModel,
	}
}

func (gs *GuestService) Add(ctx context.Context, guest entitydb.Guest) error {

	if guest == (entitydb.Guest{}) {
		return errors.New("guest is empty")
	}

	if err := gs.guestModel.Add(ctx, guest); err != nil {
		return err
	}
	return nil
}

func (gs *GuestService) FindByHash(ctx context.Context, hash string) (entitydb.Guest, error) {
	return gs.guestModel.FindByHash(ctx, hash)
}
