package services

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/repository"
	"context"
	"errors"
	"github.com/google/uuid"
)

type HitService struct {
	hitModel *repository.Hit
}

func NewHit(hitModel *repository.Hit) *HitService {
	return &HitService{
		hitModel: hitModel,
	}
}

func (hs HitService) Add(ctx context.Context, statData dto.UserData, isNewGuest bool) error {
	if statData == (dto.UserData{}) {
		return errors.New("stat data is empty")
	}

	hit := models.Hit{
		Uuid:         uuid.New(),
		PhpSessionId: statData.PHPSessionId,
		GuestHash:    statData.GuestHash,
		IsNewGuest:   isNewGuest,
		UserId:       statData.UserId,
		Url:          statData.Url,
		Url404:       statData.IsError404,
		UrlFrom:      statData.Referer,
		Ip:           statData.Ip,
		Method:       statData.Method,
		Cookies:      statData.Cookies,
		UserAgent:    statData.UserAgent,
		CountryId:    "",
		CityId:       "",
		SiteId:       statData.SiteId,
	}
	return hs.hitModel.AddHit(ctx, hit)
}
