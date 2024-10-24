package builders

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/utils"
	"fmt"
	"slices"
	"strings"
)

var filterFields = []string{
	"uuid",           //uuid РК;
	"period",         //период за который необходимо получить данные;
	"dateAdv",        //Время действия РК;
	"referer1",       //идентификатор referer1 РК;
	"referer2",       //идентификатор referer2 РК;
	"priority",       //приоритет РК;
	"newGuests",      //Новых посетителей
	"guestsBack",     //Посетителей на возврате
	"guests",         //Посетителей на прямом заходе
	"favoritesBack",  //Посетителей, добавившие сайт в "Избранное" на возврате
	"favorites",      //Посетителей, добавивших сайт в "Избранное" на прямом заходе
	"hostsBack",      //Хостов на возврате
	"hosts",          //Хостов на прямом заходе
	"sessionsBack",   //Сессий на возврате
	"sessions",       //Сессий на прямом заходе
	"hitsBack",       //Хитов на возврате
	"hits",           // Хитов на прямом заходе
	"cost",           //затраты на РК;
	"revenue",        //доходы с РК;
	"benefit",        //прибыль с РК;
	"roi",            //рентабельность РК;
	"attent",         //коэффициент внимательности посетителей РК;
	"visitorsPerDay", //среднее кол-во посетителей в день;
	"duration",       //длительность РК;
	"description",    //описание РК;
}

var advSelectFields = []string{
	//"uuid", "sessionUuid", "advUuid", "dateHit", "phpSessionId", "guestUuid", "newGuest", "userId",
	//"userAuth", "url", "url404", "urlFrom", "ip", "method", "cookies", "userAgent", "stopListUuid", "countryId",
	//"cityUuid", "siteId",

	"uuid",        // ID
	"PRIORITY",    // приоритет
	"referer1",    // идентификатор referer1
	"referer2",    // идентификатор referer2
	"description", // описание
	//[EVENTS_VIEW// режим показа списка событий, возможные значения:
	//	link - ссылкой на список
	//	list - списком
	//	event1 - сгруппированными по event1
	//	event2 - сгруппированными по event2
	"dateFirst",                 // дата первого прямого захода
	"dateLast",                  // дата последнего прямого захода или возврата
	"advTime",                   // длительность РК в секундах
	"attent",                    // внимательность на прямом заходе
	"attentBack",                // внимательность на возврате
	"newVisitors",               // процент новых посетителей
	"returnedVisitors",          // процент вернувшихся на сайт посетителей после прямого захода
	"visitorsPerDay",            // среднее кол-во посетителей в день
	"currency",                  // валюта в которой заданы финансовые показатели
	"cost",                      // затраты
	"revenue",                   // доход
	"benefit",                   // прибыль
	"sessionCost",               // стоимость сессии
	"visitorCost",               // стоимость посетителя
	"roi",                       // рентабельность
	"guests",                    // суммарное кол-во посетителей на прямом заходе
	"newGuests",                 // суммарное кол-во новых посетителей на прямом заходе
	"favorites",                 // суммарное кол-во посетителей, добавившие сайт в "Избранное" на прямом заходе
	"hosts",                     // суммарное кол-во хостов на прямом заходе
	"sessions",                  // суммарное кол-во сессий на прямом заходе
	"hits",                      // суммарное кол-во хитов на прямом заходе
	"guestsBack",                // суммарное кол-во посетителей на возврате
	"favoritesBack",             // суммарное кол-во посетителей, добавившие сайт в "Избранное" на возврате
	"hostsBack",                 // суммарное кол-во хостов на возврате
	"sessionsBack",              // суммарное кол-во сессий на возврате
	"hitsBack",                  // суммарное кол-во хитов на возврате
	"guestsToday",               // посетителей на прямом заходе за сегодня
	"guestsBackToday",           // посетителей на возврате за сегодня
	"newGuestsToday",            // новых посетителей на возврате за сегодня
	"favoritesToday",            // посетителей, добавившие сайт в "Избранное" на прямом заходе за сегодня
	"favoritesBackToday",        // посетителей, добавившие сайт в "Избранное" на возврате за сегодня
	"HostsToday",                // хостов на прямом заходе за сегодня
	"hostsBackToday",            // хостов на возврате за сегодня
	"sessionsToday",             // сессий на прямом заходе за сегодня
	"sessionsBackToday",         // сессий на возврате за сегодня
	"hitsToday",                 // хитов на прямом заходе за сегодня
	"hitsBackToday",             // хитов на возврате за сегодня
	"guestsYesterday",           // посетителей на прямом заходе за вчера
	"guestsBackYesterday",       // посетителей на возврате за вчера
	"newGuestsYesterday",        // новых посетителей на возврате за вчера
	"favoritesYesterday",        // посетителей, добавившие сайт в "Избранное" на прямом заходе за вчера
	"favoritesBackYesterday",    // посетителей, добавившие сайт в "Избранное" на возврате за вчера
	"hostsYesterday",            // хостов на прямом заходе за вчера
	"hostsBackYesterday",        // хостов на возврате за вчера
	"sessionsYesterday",         // сессий на прямом заходе за вчера
	"sessionsBackYesterday",     // сессий на возврате за вчера
	"hitsYesterday",             // хитов на прямом заходе за вчера
	"hitsBackYesterday",         // хитов на возврате за вчера
	"guestsBefYesterday",        // посетителей на прямом заходе за позавчера
	"guestsBackBefYesterday",    // посетителей на возврате за позавчера
	"newGuestsBefYesterday",     // новых посетителей на возврате за позавчера
	"favoritesBefYesterday",     // посетителей, добавившие сайт в "Избранное" на прямом заходе за позавчера
	"favoritesBackBefYesterday", // посетителей, добавившие сайт в "Избранное" на возврате за позавчера
	"hostsBefYesterday",         // хостов на прямом заходе за позавчера
	"hostsBackBefYesterday",     // хостов на возврате за позавчера
	"sessionsBefYesterday",      // сессий на прямом заходе за позавчера
	"sessionsBackBefYesterday",  // сессий на возврате за позавчера
	"hitsBefYesterday",          // хитов на прямом заходе за позавчера
	"hitsBackBefYesterday",      // хитов на возврате за позавчера
	"guestsPeriod",              // посетителей на прямом заходе за период
	"guestsBackPeriod",          // посетителей на возврате за период
	"newGuestsPeriod",           // новых посетителей на возврате за период
	"favoritesPeriod",           // посетителей, добавившие сайт в "Избранное" на прямом заходе за период
	"favoritesBackPeriod",       // посетителей, добавившие сайт в "Избранное" на возврате за период
	"hostsPeriod",               // хостов на прямом заходе за период
	"hostsBackPeriod",           // хостов на возврате за период
	"sessionsPeriod",            // сессий на прямом заходе за период
	"sessionsBackPeriod",        // сессий на возврате за период
	"hitsPeriod",                // хитов на прямом заходе за период
	"hitsBackPeriod",            // хитов на возврате за период
}

