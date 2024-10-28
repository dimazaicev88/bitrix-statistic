package converters

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/utils"
	"fmt"
	"slices"
	"strings"
)

type HitSqlBuilder struct {
	filter     filters.Filter
	sqlBuilder *FilterToSqlConverter
}

func NewHitSQLBuilder(filter filters.Filter) HitSqlBuilder {
	return HitSqlBuilder{
		filter:     filter,
		sqlBuilder: NewSqlSQLConverter(),
	}
}

var hitSelectFields = []string{
	"uuid", "sessionUuid", "advUuid", "dateHit", "phpSessionId", "guestUuid", "newGuest", "userId",
	"userAuth", "url", "url404", "urlFrom", "ip", "method", "cookies", "userAgent", "stopListUuid", "countryId",
	"cityUuid", "siteId",
}

func (hs *HitSqlBuilder) buildSelect() error {
	countFields := 0
	for _, field := range hs.filter.Fields {
		if field == "" {
			continue
		}
		if !slices.Contains(hitSelectFields, field) {
			return fmt.Errorf("unknown field: %s", field)
		}
		countFields++
	}
	if len(hs.filter.Fields) == 0 || countFields == 0 {
		hs.sqlBuilder.AddSql("SELECT * FROM hit ")
	} else {
		hs.sqlBuilder.AddSql(fmt.Sprintf("SELECT %s FROM hit ", strings.Join(hs.filter.Fields, ", ")))
	}
	return nil
}

func (hs *HitSqlBuilder) buildWhere() {
	if len(hs.filter.Operators) != 0 {
		hs.sqlBuilder.AddSql("WHERE ")
		for i := 0; i < len(hs.filter.Operators); i++ {
			op := hs.filter.Operators[i]
			if op.Field == "isRegistered" {
				if op.Value == true {
					hs.sqlBuilder.AddSql("userId>0 ")
				} else {
					hs.sqlBuilder.AddSql(" userId=0 ")
				}
				continue
			}

			if op.Operator == "or" {
				hs.sqlBuilder.AddSql(" OR ")
			} else {
				val := utils.StringConcat(op.Field, op.Operator, "?")
				hs.sqlBuilder.AddSql(val).AddArgs(op.Value)
			}

			if i+1 < len(hs.filter.Operators)-1 {
				if hs.filter.Operators[i+1].Operator != "or" || (i-1 > 0 && hs.filter.Operators[i-1].Operator != "or") {
					hs.sqlBuilder.AddSql("AND")
				}
			}
		}
	}
}

func (hs *HitSqlBuilder) buildSkipAndLimit() {
	hs.sqlBuilder.AddSql("LIMIT")
	if hs.filter.Skip != 0 {
		hs.sqlBuilder.AddSql("?,").AddArgs(hs.filter.Skip)
	} else {
		hs.sqlBuilder.AddSql("?,").AddArgs(hs.filter.Skip)
	}

	if hs.filter.Limit != 0 {
		hs.sqlBuilder.AddSql("?").AddArgs(hs.filter.Limit)
	} else if hs.filter.Limit > 1000 || hs.filter.Limit < 0 || hs.filter.Limit == 0 {
		hs.sqlBuilder.AddSql("?").AddArgs(hs.filter.Limit)
	}
}

func (hs *HitSqlBuilder) Build() (string, []any, error) {
	if err := hs.buildSelect(); err != nil {
		return "", nil, err
	}

	hs.buildWhere()
	hs.buildSkipAndLimit()

	resultSql, args := hs.sqlBuilder.Convert()
	return resultSql, args, nil
}
