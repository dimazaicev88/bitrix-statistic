package builders

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/utils"
	"errors"
	"golang.org/x/exp/slices"
	"strings"
)

type HitSqlBuilder struct {
	sqlData SQLDataForBuild
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

func (hs HitSqlBuilder) buildSelect() (WhereBuilder, error) {
	return NewSelectBuild(hs.sqlData).Build(func(sqlData SQLDataForBuild) (WhereBuilder, error) {
		var selectFields []string
		sqlData.selectBuilder.WriteString("SELECT ")
		if len(sqlData.filter.Select) == 0 {
			sqlData.selectBuilder.WriteString("* ")
		} else {
			set := utils.NewSet[string]()
			slices.Sort(set.SliceAsSet(sqlData.filter.Select).Items())
			sqlData.filter.Select = set.SliceAsSet(sqlData.filter.Select).Items()
			slices.Sort(sqlData.filter.Select)
			for _, selectField := range sqlData.filter.Select {
				if value, ok := hitFields[selectField]; ok {
					selectFields = append(selectFields, value)
				} else {
					return WhereBuilder{}, errors.New("unknown field " + selectField)
				}
				if selectField == "USER" {
					sqlData.selectBuilder.WriteString("")
					sqlData.joinBuilder.WriteString(" LEFT JOIN b_user t2 ON (t2.ID = t1.USER_ID)")
				}
				if selectField == "COUNTRY_ID" {
					sqlData.joinBuilder.WriteString(" INNER JOIN b_stat_country t3 ON (t3.ID = t1.COUNTRY_ID)")
				}
			}
		}
		sqlData.selectBuilder.WriteString(strings.Join(selectFields, ","))
		sqlData.selectBuilder.WriteString(" FROM b_stat_hit t1 ")
		sqlData.selectBuilder.WriteString(sqlData.joinBuilder.String())
		return NewWhereBuilder(sqlData), nil
	})
}

func (hs HitSqlBuilder) orderByBuild() SQLBuild {
	return NewOrderByBuilder(hs.sqlData).BuildDefault()
}

func (hs HitSqlBuilder) whereBuild() OrderByBuilder {
	return NewWhereBuilder(hs.sqlData).BuildDefault()
}

func (hs HitSqlBuilder) BuildSQL() (SQL, error) {
	return NewSQLBuild(hs.sqlData).DefaultBuild(hs.buildSelect)
}
