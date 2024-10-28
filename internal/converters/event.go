package converters

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/utils"
	"fmt"
	"slices"
	"strings"
)

var eventSelectFields = []string{
	"uuid",
	"uuidEventType",
	"eventName",
	"event1",
	"event2",
	"event3",
	"date",
	"money",
	"currency",
	"sessionUuid",
	"guestUuid",
	"advUud",
	"advBack",
	"hitUuid",
	"countryId",
	"country",
	"refererUrl",
	"refererSiteUuid",
	"url",
	"siteId",
	"redirectUrl",
}

type EventSqlBuilder struct {
	filter     filters.Filter
	sqlBuilder *FilterToSqlConverter
}

func NewEventSqlBuilder(filter filters.Filter) EventSqlBuilder {
	return EventSqlBuilder{
		filter:     filter,
		sqlBuilder: NewSqlSQLConverter(),
	}
}

func (hs *EventSqlBuilder) buildSelect() error {
	countFields := 0
	for _, field := range hs.filter.Fields {
		if field == "" {
			continue
		}
		if !slices.Contains(eventSelectFields, field) {
			return fmt.Errorf("unknown field: %s", field)
		}
		countFields++
	}
	if len(hs.filter.Fields) == 0 || countFields == 0 {
		hs.sqlBuilder.AddSql("SELECT * FROM adv")
	} else {
		hs.sqlBuilder.AddSql(fmt.Sprintf("SELECT %s FROM hit ", strings.Join(hs.filter.Fields, ", ")))
	}
	return nil
}

func (hs *EventSqlBuilder) buildWhere() {
	if len(hs.filter.Operators) != 0 {
		hs.sqlBuilder.AddSql("WHERE")
		for i := 0; i < len(hs.filter.Operators); i++ {
			op := hs.filter.Operators[i]
			if op.Field == "isRegistered" {
				if op.Value == true {
					hs.sqlBuilder.AddSql("userId>0")
				} else {
					hs.sqlBuilder.AddSql("userId=0")
				}
				continue
			}

			if op.Operator == "or" {
				hs.sqlBuilder.AddSql("OR")
			} else {
				val := utils.StringConcat(op.Field, op.Operator, "?")
				hs.sqlBuilder.AddSql(val).AddArgs(op.Value)
			}

			if i+1 < len(hs.filter.Operators)-1 {
				if hs.filter.Operators[i+1].Operator != "or" || (i-1 > 0 && hs.filter.Operators[i-1].Operator != "or") {
					hs.sqlBuilder.AddSql(" AND ")
				}
			}
		}
	}
}

func (hs *EventSqlBuilder) buildSkipAndLimit() {
	hs.sqlBuilder.AddSql("LIMIT")
	if hs.filter.Skip != 0 {
		hs.sqlBuilder.AddSql("?,").AddArgs(hs.filter.Skip)
	} else {
		hs.sqlBuilder.AddSql("?,").AddArgs(0)
	}

	if hs.filter.Limit != 0 {
		hs.sqlBuilder.AddSql("?").AddArgs(0)
	} else if hs.filter.Limit > 1000 || hs.filter.Limit < 0 || hs.filter.Limit == 0 {
		hs.sqlBuilder.AddSql("?").AddArgs(1000)
	}
}

func (hs *EventSqlBuilder) Build() (string, []any, error) {
	if err := hs.buildSelect(); err != nil {
		return "", nil, err
	}

	hs.buildWhere()
	hs.buildSkipAndLimit()

	resultSql, args := hs.sqlBuilder.Convert()
	return resultSql, args, nil
}
