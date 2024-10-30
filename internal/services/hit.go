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
	hitModel *models.Hit
	ctx      context.Context
}

func NewHit(ctx context.Context, hitModel *models.Hit) *HitService {
	return &HitService{
		ctx:      ctx,
		hitModel: hitModel,
	}
}

func (hs HitService) Add(statData dto.UserData) error {
	if statData == (dto.UserData{}) {
		return errors.New("stat data is empty")
	}

	return hs.hitModel.AddHit(entitydb.Hit{
		Favorites:    statData.IsFavorite,
		PhpSessionId: statData.PHPSessionId,
		GuestUuid:    statData.GuestUuid,
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
	})
}
