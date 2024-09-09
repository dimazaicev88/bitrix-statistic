package builders

import (
	"bitrix-statistic/internal/filters"
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
	"uuid", "sessionUuid", "advUuid", "dateHit", "phpSessionId", "guestUuid", "newGuest", "userId",
	"userAuth", "url", "url404", "urlFrom", "ip", "method", "cookies", "userAgent", "stopListUuid", "countryId",
	"cityUuid", "siteId",
}

var hitFilterFields = []string{
	"uuid", "guestUuid", "isNewGuest", "sessionUuid", "stopListUuid", "url", "isUrl404", "userId",
	"isRegistered", "date", "ip", "userAgent", "countryId", "country", "cookie", "isStop", "siteId",
}

func (hs *HitSqlBuilder) buildSelect() error {
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

func (hs *HitSqlBuilder) buildWhere() {
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
			} else if op.Operator == ">" {
				hs.sb = hs.sb.Where(squirrel.Gt{op.Field: op.Value})
			} else if op.Operator == ">=" {
				hs.sb = hs.sb.Where(squirrel.GtOrEq{op.Field: op.Value})
			} else if op.Operator == "<" {
				hs.sb = hs.sb.Where(squirrel.Lt{op.Field: op.Value})
			} else if op.Operator == "<=" {
				hs.sb = hs.sb.Where(squirrel.LtOrEq{op.Field: op.Value})
			} else if op.Operator == "like" {
				hs.sb = hs.sb.Where(squirrel.Like{op.Field: op.Value})
			} else if op.Operator == "not like" {
				hs.sb = hs.sb.Where(squirrel.NotLike{op.Field: op.Value})
			} else if op.Operator == "or" {
				hs.sb = hs.sb.Where(squirrel.Or{})
			}
		}
	}
}

func (hs *HitSqlBuilder) buildSkipAndLimit() {
	if hs.filter.Skip < 0 {
		hs.sb = hs.sb.Offset(0)
	} else if hs.filter.Skip > 0 {
		hs.sb = hs.sb.Offset(uint64(hs.filter.Skip))
	}

	if hs.filter.Limit < 0 {
		hs.sb = hs.sb.Limit(0)
	} else if hs.filter.Limit > 1000 {
		hs.sb = hs.sb.Limit(1000)
	}

}

func (hs *HitSqlBuilder) Build() (string, []interface{}, error) {
	for _, value := range hs.filter.Fields {
		if slices.Contains(hitSelectFields, value) == false {
			return "", nil, fmt.Errorf("unknown field: %s", value)
		}
	}

	if err := hs.buildSelect(); err != nil {
		return "", nil, err
	}

	hs.buildWhere()
	hs.buildSkipAndLimit()

	return hs.sb.ToSql()
}
