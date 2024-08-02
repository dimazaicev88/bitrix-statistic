package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
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

func (hs *HitService) FindByUuid(uuid string) (entitydb.Hit, error) {
	return hs.allModels.Hit.FindByUuid(uuid)
}

func (hs *HitService) Add(existsGuest bool, sessionDb entitydb.Session, advReferer entitydb.AdvReferer, statData entityjson.StatData) error {
	return hs.allModels.Hit.AddHit(entitydb.Hit{
		SessionUuid:  sessionDb.Uuid,
		AdvUuid:      advReferer.AdvUuid,
		GuestUuid:    sessionDb.GuestUuid,
		NewGuest:     existsGuest == false,
		UserId:       statData.UserId,
		IsUserAuth:   statData.UserId > 0,
		Url:          statData.Url,
		Url404:       statData.IsError404,
		UrlFrom:      statData.Referer,
		Ip:           statData.Ip,
		Method:       statData.Method,
		Cookies:      statData.Cookies,
		UserAgent:    statData.UserAgent,
		StopListUuid: "",
		CountryId:    "",
		CityUuid:     "",
		SiteId:       "",
	})
}

func (hs *HitService) FindLastHitWithoutSession(guestUuid, withoutPhpSessionId string) (entitydb.Hit, error) {
	return hs.allModels.Hit.FindLastHitWithoutSession(guestUuid, withoutPhpSessionId)
}
