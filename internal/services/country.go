package services

import (
	"bitrix-statistic/internal/entitydb"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/models"
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type CountryServices struct {
	countryModel *models.Country
}

func NewCountry(ctx context.Context, chClient driver.Conn) *CountryServices {
	return &CountryServices{
		countryModel: models.NewCountry(ctx, chClient),
	}
}

func (cs CountryServices) Find(filter filters.Filter) ([]entitydb.CountryDB, error) {
	return cs.countryModel.Find(filter)
}
