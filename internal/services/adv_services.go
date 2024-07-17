package services

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"net/url"
)

//TODO добавить авто создание рекламной компании

type AdvServices struct {
	AdvModel *models.Adv
}

func NewAdv(ctx context.Context, chClient driver.Conn) *AdvServices {
	return &AdvServices{AdvModel: models.NewAdv(ctx, chClient, models.NewOption(ctx, chClient))}
}

// AutoCreateAdv Автоматическое создание рекламной компании
func (as AdvServices) AutoCreateAdv(fullUrl string) error {
	return nil
}

// GetAdv Получить рекламную компанию
func (as AdvServices) GetAdv(fullUrl string) (entity.AdvReferer, error) {
	parse, err := url.Parse(fullUrl)

	if err != nil {
		return entity.AdvReferer{}, err
	}
	urlQuery := parse.Query()

	//TODO добавить установку дефолтной рекламной компании, в случае если  не установлена рекламная компания

	return entity.AdvReferer{
		Referer1: urlQuery.Get("referer1"),
		Referer2: urlQuery.Get("referer2"),
		Referer3: urlQuery.Get("referer3"),
	}, err
}
