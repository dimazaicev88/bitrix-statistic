package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/utils"
	"context"
	"github.com/google/uuid"
)

type GuestService struct {
	allModels *models.Models
	ctx       context.Context
}

func NewGuest(ctx context.Context, allModels *models.Models) *GuestService {
	return &GuestService{
		ctx:       ctx,
		allModels: allModels,
	}
}

func (gs GuestService) AddGuest(statData entityjson.StatData) error {
	guestHash, err := utils.GetGuestMd5(statData)
	if err != nil {
		return err
	}

	guest := entitydb.Guest{
		GuestHash:     guestHash,
		UserAgent:     statData.UserAgent,
		Ip:            statData.Ip,
		XForwardedFor: statData.HttpXForwardedFor,
	}

	if err := gs.allModels.Guest.Add(guest); err != nil {
		return err
	}
	return nil
}

func (gs GuestService) Find(filter filters.Filter) ([]entitydb.GuestStat, error) {
	return gs.allModels.Guest.Find(filter)
}

func (gs GuestService) FindByUuid(uuid uuid.UUID) (entitydb.GuestStat, error) {
	return gs.allModels.Guest.FindByUuid(uuid)
}

func (gs GuestService) ExistsGuestByHash(hash string) (bool, error) {
	return gs.allModels.Guest.ExistsByHash(hash)
}
