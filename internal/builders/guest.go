package builders

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/utils"
	"errors"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

var whereFields = []string{
	"id",          // - ID посетителя;
	"registered",  // - был ли посетитель когда-либо авторизован на сайте, возможные значения: 1 - был; 0 - не был.
	"first_date1", //- начальное значение интервала для поля "дата первого захода на сайт";
	//"first_date2", // конечное значение интервала для поля "дата первого захода на сайт";
	"last_date", //начальное значение интервала для поля "дата последнего захода на сайт";
	//"last_date2", //конечное значение интервала для поля "дата первого захода на сайт";
	"period_date1", //начальное значение интервала для даты посещения посетителем сайта;
	//"period_date2", //конечно значение интервала для даты посещения посетителем сайта;
	"site_id",       //ID сайта первого либо последнего захода;
	"first_site_id", // ID сайта первого захода;
	"last_site_id",  // ID сайта последнего захода;
	"url",           // страница откуда впервые пришел посетитель, страница на которую впервые пришел посетитель и последняя страница просмотренная посетителем;
	"url_404",       // была ли 404 ошибка на первой странице или на последней странице посещенной посетителем, возможные значения: Y - была; N - не было.
	"user_agent",    // UserAgent посетителя на последнем заходе;
	"adv",           // флаг "приходил ли посетитель когда-либо по рекламной кампании (не равной NA/NA)", возможные значения:
	//1 - посетитель приходил по какой-либо рекламной кампании (не равной NA/NA);
	//0 - не приходил никогда ни по одной рекламной кампании (не равной NA/NA).
	"adv_id",     // ID рекламной кампании первого либо последнего захода посетителя (при этом это мог быть как прямой заход, так и возврат по рекламной кампании);
	"referer1",   // идентификатор referer1 рекламной кампании первого либо последнего захода посетителя;
	"referer2",   // идентификатор referer2 рекламной кампании первого либо последнего захода посетителя;
	"referer3",   // дополнительный параметр referer3 рекламной кампании первого либо последнего захода посетителя;
	"events",     // начальное значение для интервала кол-ва событий сгенерированных посетителем;
	"sess",       // начальное значение для интервала кол-ва сессий сгенерированных посетителем;
	"hits",       // начальное значение для интервала кол-ва хитов сгенерированных посетителем;
	"favorites",  // флаг "добавлял ли посетитель сайт в "Избранное"", возможные значения: Y - добавлял; N - не добавлял.
	"ip",         // IP адрес посетителя сайта в последнем заходе;
	"lang",       // языки установленные в настройках браузера посетителя в последнем заходе;
	"country_id", // ID страны (двух символьный идентификатор) посетителя в последнем заходе;
	"country",    // название страны;;
	"user",       // ID, логин, имя, фамилия пользователя, под которыми посетитель последний раз был авторизован;
	"user_id",    // ID пользователя, под которым посетитель последний раз был авторизован;
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

type GuestSQLBuilder struct {
	filter        filters.Filter
	selectBuilder *sqlbuilder.SelectBuilder
	argsSql       []interface{}
}

func NewGuestBuilder(filter filters.Filter) GuestSQLBuilder {
	return GuestSQLBuilder{
		filter:        filter,
		selectBuilder: sqlbuilder.NewSelectBuilder(),
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
	var values []interface{}
	var whereCondition []string
	sessionJoined := false
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

		//case "last_date2":
		//	whereCondition = append(whereCondition, "g.last_date < ?")
		//	values = append(values, value.Value)

		case "period_date1":
			whereCondition = append(whereCondition, fmt.Sprintf("g.date_first %s ?", value.Operator))
			values = append(values, value.Value)
			if sessionJoined == false {
				g.selectBuilder.Join("sessions s", "s.guest_id = g.id")
			}
			sessionJoined = true
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
			}else {
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
		case "ADV_ID":
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
			$arSqlSearch[] = "G.C_EVENTS>='".intval($val)."'"
			break
		case "SESS":
			$arSqlSearch[] = "G.SESSIONS>='".intval($val)."'"
			break
		case "HITS1":
			$arSqlSearch[] = "G.HITS>='".intval($val)."'"
			break
		case "FAVORITES":
			if ($val == "Y")
			$arSqlSearch[] = "G.FAVORITES='Y'"
			elseif($val == "N")
			$arSqlSearch[] = "G.FAVORITES<>'Y'"
			break
		case "IP":
			$match = ($arFilter[$key.
			"_EXACT_MATCH"]=="Y" && $match_value_set) ? "N": "Y";
		$arSqlSearch[] = GetFilterQuery("G.LAST_IP",$val, $match, array("."))
			break
		case "LANG":
			$match = ($arFilter[$key.
			"_EXACT_MATCH"]=="Y" && $match_value_set) ? "N": "Y";
		$arSqlSearch[] = GetFilterQuery("G.LAST_LANGUAGE", $val, $match)
			break
		case "COUNTRY_ID":
			$match = ($arFilter[$key.
			"_EXACT_MATCH"]=="Y" && $match_value_set) ? "N": "Y";
		$arSqlSearch[] = GetFilterQuery("G.LAST_COUNTRY_ID", $val, $match)
			break
		case "COUNTRY":
			$match = ($arFilter[$key.
			"_EXACT_MATCH"]=="Y" && $match_value_set) ? "N": "Y";
		$arSqlSearch[] = GetFilterQuery("C.NAME", $val, $match);
		$select1.= " , C.NAME LAST_COUNTRY_NAME ";
		$from2 = " LEFT JOIN b_stat_country C ON (C.ID = G.LAST_COUNTRY_ID) ";
		$arrGroup["C.NAME"] = true;
		$bGroup = true
			break
		case "REGION":
			$match = ($arFilter[$key.
			"_EXACT_MATCH"]=="Y" && $match_value_set) ? "N": "Y";
		$arSqlSearch[] = GetFilterQuery("CITY.REGION", $val, $match)
			break
		case "CITY_ID":
			$match = ($arFilter[$key.
			"_EXACT_MATCH"]=="Y" && $match_value_set) ? "N": "Y";
		$arSqlSearch[] = GetFilterQuery("G.LAST_CITY_ID", $val, $match)
			break
		case "CITY":
			$match = ($arFilter[$key.
			"_EXACT_MATCH"]=="Y" && $match_value_set) ? "N": "Y";
		$arSqlSearch[] = GetFilterQuery("CITY.NAME", $val, $match)
			break
		case "USER":
		case "USER_ID":
			if intval($val) > 0) {
			$arSqlSearch[] = "G.LAST_USER_ID=".intval($val)
		} else {
			$arSqlSearch[] = $DB- > IsNull("G.LAST_USER_ID", "0").
			">0"
		}
			$arrGroup["G.LAST_USER_ID"] = true
			$bGroup = true
			break
		}
	}
}

//sql, args, err := BuildWhereSQL(g.filter, func(field string) bool {
//	return slices.Contains(whereFields, field)
//})
//if err != nil {
//	return "", err
//}
//g.argsSql = args
g.selectBuilder.Where(whereCondition...)

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
