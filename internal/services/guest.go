package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
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

func NewGuest(ctx context.Context, allModels *models.Models, hitService *HitService, advServices *AdvServices) *GuestService {
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

func (gs GuestService) Add(userData entityjson.UserData, advReferer entitydb.AdvReferer) (entitydb.Guest, error) {

	if userData == (entityjson.UserData{}) {
		return entitydb.Guest{}, errors.New("user data is empty")
	}

	guestDb := entitydb.Guest{
		Uuid:           userData.GuestUuid,
		FirstDate:      time.Now(),
		PhpSessionId:   userData.PHPSessionId,
		FirstUrlFrom:   userData.Referer,
		FirstUrlTo:     userData.Url,
		FirstUrlTo404:  userData.IsError404,
		FirstSiteId:    userData.SiteId,
		FirstAdvUuid:   advReferer.AdvUuid,
		FirstReferer1:  advReferer.Referer1,
		FirstReferer2:  advReferer.Referer2,
		FirstReferer3:  advReferer.Referer3,
		LastIp:         userData.Ip,
		LastUserId:     userData.UserId,
		LastUserAuth:   userData.UserId > 0,
		LastUrlLast:    userData.Url,
		LastUrlLast404: userData.IsError404,
		LastUserAgent:  userData.UserAgent,
		LastCookie:     userData.Cookies,
		LastAdvUUid:    advReferer.AdvUuid,
		LastAdvBack:    advReferer.LastAdvBack,
		LastReferer1:   advReferer.Referer1,
		LastReferer2:   advReferer.Referer2,
		LastReferer3:   advReferer.Referer3,
		LastSiteId:     userData.SiteId,
		Hits:           1,
		Sessions:       1,
		Sign:           1,
		Version:        1,
	}
	gs.cacheGuest.Set(guestDb.Uuid.String(), guestDb)
	if err := gs.allModels.Guest.Add(guestDb); err != nil {
		return entitydb.Guest{}, err
	}
	return guestDb, nil
}

func (gs GuestService) Find(filter filters.Filter) ([]entitydb.Guest, error) {
	return gs.allModels.Guest.Find(filter)
}

func (gs GuestService) FindByUuid(uuid uuid.UUID) (entitydb.Guest, error) {
	return gs.allModels.Guest.FindByUuid(uuid)
}

func (gs GuestService) UpdateGuest(oldGuest, newGuestDb entitydb.Guest) error {
	oldGuest.Sign *= -1
	newGuestDb.Sign *= 1
	newGuestDb.Version += 1
	return gs.allModels.Guest.Update(oldGuest, newGuestDb)
}
