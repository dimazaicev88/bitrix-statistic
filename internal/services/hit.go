package services

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/entitydb"
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

func (hs HitService) FindByUuid(uuid uuid.UUID) (entitydb.Hit, error) {
	return hs.allModels.Hit.FindByUuid(uuid)
}

func (hs HitService) Add(
	hitUuid uuid.UUID,
	isNewGuest bool,
	sessionUuid uuid.UUID,
	statData dto.UserData,
) (entitydb.Hit, error) {

	if sessionUuid == uuid.Nil {
		return entitydb.Hit{}, errors.New("sessionUuid is empty")
	}

	if statData == (dto.UserData{}) {
		return entitydb.Hit{}, errors.New("stat data is empty")
	}

	hit := entitydb.Hit{
		Uuid:         hitUuid,
		Favorites:    statData.IsFavorite,
		PhpSessionId: statData.PHPSessionId,
		SessionUuid:  sessionUuid,
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
