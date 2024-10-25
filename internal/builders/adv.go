package builders

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/utils"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"slices"
	"strings"
)

var advSelectFields = []string{
	"uuid", // ID
	//"PRIORITY",    // приоритет //TODO добавить
	"referer1",    // идентификатор referer1
	"referer2",    // идентификатор referer2
	"description", // описание
	//[EVENTS_VIEW// режим показа списка событий, возможные значения:
	//	link - ссылкой на список
	//	list - списком
	//	event1 - сгруппированными по event1
	//	event2 - сгруппированными по event2
	//"dateFirst",                 // дата первого прямого захода //TODO добавить
	//"dateLast",                  // дата последнего прямого захода или возврата //TODO добавить
	//"advTime",                   // длительность РК в секундах //TODO добавить
	"attent",     // внимательность на прямом заходе
	"attentBack", // внимательность на возврате
	//"newVisitors", // процент новых посетителей //TODO добавить
	//"returnedVisitors", // процент вернувшихся на сайт посетителей после прямого захода //TODO добавить
	//"visitorsPerDay",            // среднее кол-во посетителей в день //TODO добавить
	//"currency",                  // валюта в которой заданы финансовые показатели //TODO добавить
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
	"hostsToday",                // хостов на прямом заходе за сегодня
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
	//"guestsPeriod",              // посетителей на прямом заходе за период  //TODO добавить
	//"guestsBackPeriod",          // посетителей на возврате за период //TODO добавить
	//"newGuestsPeriod",           // новых посетителей на возврате за период //TODO добавить
	//"favoritesPeriod",           // посетителей, добавившие сайт в "Избранное" на прямом заходе за период //TODO добавить
	//"favoritesBackPeriod",       // посетителей, добавившие сайт в "Избранное" на возврате за период //TODO добавить
	//"hostsPeriod",               // хостов на прямом заходе за период //TODO добавить
	//"hostsBackPeriod",           // хостов на возврате за период //TODO добавить
	//"sessionsPeriod",            // сессий на прямом заходе за период //TODO добавить
	//"sessionsBackPeriod",        // сессий на возврате за период //TODO добавить
	//"hitsPeriod",     // хитов на прямом заходе за период //TODO добавить
	//"hitsBackPeriod", // хитов на возврате за период //TODO добавить
}

type AdvSqlBuilder struct {
	filter        filters.Filter
	sqlBuilder    *SqlBuilder
	groupByFields mapset.Set[string]
}

func NewAdvSQLBuilder(filter filters.Filter) AdvSqlBuilder {
	return AdvSqlBuilder{
		filter:        filter,
		sqlBuilder:    NewSqlBuilder(),
		groupByFields: mapset.NewSet[string](),
	}
}

