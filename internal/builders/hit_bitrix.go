package builders

import (
	"bitrix-statistic/internal/filters"
	"github.com/volatiletech/null/v9"
	"strings"
)

type HitBitrixSqlBuilder struct {
	filter filters.Filter
}

func NewHitBitrixSQLBuilder(filter filters.Filter) HitBitrixSqlBuilder {
	return HitBitrixSqlBuilder{
		filter: filter,
	}
}

func (hs HitBitrixSqlBuilder) buildIntField(fieldName, execMath string, value null.Int) {

}

func (hs HitBitrixSqlBuilder) buildStringField(fieldName, execMath, value string) {

}

func (hs HitBitrixSqlBuilder) buildSQL() {
	//hs.buildIntField("ID", hs.filter.Filter.IdExactMatch, hs.filter.Filter.ID)
	//hs.buildIntField("SESSION_ID", hs.filter.Filter.SessionIdExactMatch, hs.filter.Filter.SessionId)
	//hs.buildIntField("STOP_LIST_ID", hs.filter.Filter.StopListIdExactMatch, hs.filter.Filter.StopListId)
	//hs.buildIntField("COUNTRY_ID", hs.filter.Filter.CountryIdExactMatch, hs.filter.Filter.CountryId)
	//hs.buildIntField("CITY_ID", hs.filter.Filter.CityExactMatch, hs.filter.Filter.CityId)
	//
	//hs.buildStringField("URL", hs.filter.Filter.UrlExactMatch, hs.filter.Filter.Url)
	//hs.buildStringField("URL_404", hs.filter.Filter.Url404ExactMatch, hs.filter.Filter.Url404)
	//hs.buildStringField("NEW_GUEST", hs.filter.Filter.NewGuestExactMatch, hs.filter.Filter.NewGuest)
	//hs.buildStringField("IP", hs.filter.Filter.IpExactMatch, hs.filter.Filter.Ip)
	//hs.buildStringField("USER_AGENT", hs.filter.Filter.UserAgentExactMatch, hs.filter.Filter.UserAgent)
	//hs.buildStringField("COOKIE", hs.filter.Filter.CookieExactMatch, hs.filter.Filter.Cookie)
	//
	//if hs.filter.Filter.Registered == "Y" {
	//	hs.where.WriteString(" t_hit.USER_ID>0 ")
	//} else if hs.filter.Filter.Registered == "N" {
	//	hs.where.WriteString(" t_hit.USER_ID<=0 or t_hit.USER_ID is null ")
	//}
	//
	//if strings.Trim(hs.filter.Filter.User, " ") != "" {
	//	hs.join.WriteString(" LEFT JOIN b_user t_user ON (t_user.ID = t_hit.USER_ID) ")
	//	if hs.filter.Filter.UserExactMatch == "N" {
	//		hs.where.WriteString(
	//			StringConcat(
	//				" t_hit.USER_ID like '%",
	//				hs.filter.Filter.User,
	//				"%' ",
	//				hs.whereLogic,
	//				" t_user.LOGIN like '%",
	//				hs.filter.Filter.User,
	//				"%' ",
	//				hs.whereLogic,
	//				" t_user.LAST_NAME like '%",
	//				hs.filter.Filter.User,
	//				"%' ",
	//				hs.whereLogic,
	//				" t_user.NAME like '%",
	//				hs.filter.Filter.User,
	//				"%",
	//				"' ",
	//			),
	//		)
	//	} else {
	//		hs.where.WriteString(
	//			StringConcat(
	//				" t_hit.USER_ID=",
	//				hs.filter.Filter.User,
	//				" ",
	//				hs.whereLogic,
	//				" ",
	//				" t_user.LOGIN=",
	//				"'",
	//				hs.filter.Filter.User,
	//				"' ",
	//				hs.whereLogic,
	//				" t_user.LAST_NAME=",
	//				"'",
	//				hs.filter.Filter.User,
	//				"' ",
	//				hs.whereLogic,
	//				" t_user.NAME=",
	//				"'",
	//				hs.filter.Filter.User,
	//				"' ",
	//			),
	//		)
	//	}
	//}
}

func (hs HitBitrixSqlBuilder) BuildSQL() (string, error) {
	var result strings.Builder

	//hs.selectFields.WriteString("SELECT" +
	//	" t_hit.ID," +
	//	"t_hit.SESSION_ID," +
	//	"t_hit.GUEST_ID," +
	//	"t_hit.NEW_GUEST," +
	//	"t_hit.USER_ID," +
	//	"t_hit.USER_AUTH," +
	//	"t_hit.URL," +
	//	"t_hit.URL_404," +
	//	"t_hit.URL_FROM," +
	//	"t_hit.IP," +
	//	"t_hit.METHOD," +
	//	"t_hit.COOKIES," +
	//	" t_hit.USER_AGENT," +
	//	"t_hit.STOP_LIST_ID," +
	//	"t_hit.COUNTRY_ID," +
	//	"t_hit.CITY_ID," +
	//	" t_city.REGION REGION_NAME," +
	//	" t_city.NAME CITY_NAME," +
	//	" t_hit.SITE_ID," +
	//	"DATE_FORMAT(t_hit.DATE_HIT, '%d.%m.%Y %H:%i:%s') as  DATE_HIT" +
	//	" from b_stat_hit t_hit")
	//hs.join.WriteString(" LEFT JOIN b_stat_city t_city ON (t_city.ID = t_hit.CITY_ID)")
	//
	//hs.buildSQL()
	//result.WriteString(hs.selectFields.String())
	//result.WriteString(hs.join.String())
	//result.WriteString(" where ")
	//result.WriteString(hs.where.String())
	//result.WriteString(hs.order.String())
	return result.String(), nil
}
