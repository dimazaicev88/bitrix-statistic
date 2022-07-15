package builders

import (
	"bitrix-statistic/internal/filters"
)

type CitySQLBuilder struct {
	SQLBuilder
}

func NewCityBuilder(filter filters.Filter) CitySQLBuilder {
	return CitySQLBuilder{NewSQLBuilder(filter)}
}
