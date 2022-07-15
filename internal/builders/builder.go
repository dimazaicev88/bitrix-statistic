package builders

import (
	"bitrix-statistic/internal/filters"
	"strings"
)

//TODO добавить проверку что параметры указаны правильно в where, order by, select,params.

type SQL struct {
	SQL    string
	Params []interface{}
}

type SQLBuilder struct {
	selectBuilder  *strings.Builder
	joinBuilder    *strings.Builder
	whereBuilder   *strings.Builder
	orderByBuilder *strings.Builder
	filter         filters.Filter
	params         *[]interface{}
}

func NewSQLBuilder(filter filters.Filter) SQLBuilder {
	return SQLBuilder{
		selectBuilder:  &strings.Builder{},
		joinBuilder:    &strings.Builder{},
		whereBuilder:   &strings.Builder{},
		orderByBuilder: &strings.Builder{},
		filter:         filter,
		params:         &[]interface{}{},
	}
}