// TODO добавить сборку когда нету выбираемых полей
func (hs *AdvSqlBuilder) buildSelect() error {
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
	tmpListFields := make([]string, 0, len(hs.filter.Fields))
	for _, fieldName := range hs.filter.Fields {
		if fieldName == "" {
			continue
		}
		if !slices.Contains(advSelectFields, fieldName) {
			return fmt.Errorf("unknown field: %s", fieldName)
		}

		if val, ok := simpleFields[fieldName]; ok {
			tmpListFields = append(tmpListFields, val)
			hs.groupByFields.Add(val)
			continue
		}

		switch fieldName {
		// today
		case "guestsToday":
			tmpListFields = append(tmpListFields, `sumIf(t3.guestsDay, toStartOfDay(t3.dateStat) = today()) 
                                                           as guestsToday`)

		case "newGuestsToday":
			tmpListFields = append(tmpListFields, `sumIf(t3.newGuests, toStartOfDay(t3.dateStat) = today()) as newGuestsToday`)

		case "favoritesToday":
			tmpListFields = append(tmpListFields, `sumIf(t3.favorites, toStartOfDay(t3.dateStat) = today()) as favoritesToday`)

		case "hostsToday":
			tmpListFields = append(tmpListFields, `sumIf(t3.hostsDay, toStartOfDay(t3.dateStat) = today()) as hostsToday`)

		case "sessionsToday":
			tmpListFields = append(tmpListFields, `sumIf(t3.sessions, toStartOfDay(t3.dateStat) = today()) as sessionsToday`)

		case "hitsToday":
			tmpListFields = append(tmpListFields, `sumIf(t2.hits, toStartOfDay(t3.dateStat) = today()) as hitsToday`)
			hs.groupByFields.Add("t2.hits")

		case "guestsBackToday":
			tmpListFields = append(tmpListFields, `sumIf(t3.guestsDayBack, toStartOfDay(t3.dateStat) = today()) as guestsBackToday`)

		case "favoritesBackToday":
			tmpListFields = append(tmpListFields, `sumIf(t3.favoritesBack, toStartOfDay(t3.dateStat) = today())  as favoritesBackToday`)

		case "hostsBackToday":
			tmpListFields = append(tmpListFields, `sumIf(t3.hostsDayBack, toStartOfDay(t3.dateStat) = today()) as hostsBackToday`)

		case "sessionsBackToday":
			tmpListFields = append(tmpListFields, `sumIf(t3.sessionsBack, toStartOfDay(t3.dateStat) = today()) as	sessionsBackToday`)

		case "hitsBackToday":
			tmpListFields = append(tmpListFields, `sumIf(t3.hitsBack, toStartOfDay(t3.dateStat) = today()) as hitsBackToday`)

			//yesterday
		case "guestsYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.guestsDay, toStartOfDay(t3.dateStat) = yesterday()) as guestsYesterday`)

		case "newGuestsYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.newGuests, toStartOfDay(t3.dateStat) = yesterday()) as newGuestsYesterday`)

		case "favoritesYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.favorites, toStartOfDay(t3.dateStat) = yesterday()) as favoritesYesterday`)

		case "hostsYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.hostsDay, toStartOfDay(t3.dateStat) = yesterday()) as	hostsYesterday`)

		case "sessionsYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.sessions, toStartOfDay(t3.dateStat) = yesterday()) as sessionsYesterday`)

		case "hitsYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.hits, toStartOfDay(t3.dateStat) = yesterday()) as hitsYesterday`)

		case "guestsBackYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.guestsDayBack, toStartOfDay(t3.dateStat) = yesterday()) as guestsBackYesterday`)

		case "favoritesBackYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.favoritesBack, toStartOfDay(t3.dateStat) = yesterday()) as favoritesBackYesterday`)

		case "hostsBackYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.hostsDayBack, toStartOfDay(t3.dateStat) = yesterday()) as hostsBackYesterday`)

		case "sessionsBackYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.sessionsBack, toStartOfDay(t3.dateStat) = yesterday()) as sessionsBackYesterday`)

		case "hitsBackYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.hitsBack, toStartOfDay(t3.dateStat) = yesterday()) as hitsBackYesterday`)

			// the day before yesterday
		case "guestsBefYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.guestsDay, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) 
			as	guestsBefYesterday`)

		case "newGuestsBefYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.newGuests, toStartOfDay(t3.dateStat) = (yesterday() - interval	1 day)) 
			as newGuestsBefYesterday`)

		case "favoritesBefYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.favorites, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) 
			as favoritesBefYesterday`)

		case "hostsBefYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.hostsDay, toStartOfDay(t3.dateStat) = (yesterday() - interval	1 day)) 
			as hostsBefYesterday`)

		case "sessionsBefYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.sessions, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) 
			as sessionsBefYesterday`)

		case "hitsBefYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.hits, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) 
			as hitsBefYesterday`)
		case "guestsBackBefYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.guestsDayBack, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) 
			as guestsBackBefYesterday`)

		case "favoritesBackBefYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.favoritesBack, toStartOfDay(t3.dateStat) = (yesterday() - interval
			1 day)) as favoritesBackBefYesterday`)

		case "hostsBackBefYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.hostsDayBack, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) 
			as	hostsBackBefYesterday`)

		case "sessionsBackBefYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.sessionsBack, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) 
			as sessionsBackBefYesterday`)

		case "hitsBackBefYesterday":
			tmpListFields = append(tmpListFields, `sumIf(t3.hitsBack, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) 
			as hitsBackBefYesterday`)

		//audience
		case "attent":
			tmpListFields = append(tmpListFields, `if(t2.sessions > 0, round(t2.hits / t2.sessions, 2), -1) as attent`)
			hs.groupByFields.Add("t2.sessions")

		case "attentBack":
			tmpListFields = append(tmpListFields, `if(t2.sessionsBack > 0, round(t2.hitsBack / t2.sessionsBack, 2), -1) as attentBack`)
			hs.groupByFields.Add("t2.sessionsBack")
			hs.groupByFields.Add("t2.hitsBack")

		//finances
		case "cost":
			tmpListFields = append(tmpListFields, `round(t1.cost * 1.00, 2) as cost`)
			hs.groupByFields.Add("t1.cost")

		case "revenue":
			tmpListFields = append(tmpListFields, `round(t2.revenue * 1.00, 2) as revenue`)
			hs.groupByFields.Add("t2.revenue")

		case "benefit":
			tmpListFields = append(tmpListFields, `round((t2.revenue - t1.cost) * 1.00, 2) as benefit`)
			hs.groupByFields.Add("t2.revenue")

		case "sessionCost":
			tmpListFields = append(tmpListFields, `round((if(t2.sessions > 0, t1.cost / t2.sessions, null)) * 1.00, 2) as sessionCost`)
			hs.groupByFields.Add("t2.sessions")

		case "visitorCost":
			tmpListFields = append(tmpListFields, `round((if(t2.guests > 0, t1.cost / t2.guests, null)) * 1.00, 2) as visitorCost`)
			hs.groupByFields.Add("t2.guests")

		case "roi":
			tmpListFields = append(tmpListFields, `if(t1.cost > 0, round(((t2.revenue - t1.cost) / t1.cost) * 100, 2), -1) as roi`)
			hs.groupByFields.Add("t1.cost")

		default:
			fmt.Println(fieldName)
		}
	}

	hs.sqlBuilder.Add(strings.Join(tmpListFields, ","))
	hs.sqlBuilder.Add(` from adv t1 
          left join adv_stat t2 on t1.uuid = t2.advUuid
          left join adv_day t3 on t3.advUuid = t2.advUuid`)

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

