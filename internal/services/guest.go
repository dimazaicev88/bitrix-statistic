package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/maypok86/otter"
	"github.com/sirupsen/logrus"
	"time"
)

type GuestService struct {
	allModels   *models.Models
	ctx         context.Context
	cacheGuest  otter.Cache[string, entitydb.Guest]
	hitService  HitService
	advServices AdvServices
}

func NewGuest(ctx context.Context, allModels *models.Models, hitService HitService, advServices AdvServices) *GuestService {
	otterCache, err := otter.MustBuilder[string, entitydb.Guest](15000).
		CollectStats().
		WithTTL(time.Minute * 15).
		Build()

	if err != nil {
		logrus.Fatal(err)
	}

	return &GuestService{
		ctx:         ctx,
		allModels:   allModels,
		cacheGuest:  otterCache,
		hitService:  hitService,
		advServices: advServices,
	}
}

func (gs GuestService) AddGuest(statData entityjson.StatData, advReferer entitydb.AdvReferer) (string, error) {
	newUUID, err := uuid.NewUUID()
	var guestDb entitydb.Guest

	if err != nil {
		return "", err
	}

	guestDb.Uuid = newUUID.String()
	guestDb.FirstDate = time.Now()
	guestDb.PhpSessionId = statData.PHPSessionId
	guestDb.FirstUrlFrom = statData.Referer
	guestDb.FirstUrlTo = statData.Url
	guestDb.FirstUrlTo404 = statData.IsError404
	guestDb.FirstSiteId = statData.SiteId
	guestDb.FirstAdvUuid = advReferer.AdvUuid
	guestDb.FirstReferer1 = advReferer.Referer1
	guestDb.FirstReferer2 = advReferer.Referer2
	guestDb.FirstReferer3 = advReferer.Referer3
	guestDb.Sessions = 1
	guestDb.Sign = 1
	guestDb.Version = 1
	gs.cacheGuest.Set(newUUID.String(), guestDb)
	if err = gs.allModels.Guest.Add(guestDb); err != nil {
		return "", err
	}

	return newUUID.String(), nil
}

func (gs GuestService) Find(filter filters.Filter) ([]entitydb.Guest, error) {
	return gs.allModels.Guest.Find(filter)
}

func (gs GuestService) FindByUuid(uuid string) (entitydb.Guest, error) {
	return gs.allModels.Guest.FindByUuid(uuid)
}

func (gs GuestService) ExistsGuestByHash(hash string) (bool, error) {
	return gs.allModels.Guest.ExistsByHash(hash)
}

func (gs GuestService) FindByHash(hash string) (entitydb.Guest, error) {
	return gs.allModels.Guest.FindByHash(hash)
}

func (gs GuestService) UpdateGuest(guestDb entitydb.Guest, statData entityjson.StatData, referer entitydb.AdvReferer) error {
	var newGuestDbValue entitydb.Guest
	oldGuestDbValue, err := gs.FindByUuid(guestDb.Uuid)
	if err != nil {
		return err
	}

	//Если это новая сессия увеличиваем счетчик сессий
	if oldGuestDbValue.PhpSessionId != statData.PHPSessionId {
		newGuestDbValue.Sessions += oldGuestDbValue.Sessions
	}

	newGuestDbValue.FirstAdvUuid = referer.AdvUuid
	newGuestDbValue.LastAdvUUid = referer.AdvUuid
	newGuestDbValue.LastAdvBack = referer.LastAdvBack
	newGuestDbValue.FirstReferer1 = referer.Referer1
	newGuestDbValue.FirstReferer2 = referer.Referer2
	newGuestDbValue.LastReferer1 = referer.Referer1
	newGuestDbValue.LastReferer2 = referer.Referer2

	err = gs.allModels.Guest.Update(oldGuestDbValue, newGuestDbValue)
	if err != nil {
		return err
	}

	return nil
}
