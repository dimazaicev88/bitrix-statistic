package services

import (
	"bitrix-statistic/internal/entitydb"
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
func (as AdvServices) GetAdv(fullUrl string) (entitydb.AdvReferer, error) {
	parse, err := url.Parse(fullUrl)

	if err != nil {
		return entitydb.AdvReferer{}, err
	}
	urlQuery := parse.Query()

	//TODO добавить установку дефолтной рекламной компании, в случае если  не установлена рекламная компания

	return entitydb.AdvReferer{
		Referer1: urlQuery.Get("referer1"),
		Referer2: urlQuery.Get("referer2"),
		Referer3: urlQuery.Get("referer3"),
	}, err
}

func (as AdvServices) FindByPage(direction, page string) (entitydb.AdvDb, error) {
	return as.AdvModel.FindByByPage(page, direction)
}

func (as AdvServices) FindByUuid(advUuid string) (entitydb.AdvDb, error) {
	return as.AdvModel.FindByUuid(advUuid)
}

// DeleteByUuid Удаление рекламной компании по uuid
func (as AdvServices) DeleteByUuid(uuid string) error {
	return as.AdvModel.DeleteByUuid(uuid)
}
