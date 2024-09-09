package builders

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/utils"
	"fmt"
	"github.com/Masterminds/squirrel"
	"slices"
)

type HitSqlBuilder struct {
	filter filters.Filter
	sb     squirrel.SelectBuilder
}

func NewHitSQLBuilder(filter filters.Filter) HitSqlBuilder {
	return HitSqlBuilder{
		filter: filter,
	}
}

var hitSelectFields = []string{
	"uuid", "session_uuid", "advUuid", "dateHit", "phpSessionId", "guestUuid", "newGuest", "userId",
	"userAuth", "url", "url404", "urlFrom", "ip", "method", "cookies", "userAgent", "stopListUuid", "countryId",
	"cityUuid", "siteId",
}

var hitFilterFields = []string{
	"uuid", "guestUuid", "isNewGuest", "sessionUuid", "stopListUuid", "url", "isUrl404", "userId",
	"isRegistered", "date", "ip", "userAgent", "countryId", "country", "cookie", "isStop", "siteId",
}

func (hs HitSqlBuilder) buildSelect() error {
	for _, field := range hs.filter.Fields {
		if !slices.Contains(hitSelectFields, field) {
			return fmt.Errorf("unknown field: %s", field)
		}
	}
	if len(hs.filter.Fields) == 0 {
		hs.sb = squirrel.Select("*")
	} else {
		hs.sb = squirrel.Select(hs.filter.Fields...)
	}
	hs.sb = hs.sb.From("hit")
	return nil
}

func (hs HitSqlBuilder) buildWhere() {
	if len(hs.filter.Operators) == 0 {
		hs.sb = hs.sb.Where("")
	} else {
		for _, op := range hs.filter.Operators {
			if op.Field == "isRegistered" {
				if op.Value == true {
					hs.sb = hs.sb.Where(squirrel.Gt{"userId": 0})
				} else {
					hs.sb = hs.sb.Where(squirrel.Eq{"userId": 0})
				}
				continue
			}

			if op.Operator == "=" {
				hs.sb = hs.sb.Where(squirrel.Eq{op.Field: op.Value})
			} else if op.Operator == "!=" {
				hs.sb = hs.sb.Where(squirrel.NotEq{op.Field: op.Value})
			}

		}
	}
}

func (hs HitSqlBuilder) Build() (string, []interface{}, error) {
	for _, value := range hs.filter.Fields {
		if slices.Contains(hitSelectFields, value) == false {
			return "", nil, fmt.Errorf("unknown field: %s", value)
		}
	}

	var allArgs []interface{}
	sqlSelect, err := BuildSelect(hs.filter.Fields, "hit")
	if err != nil {
		return "", nil, err
	}

	sqlWhere, whereArgs, err := BuildWhereSQL(hs.filter)
	if err != nil {
		return "", nil, err
	}
	allArgs = append(allArgs, whereArgs)

	limitSql, limitArgs := BuildLimit(hs.filter)
	allArgs = append(allArgs, limitArgs)
	return utils.StringConcat(sqlSelect, sqlWhere, limitSql), allArgs, nil
}
