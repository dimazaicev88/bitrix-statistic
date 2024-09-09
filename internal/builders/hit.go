package builders

import (
	"bitrix-statistic/internal/filters"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"slices"
)

type HitSqlBuilder struct {
	filter     filters.Filter
	sqlBuilder *sqlbuilder.SelectBuilder
}

func NewHitSQLBuilder(filter filters.Filter) HitSqlBuilder {
	return HitSqlBuilder{
		filter:     filter,
		sqlBuilder: sqlbuilder.NewSelectBuilder(),
	}
}

var hitSQLFields = []string{
	"uuid",
	"session_uuid",
	"advUuid",
	"dateHit",
	"phpSessionId",
	"guestUuid",
	"newGuest",
	"userId",
	"userAuth",
	"url",
	"url404",
	"urlFrom",
	"ip",
	"method",
	"cookies",
	"userAgent",
	"stopListUuid",
	"countryId",
	"cityUuid",
	"siteId",
}

var hitFilterFields = []string{
	"uuid", "guestUuid", "isNewGuest", "sessionUuid", "stopListUuid", "url", "isUrl404", "userId",
	"isRegistered", "date", "ip", "userAgent", "countryId", "country", "cookie", "isStop", "siteId",
}

func (hs HitSqlBuilder) buildSelect() (HitSqlBuilder, error) {
	if len(hs.filter.Fields) == 0 {
		hs.sqlBuilder.Select("*")
	} else {
		for _, value := range hs.filter.Fields {
			if slices.Contains(hitSQLFields, value) == false {
				return HitSqlBuilder{}, fmt.Errorf("unknown field: %s", value)
			}
		}
		hs.sqlBuilder.Select(hs.filter.Fields...)
	}

	return hs, nil
}

//
//func (hs HitSqlBuilder) orderByBuild() SQLBuild {
//	return NewOrderByBuilder(hs.sqlData).BuildDefault()
//}
//
//func (hs HitSqlBuilder) whereBuild() OrderByBuilder {
//	return NewWhereBuilder(hs.sqlData).BuildDefault()
//}
//
//func (hs HitSqlBuilder) BuildSQL() (SQL, error) {
//	return NewSQLBuild(hs.sqlData).DefaultBuild(hs.buildSelect)
//}
//
//func (hs HitSqlBuilder) template() string {
//
//	sql := "SELECT H.ID,  H.SESSION_ID, H.GUEST_ID, H.NEW_GUEST, H.USER_ID, H.USER_AUTH," +
//		"	H.URL,		H.URL_404,		H.URL_FROM,		H.IP,		H.METHOD,		H.COOKIES,		H.USER_AGENT," +
//		"	H.STOP_LIST_ID,		H.COUNTRY_ID,		H.CITY_ID,		CITY.REGION REGION_NAME,		CITY.NAME CITY_NAME," +
//		"H.SITE_ID,		DATE_FORMAT(H.DATE_HIT, '%d.%m.%Y %H:%i:%s') as  DATE_HIT" +
//		" . $select . " +
//		"FROM" +
//		"b_stat_hit H" +
//		"LEFT JOIN b_stat_city CITY ON (CITY.ID = H.CITY_ID)" +
//		" . $from1 . " +
//		" . $from2 . " +
//		"	WHERE" +
//		" . $strSqlSearch . " +
//		" . $strSqlOrder . "
//
//	return sql
//}
