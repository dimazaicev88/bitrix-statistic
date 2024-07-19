package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"time"
)

type GuestService struct {
	GuestModel *models.Guest
}

func (s GuestService) AddGuest(statData entityjson.StatData, adv entitydb.AdvReferer) error {
	guest := entitydb.GuestDb{
		Timestamp: time.Now(),
		UrlFrom:   statData.Referer,
		UrlTo:     statData.Url,
		UrlTo404:  statData.IsError404,
		SiteId:    statData.SiteId,
		AdvUuid:   "",
		Referer1:  adv.Referer1,
		Referer2:  adv.Referer2,
		Referer3:  adv.Referer3,
		GuestHash: statData.GuestHash,
	}

	if err := s.GuestModel.AddGuest(guest); err != nil {
		return err
	}
	return nil
}

func NewGuest(ctx context.Context, chClient driver.Conn) *GuestService {
	return &GuestService{
		GuestModel: models.NewGuest(ctx, chClient),
	}
}