type AdvSqlBuilder struct {
	filter            filters.Filter
	sqlBuilder        *SqlBuilder
	groupBySqlBuilder strings.Builder
}

func NewAdvSQLBuilder(filter filters.Filter) AdvSqlBuilder {
	return AdvSqlBuilder{
		filter:     filter,
		sqlBuilder: NewSqlBuilder(),
	}
}

// TODO добавить сборку когда нету выбираемых полей
func (hs *AdvSqlBuilder) buildSelect() error {
	countFields := 0
	for _, field := range hs.filter.Fields {
		if field == "" {
			continue
		}
		if !slices.Contains(advSelectFields, field) {
			return fmt.Errorf("unknown field: %s", field)
		}
		countFields++
	}

	simpleFields := map[string]string{
		"uuid":          "t1.uuid",
		"referer1":      "t1.referer1",
		"referer2":      "t1.referer2",
		"eventsView":    "t1.eventsView",
		"description":   "t1.description",
		"guests":        "t2.guests",
		"newGuests":     "t2.newGuests",
		"favorites":     "t2.favorites",
		"hosts":         "t2.hosts",
		"sessions":      "t2.sessions",
		"hits":          "t2.hits",
		"guestsBack":    "t2.guestsBack",
		"favoritesBack": "t2.favoritesBack",
		"hostsBack":     "t2.hostsBack",
		"sessionsBack":  "t2.sessionsBack",
		"hitsBack":      "t2.hitsBack",
	}

	hs.sqlBuilder.Add(`SELECT `)
	hs.groupBySqlBuilder.WriteString(`GROUP BY `)
	for _, fieldName := range hs.filter.Fields {
		if _, ok := simpleFields[fieldName]; !ok {
			hs.sqlBuilder.Add(fieldName)
			hs.sqlBuilder.Add(` `)
		}
		hs.groupBySqlBuilder.WriteString(fieldName)
		hs.groupBySqlBuilder.WriteString(` `)

		switch fieldName {
		// today
		case "guestsToday":
			hs.sqlBuilder.Add(`sumIf(t3.guests_day, toStartOfDay(t3.date_stat) = today()) as guestsToday`)
			hs.sqlBuilder.Add(` `)

		case "newGuestsToday":
			hs.sqlBuilder.Add(`sumIf(t3.new_guests, toStartOfDay(t3.date_stat) = today()) as newGuestsToday`)
			hs.sqlBuilder.Add(` `)

		case "favoritesToday":
			hs.sqlBuilder.Add(`sumIf(t3.favorites, toStartOfDay(t3.date_stat) = today()) as	favoritesToday`)
			hs.sqlBuilder.Add(` `)

		case "hostsToday":
			hs.sqlBuilder.Add(`sumIf(t3.hosts_day, toStartOfDay(t3.date_stat) = today()) as	hostsToday`)
			hs.sqlBuilder.Add(` `)

		case "sessionsToday":
			hs.sqlBuilder.Add(`sumIf(t3.sessions, toStartOfDay(t3.date_stat) = today()) as sessionsToday`)
			hs.sqlBuilder.Add(` `)

		case "hitsToday":
			hs.sqlBuilder.Add(`sumIf(t2.hits, toStartOfDay(t3.date_stat) = today()) as hitsToday`)
			hs.sqlBuilder.Add(` `)

		case "guestsBackToday":
			hs.sqlBuilder.Add(`sumIf(t3.guests_day_back, toStartOfDay(t3.date_stat) = today()) as guestsBackToday`)
			hs.sqlBuilder.Add(` `)

		case "favoritesBackToday":
			hs.sqlBuilder.Add(`sumIf(t3.favorites_back, toStartOfDay(t3.date_stat) = today())  as favoritesBackToday`)
			hs.sqlBuilder.Add(` `)

		case "hostsBackToday":
			hs.sqlBuilder.Add(`sumIf(t3.hosts_day_back, toStartOfDay(t3.date_stat) = today()) as hostsBackToday`)

		case "sessionsBackToday":
			hs.sqlBuilder.Add(`sumIf(t3.sessions_back, toStartOfDay(t3.date_stat) = today()) as	sessionsBackToday`)
			hs.sqlBuilder.Add(` `)

		case "hitsBackToday":
			hs.sqlBuilder.Add(`sumIf(t3.hits_back, toStartOfDay(t3.date_stat) = today()) as hitsBackToday`)
			hs.sqlBuilder.Add(` `)

			//yesterday

		case "guestsYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.guests_day, toStartOfDay(t3.date_stat) = yesterday()) as guestsYesterday`)
			hs.sqlBuilder.Add(` `)

		case "newGuestsYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.new_guests, toStartOfDay(t3.date_stat) = yesterday()) as newGuestsYesterday`)
			hs.sqlBuilder.Add(` `)

		case "favoritesYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.favorites, toStartOfDay(t3.date_stat) = yesterday()) as favoritesYesterday`)
			hs.sqlBuilder.Add(` `)

		case "hostsYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.hosts_day, toStartOfDay(t3.date_stat) = yesterday()) as	hostsYesterday`)
			hs.sqlBuilder.Add(` `)

		case "sessionsYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.sessions, toStartOfDay(t3.date_stat) yesterday()) as sessionsYesterday`)
			hs.sqlBuilder.Add(` `)

		case "hitsYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.hits, toStartOfDay(t3.date_stat) = yesterday()) as hitsYesterday`)
			hs.sqlBuilder.Add(` `)

		case "guestsBackYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.guests_day_back, toStartOfDay(t3.date_stat) = yesterday()) as guestsBackYesterday`)
			hs.sqlBuilder.Add(` `)

		case "favoritesBackYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.favorites_back, toStartOfDay(t3.date_stat) = yesterday()) as favoritesBackYesterday`)
			hs.sqlBuilder.Add(` `)

		case "hostsBackYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.hosts_day_back, toStartOfDay(t3.date_stat) = yesterday()) as hostsBackYesterday`)
			hs.sqlBuilder.Add(` `)

		case "sessionsBackYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.sessions_back, toStartOfDay(t3.date_stat) = yesterday()) as sessionsBackYesterday`)
			hs.sqlBuilder.Add(` `)

		case "hitsBackYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.hits_back, toStartOfDay(t3.date_stat) = yesterday()) as hitsBackYesterday`)
			hs.sqlBuilder.Add(` `)

			// the day before yesterday
		case "guestsBefYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.guests_day, toStartOfDay(t3.date_stat) = (yesterday() - interval 1 day)) 
			as	guestsBefYesterday`)
			hs.sqlBuilder.Add(` `)

		case "newGuestsBefYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.new_guests, toStartOfDay(t3.date_stat) =(yesterday() - interval	1 day)) 
			as newGuestsBefYesterday`)
			hs.sqlBuilder.Add(` `)

		case "favoritesBefYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.favorites, toStartOfDay(t3.date_stat) = (yesterday() - interval 1 day)) 
			as favoritesBefYesterday`)
			hs.sqlBuilder.Add(` `)

		case "hostsBefYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.hosts_day, toStartOfDay(t3.date_stat) = (yesterday() - interval	1 day)) 
			as hostsBefYesterday`)
			hs.sqlBuilder.Add(` `)

		case "sessionsBefYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.sessions, toStartOfDay(t3.date_stat) = (yesterday() - interval 1 day)) 
			as sessionsBefYesterday`)
			hs.sqlBuilder.Add(` `)

		case "hitsBefYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.hits, toStartOfDay(t3.date_stat) = (yesterday() - interval 1 day)) 
			as hitsBefYesterday`)
			hs.sqlBuilder.Add(` `)
		case "guestsBackBefYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.guests_day_back, toStartOfDay(t3.date_stat) = (yesterday() - interval 1 day)) 
			as guestsBackBefYesterday`)
			hs.sqlBuilder.Add(` `)

		case "favoritesBackBefYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.favorites_back, toStartOfDay(t3.date_stat) = (yesterday() - interval
			1 day)) as favoritesBackBefYesterday`)
			hs.sqlBuilder.Add(` `)

		case "hostsBackBefYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.hosts_day_back, toStartOfDay(t3.date_stat) = (yesterday() - interval 1 day)) 
			as	hostsBackBefYesterday`)
			hs.sqlBuilder.Add(` `)

		case "sessionsBackBefYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.sessions_back, toStartOfDay(t3.date_stat) = (yesterday() - interval 1 day)) 
			as sessionsBackBefYesterday`)
			hs.sqlBuilder.Add(` `)

		case "hitsBackBefYesterday":
			hs.sqlBuilder.Add(`sumIf(t3.hits_back, toStartOfDay(t3.date_stat) = (yesterday() - interval 1 day)) 
			as hitsBackBefYesterday`)
			hs.sqlBuilder.Add(` `)
		}
	}
	hs.sqlBuilder.Add(`from adv t1 
          left join adv_stat t2 on t1.uuid = t2.adv_uuid
          left join adv_day t3 on t3.adv_uuid = t2.adv_uuid`)

	return nil
}

