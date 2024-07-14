package services

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"time"
)

type GuestService struct {
	GuestModel *models.GuestModel
}

func (s GuestService) AddGuest(statData entity.StatData, adv entity.AdvReferer) error {
	guest := entity.GuestDb{
		Timestamp:     time.Now(),
		FirstUrlFrom:  statData.Referer,
		FirstUrlTo:    statData.Url,
		FirstUrlTo404: statData.IsError404,
		FirstSiteId:   statData.SiteId,
		FirstAdvUuid:  "",
		FirstReferer1: adv.Referer1,
		FirstReferer2: adv.Referer2,
		FirstReferer3: adv.Referer3,
		GuestHash:     statData.GuestHash,
	}

	if err := s.GuestModel.AddGuest(guest); err != nil {
		return err
	}
	return nil
}

func NewGuestService(ctx context.Context, chClient driver.Conn) *GuestService {
	return &GuestService{
		GuestModel: models.NewGuestModel(ctx, chClient),
	}
}
