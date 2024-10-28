package services

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/maypok86/otter"
	"github.com/sirupsen/logrus"
	"time"
)

type GuestService struct {
	allModels   *models.Models
	ctx         context.Context
	cacheGuest  otter.Cache[string, entitydb.Guest]
	hitService  *HitService
	advServices *AdvServices
}

func NewGuest(ctx context.Context, allModels *models.Models) *GuestService {
	otterCache, err := otter.MustBuilder[string, entitydb.Guest](15000).
		CollectStats().
		WithTTL(time.Minute * 15).
		Build()

	if err != nil {
		logrus.Fatal(err)
	}
	return &GuestService{
		ctx:        ctx,
		allModels:  allModels,
		cacheGuest: otterCache,
	}
}

func (gs *GuestService) SetHitService(hitService *HitService) {
	gs.hitService = hitService
}

func (gs *GuestService) SetAdvService(advServices *AdvServices) {
	gs.advServices = advServices
}

func (gs *GuestService) Add(userData dto.UserData) (entitydb.Guest, error) {

	if userData == (dto.UserData{}) {
		return entitydb.Guest{}, errors.New("user data is empty")
	}

	guestDb := entitydb.Guest{
		Uuid:    userData.GuestUuid,
		DateAdd: time.Now(),
		//TODO добавить repair
	}
	if err := gs.allModels.Guest.Add(guestDb); err != nil {
		return entitydb.Guest{}, err
	}

	gs.cacheGuest.Set(guestDb.Uuid.String(), guestDb)
	return guestDb, nil
}

func (gs *GuestService) Find(filter filters.Filter) ([]entitydb.Guest, error) {
	return gs.allModels.Guest.Find(filter)
}

func (gs *GuestService) FindByUuid(uuid uuid.UUID) (entitydb.Guest, error) {
	return gs.allModels.Guest.FindByUuid(uuid)
}

func (gs *GuestService) ClearCache() {
	gs.cacheGuest.Close()
}

func (gs *GuestService) FindAll(skip uint32, limit uint32) ([]entitydb.Guest, error) {
	return nil, nil
}

func (gs *GuestService) ConvertToJSONListGuest(guests []entitydb.Guest) ([]entitydb.Guest, error) {
	return nil, nil
}
