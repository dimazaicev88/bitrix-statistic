package models

import (
	"bitrix-statistic/internal/filters"
)

type OptionModel struct {
}

func (o OptionModel) Add(filter filters.Filter) (interface{}, interface{}) {
	return nil, nil
}
