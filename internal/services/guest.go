package services

import (
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/repository"
	"context"
	"errors"
)

type GuestService struct {
	guestRepository *repository.Guest
}

func NewGuest(guestRepository *repository.Guest) *GuestService {
	return &GuestService{
		guestRepository: guestRepository,
	}
}

func (gs *GuestService) Add(ctx context.Context, guest models.Guest) error {
	if guest == (models.Guest{}) {
		return errors.New("guest is empty")
	}

	if err := gs.guestRepository.Add(ctx, guest); err != nil {
		return err
	}
	return nil
}

func (gs *GuestService) FindByHash(ctx context.Context, hash string) (models.Guest, error) {
	return gs.guestRepository.FindByHash(ctx, hash)
}
