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

func NewGuest(ctx context.Context, chClient driver.Conn) *GuestService {
	return &GuestService{
		GuestModel: models.NewGuest(ctx, chClient),
	}
}
