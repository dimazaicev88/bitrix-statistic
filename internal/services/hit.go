package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
	"errors"
	"github.com/google/uuid"
)

type HitService struct {
	allModels *models.Models
	ctx       context.Context
}

func NewHit(ctx context.Context, allModels *models.Models) *HitService {
	return &HitService{
		ctx:       ctx,
		allModels: allModels,
	}
}

func (hs *HitService) Find(filter filters.Filter) ([]entitydb.Hit, error) {
	return hs.allModels.Hit.Find(filter)
}

func (hs *HitService) FindByUuid(uuid uuid.UUID) (entitydb.Hit, error) {
	return hs.allModels.Hit.FindByUuid(uuid)
}

func (hs *HitService) Add(existsGuest bool, sessionDb entitydb.Session, advReferer entitydb.AdvReferer, statData entityjson.UserData) (entitydb.Hit, error) {

	if sessionDb == (entitydb.Session{}) {
		return entitydb.Hit{}, errors.New("session is empty")
	}

	if statData == (entityjson.UserData{}) {
		return entitydb.Hit{}, errors.New("stat data is empty")
	}

	hit := entitydb.Hit{
		Uuid:         uuid.New(),
		PhpSessionId: statData.PHPSessionId,
		SessionUuid:  sessionDb.Uuid,
		AdvUuid:      advReferer.AdvUuid,
		GuestUuid:    sessionDb.GuestUuid,
		IsNewGuest:   existsGuest == false,
		UserId:       uint32(statData.UserId),
		IsUserAuth:   statData.UserId > 0,
		Url:          statData.Url,
		Url404:       statData.IsError404,
		UrlFrom:      statData.Referer,
		Ip:           statData.Ip,
		Method:       statData.Method,
		Cookies:      statData.Cookies,
		UserAgent:    statData.UserAgent,
		StopListUuid: uuid.New(),
		CountryId:    "",
		CityUuid:     uuid.New(),
		SiteId:       statData.SiteId,
	}
	if err := hs.allModels.Hit.AddHit(hit); err != nil {
		return entitydb.Hit{}, err
	}
	return hit, nil
}

// FindLastHitWithoutSession Найти хит, не включая указанную сессию
func (hs *HitService) FindLastHitWithoutSession(guestUuid uuid.UUID, withoutPhpSessionId string) (entitydb.Hit, error) {
	return hs.allModels.Hit.FindLastHitWithoutSession(guestUuid, withoutPhpSessionId)
}
