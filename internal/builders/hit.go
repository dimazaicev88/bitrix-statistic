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
	"SessionId":          " t1.SessionId ",
	"GuestId":            " t1.GuestId ",
	"NewGuest":           " t1.NewGuest ",
	"USER_ID":            " t1.USER_ID ",
	"USER_AUTH":          " t1.USER_AUTH ",
	"Url":                " t1.Url ",
	"Url404":             " t1.Url404 ",
	"URL_FROM":           " t1.URL_FROM ",
	"Ip":                 " t1.Ip ",
	"METHOD":             " t1.METHOD ",
	"COOKIES":            " t1.COOKIES ",
	"UserAgent":          " t1.UserAgent ",
	"StopListId":         " t1.StopListId ",
	"CountryId":          " t1.CountryId ",
	"CityId":             " t1.CityId ",
	"Region REGION_NAME": " t3.Region ",
	"USER":               " t2.LOGIN, t2.NAME ",
	"NAME CITY_NAME":     " t3.CITY_NAME ",
	"SiteId":             " t1.SiteId ",
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
				if selectField == "CountryId" {
					sqlData.joinBuilder.WriteString(" INNER JOIN b_stat_country t3 ON (t3.ID = t1.CountryId)")
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

func (hs HitSqlBuilder) template() string {

	sql := "	SELECT H.ID,  H.SESSION_ID, H.GUEST_ID, H.NEW_GUEST, H.USER_ID, H.USER_AUTH," +
		"	H.URL,		H.URL_404,		H.URL_FROM,		H.IP,		H.METHOD,		H.COOKIES,		H.USER_AGENT," +
		"	H.STOP_LIST_ID,		H.COUNTRY_ID,		H.CITY_ID,		CITY.REGION REGION_NAME,		CITY.NAME CITY_NAME," +
		"H.SITE_ID,		DATE_FORMAT(H.DATE_HIT, '%d.%m.%Y %H:%i:%s') as  DATE_HIT" +
		" . $select . " +
		"FROM" +
		"b_stat_hit H" +
		"LEFT JOIN b_stat_city CITY ON (CITY.ID = H.CITY_ID)" +
		" . $from1 . " +
		" . $from2 . " +
		"	WHERE" +
		" . $strSqlSearch . " +
		" . $strSqlOrder . "

	return sql
}
