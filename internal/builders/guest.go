package builders

import (
	"bitrix-statistic/internal/filters"
	"errors"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

type GuestSQLBuilder struct {
	filter        filters.Filter
	selectBuilder *sqlbuilder.SelectBuilder
	whereBuilder  *sqlbuilder.WhereClause
}

var selectFields = map[string]string{
	"id":                "g.id",
	"favorites":         "g.favorites",         //[0|1] флаг добавления сайта в "Избранное"
	"c_events":          "g.c_events",          //количество событий сгенерированных данным посетителей
	"sessions":          "g.sessions",          //количество сессий данного посетителя
	"hits":              "g.hits",              //количество хитов данного посетителя
	"first_session_id":  "g.first_session_id",  //ID сессии первого захода на сайт
	"first_site_id":     "g.first_site_id",     //ID сайта на который посетитель впервые пришел
	"last_site_id":      "g.last_site_id",      //ID сайта последнего захода
	"first_url_from":    "g.first_url_from",    //адрес страницы с которой посетитель впервые пришел на сайт
	"first_url_to":      "g.first_url_to",      //адрес страницы сайта на которую посетитель впервые пришел
	"first_url_to_404":  "g.first_url_to_404",  // флаг 404 ошибки (страница не существует) на странице сайта на которую посетитель впервые пришел
	"first_adv_id":      "g.first_adv_id",      //ID рекламной кампании по которой посетитель впервые пришел на сайт
	"first_referer1":    "g.first_referer1",    //идентификатор referer1 рекламной кампании FIRST_ADV_ID
	"first_referer2":    "g.first_referer2",    //идентификатор referer2 рекламной кампании FIRST_ADV_ID
	"first_referer3":    "g.first_referer3",    //дополнительный параметр referer3 рекламной кампании FIRST_ADV_ID
	"last_adv_id":       "g.last_adv_id",       //ID рекламной кампании по которой посетитель пришел на сайт в последнем заходе
	"last_adv_back":     "g.last_adv_back",     //[1|0] флаг того был ли это возврат (1) или прямой заход (0) по рекламной кампании LAST_ADV_ID
	"last_referer1":     "g.last_referer1",     //идентификатор referer1 рекламной кампании LAST_ADV_ID
	"last_referer2":     "g.last_referer2",     //идентификатор referer2 рекламной кампании LAST_ADV_ID
	"last_referer3":     "g.last_referer3",     //дополнительный параметр referer3 рекламной кампании LAST_ADV_ID
	"last_user_id":      "g.last_user_id",      //ID пользователя
	"last_user_auth":    "g.last_user_auth",    //[1|0] был ли авторизован посетитель в последнем заходе на сайт
	"last_url_last":     "g.last_url_last",     //адрес последней страницы на которую зашел посетитель
	"last_url_last_404": "g.last_url_last_404", //флаг 404 ошибки (страница не существует) на последней странице сайта на которую зашел посетитель
	"last_user_agent":   "g.last_user_agent",   //UserAgent посетителя в последнем заходе
	"last_ip":           "g.last_ip",           //IP адрес посетителя сайта в последнем заходе
	"last_language":     "g.last_language",     //языки установленные в настройках браузера посетителя в последнем заходе
	"last_country_id":   "g.last_country_id",   //ID страны посетителя в последнем заходе
	//"last_city_id":      "g.last_city_id",      //TODO ?????
	"first_date":        "g.first_date",      //время первого захода на сайт
	"last_date":         "g.last_date",       //время первого захода на сайт
	"last_session_id":   "g.last_session_id", //ID сессии последнего захода на сайт
	"last_country_name": "c.name",
	"last_city_name":    "city.last_city_name",
}

func NewGuestBuilder(filter filters.Filter) GuestSQLBuilder {
	return GuestSQLBuilder{
		filter:        filter,
		selectBuilder: sqlbuilder.NewSelectBuilder(),
		whereBuilder:  sqlbuilder.NewWhereClause(),
	}
}

func (g *GuestSQLBuilder) Select() error {
	arrFields := make([]string, 0, len(g.filter.Fields))
	for _, field := range g.filter.Fields {
		if _, ok := selectFields[field]; !ok {
			return errors.New("unknown field: " + field)
		}
		arrFields = append(arrFields, fmt.Sprintf("%s as %s", selectFields[field], field))
	}
	if len(arrFields) == 0 {
		arrFields = append(arrFields, "id as g.id")
	}
	g.selectBuilder.Select(arrFields...).From("guest g")
	return nil
}

func (g *GuestSQLBuilder) Where() (string, error) {
	return "", nil
}

func (g *GuestSQLBuilder) ToString() (string, error) {
	if err := g.Select(); err != nil {
		return "", err
	}
	return g.selectBuilder.String(), nil
	//err := g.Select()
	//if err != nil {
	//	return "", err
	//}
	//
	//where, err := g.Where()
	//if err != nil {
	//	return "", err
	//}
	//return utils.StringConcat(selectFields, where), nil
}
