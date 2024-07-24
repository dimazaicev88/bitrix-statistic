package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/entityjson"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
	"time"
)

type GuestService struct {
	guestModel *models.Guest
}

func NewGuest(ctx context.Context, chClient driver.Conn) *GuestService {
	return &GuestService{
		guestModel: models.NewGuest(ctx, chClient),
	}
}

func (gs GuestService) AddGuest(statData entityjson.StatData, adv entitydb.AdvReferer) error {
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

	if err := gs.guestModel.AddGuest(guest); err != nil {
		return err
	}
	return nil
}

func (gs GuestService) Find(filter filters.Filter) ([]entitydb.GuestDb, error) {
	return gs.guestModel.Find(filter)
}

func (gs GuestService) FindByUuid(uuid uuid.UUID) (entitydb.GuestDb, error) {
	return gs.guestModel.FindByUuid(uuid)
}
