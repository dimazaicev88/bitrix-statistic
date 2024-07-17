package services

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type OptionService struct {
	optionModel *models.Option
}

func NewOption(ctx context.Context, chClient driver.Conn) *OptionService {
	return &OptionService{
		optionModel: models.NewOption(ctx, chClient),
	}
}

func (o OptionService) Add(options []entity.Option) error {
	return nil
}

func (o OptionService) Set(key, value, desc, site string) error {
	return nil
}

func (o OptionService) Get(name string) string {
	return ""
}