func (hs *AdvSqlBuilder) buildOrder() {
	if len(hs.filter.Order) > 0 {
		hs.sqlBuilder.Add(" ORDER BY ")
		hs.sqlBuilder.Add(strings.Join(hs.filter.Order, ","))

		if hs.filter.OrderBy != "" {
			hs.sqlBuilder.Add(" ")
			hs.sqlBuilder.Add(hs.filter.OrderBy)
		} else {
			hs.sqlBuilder.Add(" DESC ")
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
		hs.sqlBuilder.Add("?", hs.filter.Limit)
	} else if hs.filter.Limit > 1000 || hs.filter.Limit < 0 || hs.filter.Limit == 0 {
		hs.sqlBuilder.Add("?", 1000)
	}
}

func (hs *AdvSqlBuilder) Build() (string, []any, error) {
	if err := hs.buildSelect(); err != nil {
		return "", nil, err
	}

	hs.buildWhere()
	if hs.groupByFields.IsEmpty() == false {
		hs.sqlBuilder.Add(" GROUP BY ")
		listGroupByFields := hs.groupByFields.ToSlice()
		slices.Sort(listGroupByFields)
		hs.sqlBuilder.Add(strings.Join(listGroupByFields, ","))
		hs.sqlBuilder.Add(" ")
	}

	hs.buildOrder()
	hs.buildSkipAndLimit()

	resultSql, args := hs.sqlBuilder.Build()
	return resultSql, args, nil
}
