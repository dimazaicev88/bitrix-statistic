package converters

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/utils"
	"errors"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

var whereFields = []string{
	"id",                    // ID посетителя;
	"registered",            // был ли посетитель когда-либо авторизован на сайте, возможные значения: 1 - был; 0 - не был.
	"date_first_visit_site", // дата первого захода на сайт;
	"date_last_visit_site",  // дата последнего захода на сайт;
	"site_id",               //ID сайта первого либо последнего захода;
	"first_site_id",         // ID сайта первого захода;
	"last_site_id",          // ID сайта последнего захода;
	"url",                   // страница откуда впервые пришел посетитель, страница на которую впервые пришел посетитель и последняя страница просмотренная посетителем;
	"url_404",               // была ли 404 ошибка на первой странице или на последней странице посещенной посетителем, возможные значения: Y - была; N - не было.
	"user_agent",            // UserAgent посетителя на последнем заходе;
	"adv",                   // флаг "приходил ли посетитель когда-либо по рекламной кампании (не равной NA/NA)", возможные значения:
	//1 - посетитель приходил по какой-либо рекламной кампании (не равной NA/NA);
	//0 - не приходил никогда ни по одной рекламной кампании (не равной NA/NA).
	"adv_id",         // ID рекламной кампании первого либо последнего захода посетителя (при этом это мог быть как прямой заход, так и возврат по рекламной кампании);
	"referer1",       // идентификатор referer1 рекламной кампании первого либо последнего захода посетителя;
	"referer2",       // идентификатор referer2 рекламной кампании первого либо последнего захода посетителя;
	"referer3",       // дополнительный параметр referer3 рекламной кампании первого либо последнего захода посетителя;
	"events",         // начальное значение для интервала кол-ва событий сгенерированных посетителем;
	"count_sessions", // кол-ва сессий сгенерированных посетителем;
	"count_hits",     // кол-ва хитов сгенерированных посетителем;
	"favorites",      // флаг "добавлял ли посетитель сайт в "Избранное"", возможные значения: Y - добавлял; N - не добавлял.
	"ip",             // IP адрес посетителя сайта в последнем заходе;
	"lang",           // языки установленные в настройках браузера посетителя в последнем заходе;
	"country_id",     // ID страны (двух символьный идентификатор) посетителя в последнем заходе;
	"country",        // название страны;;
	"user",           // ID, логин, имя, фамилия пользователя, под которыми посетитель последний раз был авторизован;
	"user_id",        // ID пользователя, под которым посетитель последний раз был авторизован;
}

