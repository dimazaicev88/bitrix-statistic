package adv

import (
	"bitrix-statistic/internal/converters"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/utils"
	"fmt"
	"slices"
	"strings"
)

type EventConverter struct {
	filter     filters.Filter
	sqlBuilder *converters.FilterToSqlConverter
}

func NewEventConverter(filter filters.Filter) EventConverter {
	return EventConverter{
		filter:     filter,
		sqlBuilder: converters.NewSqlSQLConverter(),
	}
}

func (hs *EventConverter) buildSelect() error {
	hs.sqlBuilder.AddSql(`SELECT`)

	return nil
}

func (hs *EventConverter) buildWhere() {
	if len(hs.filter.Operators) != 0 {
		hs.sqlBuilder.AddSql(`WHERE`)
		itemsAnd := make([]string, 0, len(hs.filter.Operators))
		itemsOr := make([]string, 0, len(hs.filter.Operators))

		for i := 0; i < len(hs.filter.Operators); i++ {
			op := hs.filter.Operators[i]
			if op.TextOperator == "or" {
				itemsOr = hs.appendSqlWhere(op.Field, op.Operator, op.Value, itemsOr)
			} else if op.TextOperator == "and" || op.TextOperator == "" {
				itemsAnd = hs.appendSqlWhere(op.Field, op.Operator, op.Value, itemsAnd)
			}
		}

		hs.sqlBuilder.AddSql(strings.Join(itemsAnd, ` AND `))
		hs.sqlBuilder.AddSql(strings.Join(itemsOr, ` OR `))
	}
}

func (hs *EventConverter) appendSqlWhere(field, operator string, value any, listSql []string) []string {
	if value != nil {
		if fieldName, ok := advSimpleFields[field]; ok {
			val := utils.StringConcat(fieldName, operator, "?")
			listSql = append(listSql, val)
			hs.sqlBuilder.AddArgs(value)
		} else {
			val := utils.StringConcat(field, operator, "?")
			listSql = append(listSql, val)
			hs.sqlBuilder.AddArgs(value)
		}
	}

	return listSql
}

func (hs *EventConverter) buildOrder() error {
	if len(hs.filter.Order) > 0 {
		hs.sqlBuilder.AddSql("ORDER BY")
		fieldsOrder := make([]string, 0, len(hs.filter.Order))
		for _, fieldName := range hs.filter.Order {
			if slices.Contains(advSelectFields, fieldName) == false {
				return fmt.Errorf("unknown field: %s", fieldName)
			}
			if val, ok := advSimpleFields[fieldName]; ok {
				fieldsOrder = append(fieldsOrder, val)
				continue
			} else {
				fieldsOrder = append(fieldsOrder, fieldName)
			}
		}

		hs.sqlBuilder.AddSql(strings.Join(fieldsOrder, ","))

		if hs.filter.OrderBy != "" {
			hs.sqlBuilder.AddSql(hs.filter.OrderBy)
		} else {
			hs.sqlBuilder.AddSql("DESC")
		}
	}
	return nil
}

func (hs *EventConverter) buildSkipAndLimit() {
	hs.sqlBuilder.AddSql("LIMIT")
	if hs.filter.Skip != 0 {
		hs.sqlBuilder.AddSql("?,").AddArgs(hs.filter.Skip)
	} else {
		hs.sqlBuilder.AddSql("?,").AddArgs(0)
	}

	if hs.filter.Limit != 0 {
		hs.sqlBuilder.AddSql("?").AddArgs(hs.filter.Limit)
	} else if hs.filter.Limit > 1000 || hs.filter.Limit < 0 || hs.filter.Limit == 0 {
		hs.sqlBuilder.AddSql("?").AddArgs(1000)
	}
}

func (hs *EventConverter) Convert() (string, []any, error) {
	if err := hs.buildSelect(); err != nil {
		return "", nil, err
	}

	hs.buildWhere()

	if err := hs.buildOrder(); err != nil {
		return "", nil, err
	}

	hs.buildSkipAndLimit()

	resultSql, args := hs.sqlBuilder.Convert()
	return resultSql, args, nil
}
