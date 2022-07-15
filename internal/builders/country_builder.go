package builders

import (
	"bitrix-statistic/internal/filters"
	"errors"
	"regexp"
	"strings"
	"unicode/utf8"
)

var cityFields = map[string]string{
	"ID":         "t1.ID",
	"SHORT_NAME": "t1.SHORT_NAME",
	"NAME":       "t1.NAME",
	"SESSIONS":   "t1.SESSIONS",
	"NEW_GUESTS": "t1.NEW_GUESTS",
	"HITS":       "t1.HITS",
	"C_EVENTS":   "t1.C_EVENTS",
}

type CountrySQLBuilder struct {
	SQLBuilder
}

func NewCountrySQLBuilder(filter filters.Filter) CountrySQLBuilder {
	return CountrySQLBuilder{NewSQLBuilder(filter)}
}

func (cb CountrySQLBuilder) buildSelect() error {
	var selectFields []string
	cb.selectBuilder.WriteString("SELECT ")
	if len(cb.filter.Select) == 0 {
		cb.selectBuilder.WriteString("* ")
	} else {
		for _, selectField := range cb.filter.Select {
			if value, ok := cityFields[selectField]; ok {
				selectFields = append(selectFields, value)
			}
		}
	}
	if len(selectFields) == 0 {
		return errors.New("unknown fields is select")
	}
	cb.selectBuilder.WriteString(strings.Join(selectFields, ","))
	cb.selectBuilder.WriteString(" FROM b_stat_country t1 ")
	cb.selectBuilder.WriteString(cb.joinBuilder.String())
	return nil
}

func (cb CountrySQLBuilder) orderByBuild() CountrySQLBuilder {
	if len(cb.filter.OrderBy) == 0 {
		return cb
	}
	cb.orderByBuilder.WriteString(" ORDER BY ")
	var orderByFields []string
	for _, by := range cb.filter.OrderBy {
		orderByFields = append(orderByFields, hitFields[by])
	}
	cb.orderByBuilder.WriteString(strings.Join(orderByFields, ","))
	cb.orderByBuilder.WriteString(" ")
	cb.orderByBuilder.WriteString(cb.filter.TypeSort)
	return cb
}

func (cb CountrySQLBuilder) whereBuild() CountrySQLBuilder {
	where := cb.filter.Where
	if utf8.RuneCountInString(where) == 0 {
		return cb
	}
	for key, value := range cb.filter.Params {
		where = strings.ReplaceAll(where, key, " ? ")
		*cb.params = append(*cb.params, value)
	}

	for key, value := range hitFields {
		var re = regexp.MustCompile(`\b` + key + `\b`)
		where = re.ReplaceAllString(where, value)
	}
	cb.whereBuilder.WriteString(" WHERE ")
	cb.whereBuilder.WriteString(where)
	return cb
}

func (cb CountrySQLBuilder) BuildSQL() (error, SQL) {
	var resultSQL strings.Builder
	err := cb.buildSelect()
	if err != nil {
		return err, SQL{}
	}
	cb.whereBuild()
	cb.orderByBuild()
	resultSQL.WriteString(cb.selectBuilder.String())
	resultSQL.WriteString(cb.whereBuilder.String())
	resultSQL.WriteString(cb.orderByBuilder.String())
	return nil, SQL{
		SQL:    resultSQL.String(),
		Params: *cb.params,
	}
}