func (hs *AdvSqlBuilder) buildWhere() {
	if len(hs.filter.Operators) != 0 {
		hs.sqlBuilder.Add("WHERE ")
		for i := 0; i < len(hs.filter.Operators); i++ {
			op := hs.filter.Operators[i]
			if op.Field == "isRegistered" {
				if op.Value == true {
					hs.sqlBuilder.Add("userId>0 ")
				} else {
					hs.sqlBuilder.Add(" userId=0 ")
				}
				continue
			}

			if op.Operator == "or" {
				hs.sqlBuilder.Add(" OR ")
			} else {
				val := utils.StringConcat(op.Field, op.Operator, "?")
				hs.sqlBuilder.Add(val, op.Value)
			}

			if i+1 < len(hs.filter.Operators)-1 {
				if hs.filter.Operators[i+1].Operator != "or" || (i-1 > 0 && hs.filter.Operators[i-1].Operator != "or") {
					hs.sqlBuilder.Add(" AND ")
				}
			}
		}
	}
}

func (hs *AdvSqlBuilder) buildSkipAndLimit() {
	hs.sqlBuilder.Add(" LIMIT ")
	if hs.filter.Skip != 0 {
		hs.sqlBuilder.Add("?, ", hs.filter.Skip)
	} else {
		hs.sqlBuilder.Add("?, ", 0)
	}

	if hs.filter.Limit != 0 {
		hs.sqlBuilder.Add("?", 0)
	} else if hs.filter.Limit > 1000 || hs.filter.Limit < 0 || hs.filter.Limit == 0 {
		hs.sqlBuilder.Add("?", 1000)
	}
}

func (hs *AdvSqlBuilder) Build() (string, []any, error) {
	if err := hs.buildSelect(); err != nil {
		return "", nil, err
	}

	hs.buildWhere()
	hs.buildSkipAndLimit()

	resultSql, args := hs.sqlBuilder.Build()
	return resultSql, args, nil
}
