package builders

import (
	"bitrix-statistic/internal/filters"
	"regexp"
	"strings"
	"unicode/utf8"
)

type HitSqlBuilder struct {
	SQLBuilder
}

func NewHitSQLBuilder(filter filters.Filter) HitSqlBuilder {
	return HitSqlBuilder{NewSQLBuilder(filter)}
}

var hitFields = map[string]string{
	"ID":                 " t1.ID",
	"SESSION_ID":         " t1.SESSION_ID ",
	"GUEST_ID":           " t1.GUEST_ID ",
	"NEW_GUEST":          " t1.NEW_GUEST ",
	"USER_ID":            " t1.USER_ID ",
	"USER_AUTH":          " t1.USER_AUTH ",
	"URL":                " t1.URL ",
	"URL_404":            " t1.URL_404 ",
	"URL_FROM":           " t1.URL_FROM ",
	"IP":                 " t1.IP ",
	"METHOD":             " t1.METHOD ",
	"COOKIES":            " t1.COOKIES ",
	"USER_AGENT":         " t1.USER_AGENT ",
	"STOP_LIST_ID":       " t1.STOP_LIST_ID ",
	"COUNTRY_ID":         " t1.COUNTRY_ID ",
	"CITY_ID":            " t1.CITY_ID ",
	"REGION REGION_NAME": " t3.REGION ",
	"USER":               " t2.LOGIN, t2.NAME ",
	"NAME CITY_NAME":     " t3.CITY_NAME ",
	"SITE_ID":            " t1.SITE_ID ",
}

func (hs HitSqlBuilder) buildSelectAndJoin() HitSqlBuilder {
	var selectFields []string
	hs.selectBuilder.WriteString("SELECT ")
	if len(hs.filter.Select) == 0 {
		hs.selectBuilder.WriteString("* ")
	} else {
		for _, selectField := range hs.filter.Select {
			if value, ok := hitFields[selectField]; ok {
				selectFields = append(selectFields, value)
			}
			if selectField == "USER" {
				hs.selectBuilder.WriteString("")
				hs.joinBuilder.WriteString(" LEFT JOIN b_user t2 ON (t2.ID = t1.USER_ID)")
			}
			if selectField == "COUNTRY" {
				hs.joinBuilder.WriteString(" INNER JOIN b_stat_country t3 ON (t3.ID = t1.COUNTRY_ID)")
			}
		}
	}
	hs.selectBuilder.WriteString(strings.Join(selectFields, ","))
	hs.selectBuilder.WriteString(" FROM b_stat_hit t1 ")
	hs.selectBuilder.WriteString(hs.joinBuilder.String())
	return hs
}

func (hs HitSqlBuilder) orderByBuild() HitSqlBuilder {
	if len(hs.filter.OrderBy) == 0 {
		return hs
	}
	hs.orderByBuilder.WriteString(" ORDER BY ")
	var orderByFields []string
	for _, by := range hs.filter.OrderBy {
		orderByFields = append(orderByFields, hitFields[by])
	}
	hs.orderByBuilder.WriteString(strings.Join(orderByFields, ","))
	hs.orderByBuilder.WriteString(" ")
	hs.orderByBuilder.WriteString(hs.filter.TypeSort)
	return hs
}

func (hs HitSqlBuilder) whereBuild() HitSqlBuilder {
	where := hs.filter.Where
	if utf8.RuneCountInString(where) == 0 {
		return hs
	}
	for key, value := range hs.filter.Params {
		where = strings.ReplaceAll(where, key, " ? ")
		*hs.params = append(*hs.params, value)
	}

	for key, value := range hitFields {
		var re = regexp.MustCompile(`\b` + key + `\b`)
		where = re.ReplaceAllString(where, value)
	}
	hs.whereBuilder.WriteString(" WHERE ")
	hs.whereBuilder.WriteString(where)
	return hs
}

func (hs HitSqlBuilder) BuildSQL() SQL {
	var resultSQL strings.Builder
	hs.buildSelectAndJoin().
		whereBuild().
		orderByBuild()
	resultSQL.WriteString(hs.selectBuilder.String())
	resultSQL.WriteString(hs.whereBuilder.String())
	resultSQL.WriteString(hs.orderByBuilder.String())
	return SQL{
		SQL:    resultSQL.String(),
		Params: *hs.params,
	}
}
