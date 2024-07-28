package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/google/uuid"
	"time"
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

func (gs GuestService) AddGuest(statData entityjson.StatData, adv entitydb.AdvReferer) error {
	guest := entitydb.Guest{
		Timestamp: time.Now(),
		UrlFrom:   statData.Referer,
		UrlTo:     statData.Url,
		UrlTo404:  statData.IsError404,
		SiteId:    statData.SiteId,
		AdvUuid:   "",
		Referer1:  adv.Referer1,
		Referer2:  adv.Referer2,
		Referer3:  adv.Referer3,
		GuestHash: statData.GuestHash,
	}

	if err := gs.allModels.Guest.AddGuest(guest); err != nil {
		return err
	}
	return nil
}

func (gs GuestService) Find(filter filters.Filter) ([]entitydb.Guest, error) {
	return gs.allModels.Guest.Find(filter)
}

func (gs GuestService) FindByUuid(uuid uuid.UUID) (entitydb.Guest, error) {
	return gs.allModels.Guest.FindByUuid(uuid)
}

func (gs GuestService) ExistsGuestByHash(hash string) (bool, error) {
	return gs.allModels.Guest.ExistsGuestByHash(hash)
}
