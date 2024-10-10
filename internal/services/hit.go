package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
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

func (hs HitService) Find(filter filters.Filter) ([]entitydb.Hit, error) {
	return hs.allModels.Hit.Find(filter)
}

func (hs HitService) FindByUuid(uuid uuid.UUID) (entitydb.Hit, error) {
	return hs.allModels.Hit.FindByUuid(uuid)
}

func (hs HitService) Add(
	hitUuid uuid.UUID,
	isNewGuest bool,
	sessionUuid uuid.UUID,
	advReferer entitydb.AdvCompany,
	statData entityjson.UserData,
) (entitydb.Hit, error) {

	if sessionUuid == uuid.Nil {
		return entitydb.Hit{}, errors.New("sessionUuid is empty")
	}

	if statData == (entityjson.UserData{}) {
		return entitydb.Hit{}, errors.New("stat data is empty")
	}

	hit := entitydb.Hit{
		Uuid:         hitUuid,
		Favorites:    statData.IsFavorite,
		PhpSessionId: statData.PHPSessionId,
		SessionUuid:  sessionUuid,
		AdvUuid:      advReferer.AdvUuid,
		GuestUuid:    statData.GuestUuid,
		IsNewGuest:   isNewGuest,
		UserId:       statData.UserId,
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
func (hs HitService) FindLastHitWithoutSession(guestUuid uuid.UUID, withoutPhpSessionId string) (entitydb.Hit, error) {
	return hs.allModels.Hit.FindLastHitWithoutSession(guestUuid, withoutPhpSessionId)
}

func (hs HitService) FindAll(skip, limit uint32) ([]entitydb.Hit, error) {
	return hs.allModels.Hit.FindAll(skip, limit)
}

// ConvertToJSONHit converts an entitydb.Hit to entityjson.Hit
func (hs HitService) ConvertToJSONHit(dbHit entitydb.Hit) entityjson.Hit {
	return entityjson.Hit{
		Uuid:         dbHit.Uuid,
		SessionUuid:  dbHit.SessionUuid,
		DateHit:      dbHit.DateHit.Format(time.RFC3339), // Format time as needed
		GuestUuid:    dbHit.GuestUuid,
		NewGuest:     dbHit.IsNewGuest,
		UserId:       dbHit.UserId,
		UserAuth:     dbHit.IsUserAuth,
		Url:          dbHit.Url,
		Url404:       dbHit.Url404,
		UrlFrom:      dbHit.UrlFrom,
		Ip:           dbHit.Ip,
		Method:       dbHit.Method,
		Cookies:      dbHit.Cookies,
		UserAgent:    dbHit.UserAgent,
		StopListUuid: dbHit.StopListUuid, // Assuming you want to keep it as uint32
		CountryId:    dbHit.CountryId,
		CountryName:  "", // Set this if you have a way to determine country name
		SiteId:       dbHit.SiteId,
	}
}

func (hs HitService) ConvertToJSONListHits(dbHits []entitydb.Hit) []entityjson.Hit {
	var hits []entityjson.Hit

	for _, dbHit := range dbHits {
		hits = append(hits, hs.ConvertToJSONHit(dbHit))
	}
	return hits
}
