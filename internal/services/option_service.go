package services

import (
	"bitrix-statistic/internal/entity"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/maypok86/otter"
	"time"
)

type OptionService struct {
	optionModel *models.Option
}

func NewOption(ctx context.Context, chClient driver.Conn) *OptionService {

	cache, err := otter.MustBuilder[string, string](10_000).
		CollectStats().
		Cost(func(key string, value string) uint32 {
			return 1
		}).
		WithTTL(time.Hour).
		Build()

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
