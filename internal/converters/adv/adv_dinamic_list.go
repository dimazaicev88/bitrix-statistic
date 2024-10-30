package adv

import (
	"bitrix-statistic/internal/converters"
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/utils"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"slices"
	"strings"
)

var advDynamicListSelectFields = []string{
	"dateStat",      // приоритет
	"day",           // идентификатор referer1
	"month",         // идентификатор referer2
	"year",          // описание
	"guests",        // посетителей на прямом заходе
	"newGuests",     // новых посетителей на прямом заходе
	"favorites",     // посетителей, добавивших сайт в "Избранное" на прямом заходе
	"hosts",         // хостов на прямом заходе
	"sessions",      // сессий на прямом заходе
	"hits",          // хитов на прямом заходе
	"guestsBack",    // посетителей на возврате
	"favoritesBack", // посетителей, добавивших сайт в "Избранное" на возврате
	"hostsBack",     // хостов на возврате
	"sessionsBack",  // сессий на возврате
	"hitsBack",      // хитов на возврате
}

type DynamicListConverter struct {
	filter        filters.Filter
	sqlBuilder    *converters.FilterToSqlConverter
	groupByFields mapset.Set[string]
}

func NewDynamicListConverter(filter filters.Filter, maxMin bool) DynamicListConverter {
	return DynamicListConverter{
		filter:        filter,
		sqlBuilder:    converters.NewSqlSQLConverter(),
		groupByFields: mapset.NewSet[string](),
	}
}

// TODO добавить сборку когда нету выбираемых полей
func (hs *DynamicListConverter) buildSelectAndGroupBy() error {

	hs.sqlBuilder.AddSql(`SELECT`)
	tmpListFields := make([]string, 0, len(hs.filter.Fields))
	for _, fieldName := range hs.filter.Fields {
		if fieldName == "" {
			continue
		}
		if !slices.Contains(advDynamicListSelectFields, fieldName) {
			return fmt.Errorf("unknown field: %s", fieldName)
		}
	}

	if len(tmpListFields) > 0 {
		hs.sqlBuilder.AddSql(strings.Join(tmpListFields, ","))
	} else {
		hs.sqlBuilder.AddSql("*")
	}

	hs.sqlBuilder.AddSql(`FROM adv_day`)

	return nil
}

func (hs *DynamicListConverter) BuildWhere() {
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

func (hs *DynamicListConverter) appendSqlWhere(field, operator string, value any, listSql []string) []string {
	if value != nil {
		//if !slices.Contains(advSelectFields, field) { //TODO заменить проверкой на фильтруемые поля
		val := utils.StringConcat(field, operator, "?")
		listSql = append(listSql, val)
		hs.sqlBuilder.AddArgs(value)
		//}
	}

	return listSql
}

func (hs *DynamicListConverter) buildOrder() error {
	if len(hs.filter.Order) > 0 {
		hs.sqlBuilder.AddSql("ORDER BY")
		fieldsOrder := make([]string, 0, len(hs.filter.Order))
		for _, fieldName := range hs.filter.Order {
			if slices.Contains(advDynamicListSelectFields, fieldName) == false {
				return fmt.Errorf("unknown field: %s", fieldName)
			}
			fieldsOrder = append(fieldsOrder, fieldName)
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

func (hs *DynamicListConverter) buildSkipAndLimit() {
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

func (hs *DynamicListConverter) Convert() (string, string, []any, error) {
	if err := hs.buildSelectAndGroupBy(); err != nil {
		return "", "", nil, err
	}

	hs.BuildWhere()
	if hs.groupByFields.IsEmpty() == false {
		hs.sqlBuilder.AddSql("GROUP BY")
		listGroupByFields := hs.groupByFields.ToSlice()
		slices.Sort(listGroupByFields)
		hs.sqlBuilder.AddSql(strings.Join(listGroupByFields, ","))
	}

	if err := hs.buildOrder(); err != nil {
		return "", "", nil, err
	}

	hs.buildSkipAndLimit()

	resultSql, args := hs.sqlBuilder.Convert()
	return resultSql, "", args, nil
}

//TODO добавить
//sqlMaxMin := `
//				SELECT
//					max(dateStat) AS dateLast,
//					min(dateStat) AS dateFirst,
//					dayOfMonth(max(dateStat)) AS maxDay,
//					month(max(dateStat)) AS maxMonth,
//					year(max(dateStat)) AS maxYear,
//					dayOfMonth(min(dateStat)) AS minDay,
//					month(min(dateStat)) AS minMonth,
//					year(min(dateStat)) AS minYear
//				FROM
//					adv_day
//			`
