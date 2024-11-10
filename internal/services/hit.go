package services

import (
	"bitrix-statistic/internal/dto"
	"bitrix-statistic/internal/models"
	"bitrix-statistic/internal/repository"
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

type HitService struct {
	hitRepository repository.IHitRepository
}

func NewHit(hitRepository repository.IHitRepository) *HitService {
	return &HitService{
		hitRepository: hitRepository,
	}
}

func (hs HitService) Add(ctx context.Context, statData dto.UserData, isNewGuest bool, waitAdd bool) error {
	if statData == (dto.UserData{}) {
		return errors.New("stat data is empty")
	}

	hit := models.Hit{
		Uuid:         uuid.New(),
		PhpSessionId: statData.PHPSessionId,
		DateHit:      time.Now(),
		GuestHash:    statData.GuestHash,
		IsNewGuest:   isNewGuest,
		UserId:       statData.UserId,
		Url:          statData.Url,
		Url404:       statData.IsError404,
		Referer:      statData.Referer,
		Ip:           statData.Ip,
		Method:       statData.Method,
		Cookies:      statData.Cookies,
		UserAgent:    statData.UserAgent,
		SiteId:       statData.SiteId,
		Event1:       statData.Event1,
		Event2:       statData.Event2,
		Event3:       statData.Event3,
	}

	return hs.hitRepository.AddHit(ctx, hit, waitAdd)
}
