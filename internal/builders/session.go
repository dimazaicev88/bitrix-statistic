package builders

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/utils"
	"fmt"
	"slices"
	"strings"
)

var sessionSelectFields = []string{
	"uuid",
	"guestUuid",
	"isNewGuest",
	"userId",
	"isUserAuth",
	"isRegistered",
	"isFavorites",
	"events",
	"hits",
	"isAdv",
	"advUuid",
	"isAdvBack",
	"referer1",
	"referer2",
	"referer3",
	"stop",
	"stopListId",
	"countryId",
	"country",
	"ip",
	"userAgent",
	"date",
	"urlTo",
	"isUrlTo404",
	"firstSiteId",
	"urlLast",
	"isUrlLast404",
	"lastSiteId",
}

type SessionSqlBuilder struct {
	filter     filters.Filter
	sqlBuilder *SqlBuilder
}

func NewSessionSqlBuilder(filter filters.Filter) SessionSqlBuilder {
	return SessionSqlBuilder{
		filter:     filter,
		sqlBuilder: NewSqlBuilder(),
	}
}

func (hs *SessionSqlBuilder) buildSelect() error {
	countFields := 0
	for _, field := range hs.filter.Fields {
		if field == "" {
			continue
		}
		if !slices.Contains(sessionSelectFields, field) {
			return fmt.Errorf("unknown field: %s", field)
		}
		countFields++
	}
	if len(hs.filter.Fields) == 0 || countFields == 0 {
		hs.sqlBuilder.Add("SELECT * FROM adv")
	} else {
		hs.sqlBuilder.Add(fmt.Sprintf("SELECT %s FROM hit ", strings.Join(hs.filter.Fields, ", ")))
	}
	return nil
}

func (hs *SessionSqlBuilder) buildWhere() {
	if len(hs.filter.Operators) != 0 {
		hs.sqlBuilder.Add("WHERE ")
		for i := 0; i < len(hs.filter.Operators); i++ {
			op := hs.filter.Operators[i]
			if op.Field == "isRegistered" {
				if op.Value == true {
					hs.sqlBuilder.Add("userId>0 ")
				} else {
					hs.sqlBuilder.Add(" userId=0 ")
				}
				continue
			}

			if op.Operator == "or" {
				hs.sqlBuilder.Add(" OR ")
			} else {
				val := utils.StringConcat(op.Field, op.Operator, "?")
				hs.sqlBuilder.Add(val, op.Value)
			}

			if i+1 < len(hs.filter.Operators)-1 {
				if hs.filter.Operators[i+1].Operator != "or" || (i-1 > 0 && hs.filter.Operators[i-1].Operator != "or") {
					hs.sqlBuilder.Add(" AND ")
				}
			}
		}
	}
}

func (hs *SessionSqlBuilder) buildSkipAndLimit() {
	hs.sqlBuilder.Add(" LIMIT ")
	if hs.filter.Skip != 0 {
		hs.sqlBuilder.Add("?, ", hs.filter.Skip)
	} else {
		hs.sqlBuilder.Add("?, ", 0)
	}

	if hs.filter.Limit != 0 {
		hs.sqlBuilder.Add("?", 0)
	} else if hs.filter.Limit > 1000 || hs.filter.Limit < 0 || hs.filter.Limit == 0 {
		hs.sqlBuilder.Add("?", 1000)
	}
}

func (hs *SessionSqlBuilder) Build() (string, []any, error) {
	if err := hs.buildSelect(); err != nil {
		return "", nil, err
	}

	hs.buildWhere()
	hs.buildSkipAndLimit()

	resultSql, args := hs.sqlBuilder.Build()
	return resultSql, args, nil
}