var selectFields = map[string]string{
	"id":                "g.id",
	"favorites":         "g.favorites",         //[0|1] флаг добавления сайта в "Избранное"
	"events":            "g.events",            //количество событий сгенерированных данным посетителей
	"sessions":          "g.sessions",          //количество сессий данного посетителя
	"hits":              "g.hits",              //количество хитов данного посетителя
	"session_id":        "g.session_id",        //ID сессии первого захода на сайт
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

type GuestSQLBuilder struct {
	filter        filters.Filter
	selectBuilder *sqlbuilder.SelectBuilder
	argsSql       []any
	selectBuffer  []string
	joinBuffer    []string
}

func NewGuestBuilder(filter filters.Filter) GuestSQLBuilder {
	return GuestSQLBuilder{
		filter:        filter,
		selectBuilder: sqlbuilder.NewSelectBuilder(),
		selectBuffer:  []string{},
	}
}

func (g *GuestSQLBuilder) build() (string, error) {
	var values []any
	var whereCondition []string
	needJoinSession := false
	needJoinCountry := false

	for _, field := range g.filter.Fields {
		if _, ok := selectFields[field]; !ok {
			return "", errors.New("unknown field: " + field)
		}
		g.selectBuffer = append(g.selectBuffer, fmt.Sprintf("%s as %s", selectFields[field], field))
	}
	if len(g.selectBuffer) == 0 {
		g.selectBuffer = append(g.selectBuffer, "id as g.id")
	}

	for _, value := range g.filter.Operators {
		switch value.Field {

		case "registered":
			if utils.IsInt(value.Value) == false {
				return "", errors.New("invalid value type: " + value.Field)
			}
			if value.Value.(int) > 0 {
				whereCondition = append(whereCondition, fmt.Sprintf("g.last_user_id > 0"))
			} else {
				whereCondition = append(whereCondition, fmt.Sprintf("g.last_user_id <= 0 or g.last_user_id is null"))
			}
			values = append(values, value.Value)

		case "first_date": //TODO передать на IN
			whereCondition = append(whereCondition, fmt.Sprintf("g.first_date %s ?", value.Operator))
			values = append(values, value.Value)

		case "last_date":
			whereCondition = append(whereCondition, fmt.Sprintf("g.last_date %s ?", value.Operator))
			values = append(values, value.Value)

		case "period_date1":
			whereCondition = append(whereCondition, fmt.Sprintf("g.date_first %s ?", value.Operator))
			values = append(values, value.Value)
			needJoinSession = true
			//"count(S.ID) as SESS,";  TODO зачем ???

		case "site_id":
			whereCondition = append(whereCondition, fmt.Sprintf("g.site_id %s ?", value.Operator))
			values = append(values, value.Value)

		case "last_site_id":
			whereCondition = append(whereCondition, fmt.Sprintf("g.last_site_id %s ?", value.Operator))
			values = append(values, value.Value)

		case "first_site_id":
			whereCondition = append(whereCondition, fmt.Sprintf("g.first_site_id %s ?", value.Operator))
			values = append(values, value.Value)

		case "url":
			whereCondition = append(whereCondition, fmt.Sprintf("g.first_site_id %s ?", value.Operator))
			values = append(values, value.Value)
		//TODO ADD G.FIRST_URL_FROM,G.FIRST_URL_TO,G.LAST_URL_LAST

		case "url_404":
			if utils.IsInt(value.Value) == false {
				return "", errors.New("invalid value type: " + value.Field)
			}
			whereCondition = append(whereCondition, "g.first_url_to_404=? or g.last_url_last_404=?")
			if value.Value.(int) == 1 {
				values = append(values, 1, 1)
			} else {
				values = append(values, 0, 0)
			}

		case "user_agent":
			whereCondition = append(whereCondition, fmt.Sprintf("g.last_user_agent %s ?", value.Operator))
			values = append(values, value.Value)

		case "adv":
			if utils.IsInt(value.Value) == false {
				return "", errors.New("invalid value type: " + value.Field)
			}
			if value.Value.(int) == 1 {
				whereCondition = append(whereCondition, `g.first_adv_id >0 and g.first_referer1<>'NA' and g.first_referer2<>'NA' 
                              or G.last_adv_id > 0 and G.last_adv_id is not null and G.last_referer1 <> 'NA' and G.last_referer2 <> 'NA'`)
			} else {
				whereCondition = append(whereCondition, `G.FIRST_ADV_ID<=0 or
				G.FIRST_ADV_ID is null or
				(G.FIRST_REFERER1='NA' and G.FIRST_REFERER2='NA')
			) and (
					G.LAST_ADV_ID<=0 or
				G.LAST_ADV_ID is null or
				(G.LAST_REFERER1='NA' and G.LAST_REFERER2='NA')
			))`)

			}
			values = append(values, value.Value)
		case "adv_id":
			whereCondition = append(whereCondition, fmt.Sprintf("g.first_adv_id %s ? or g.last_adv_id %s ?", value.Operator, value.Operator))
			values = append(values, value.Value)

		case "REFERER1":
			whereCondition = append(whereCondition, fmt.Sprintf("g.FIRST_REFERER1 %s ? or g.LAST_REFERER1 %s ?", value.Operator, value.Operator))
			values = append(values, value.Value)
		case "REFERER2":
			whereCondition = append(whereCondition, fmt.Sprintf("g.FIRST_REFERER2 %s ? or g.LAST_REFERER2 %s ?", value.Operator, value.Operator))
			values = append(values, value.Value)
		case "REFERER3":
			whereCondition = append(whereCondition, fmt.Sprintf("g.FIRST_REFERER3 %s ? or g.LAST_REFERER3 %s ?", value.Operator, value.Operator))
			values = append(values, value.Value)
		case "EVENTS":
			whereCondition = append(whereCondition, fmt.Sprintf("g.c_events %s ? ", value.Operator))
			values = append(values, value.Value)
		case "SESS":
			whereCondition = append(whereCondition, fmt.Sprintf("g.sessions %s ? ", value.Operator))
			values = append(values, value.Value)
		case "HITS":
			whereCondition = append(whereCondition, fmt.Sprintf("g.hits %s ? ", value.Operator))
			values = append(values, value.Value)
		case "favorites":
			if utils.IsInt(value.Value) == false {
				return "", errors.New("invalid value type: " + value.Field)
			}
			whereCondition = append(whereCondition, fmt.Sprintf("g.favorites %s ? ", value.Operator))
			values = append(values, value.Value)
		case "ip":
			whereCondition = append(whereCondition, fmt.Sprintf("g.last_ip %s ? ", value.Operator))
			values = append(values, value.Value)
		case "lang":
			whereCondition = append(whereCondition, fmt.Sprintf("g.last_language %s ? ", value.Operator))
			values = append(values, value.Value)

		case "country_id":
			//, C.NAME LAST_COUNTRY_NAME
			whereCondition = append(whereCondition, fmt.Sprintf("g.last_language %s ? ", value.Operator))
			values = append(values, value.Value)

		case "country":
			g.selectBuffer = append(g.selectBuffer, "c.name as last_country_name")
			needJoinCountry = true
			whereCondition = append(whereCondition, fmt.Sprintf("c.name %s ? ", value.Operator))
			values = append(values, value.Value)

		case "region":
			whereCondition = append(whereCondition, fmt.Sprintf("city.region %s ? ", value.Operator))
			values = append(values, value.Value)
		case "city_id":
			whereCondition = append(whereCondition, fmt.Sprintf("g.last_city_id %s ? ", value.Operator))
			values = append(values, value.Value)
		case "city":
			whereCondition = append(whereCondition, fmt.Sprintf("city.name %s ? ", value.Operator))
			values = append(values, value.Value)
		case "USER": //TODO добавить фильтрацию по логину
		case "USER_ID":
			whereCondition = append(whereCondition, fmt.Sprintf("g.last_user_id %s ? ", value.Operator))
			values = append(values, value.Value)
		}

		if needJoinSession == false {
			g.selectBuilder.Join("sessions s", "s.guest_id = g.id")
		}

		if needJoinCountry == false {
			g.selectBuilder.Join("country c", "c.id = g.last_country_id")
		}

	}

	g.selectBuilder.Where(whereCondition...)
	return "", nil
}

func (g *GuestSQLBuilder) ToString() (string, error) {
	return g.selectBuilder.String(), nil
}
