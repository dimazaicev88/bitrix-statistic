package entityjson

import (
	"github.com/google/uuid"
	"time"
)

type TrafficCommonValues struct {
	TotalHits           uint32 `json:"totalHits"`                     //Суммарное количество хитов за все время ведения статистики
	TotalSessions       uint32 `json:"totalSessions,omitempty"`       //Суммарное количество сессий за все время ведения статистики
	TotalEvents         uint32 `json:"totalEvents,omitempty"`         //Суммарное количество событий за все время ведения статистики
	TotalHosts          uint32 `json:"totalHosts,omitempty"`          //Суммарное количество хостов за все время ведения статистики
	TotalGuests         uint32 `json:"totalGuests,omitempty"`         //Суммарное количество посетителей за все время ведения статистики
	TotalFavorites      uint32 `json:"totalFavorites,omitempty"`      //Суммарное количество посетителей, добавивших сайт в "Избранное" за все время ведения статистики
	TodayHits           uint32 `json:"todayHits,omitempty"`           //Количество хитов за сегодня
	TodaySessions       uint32 `json:"todaySessions,omitempty"`       //Количество сессий за сегодня
	TodayEvents         uint32 `json:"todayEvents,omitempty"`         //Количество событий за сегодня
	TodayHosts          uint32 `json:"todayHosts,omitempty"`          //Количество хостов за сегодня
	TodayGuests         uint32 `json:"todayGuests,omitempty"`         //Суммарное количество посетителей за сегодня
	TodayNewGuests      uint32 `json:"todayNewGuests,omitempty"`      //Количество новых посетителей  за сегодня
	TodayFavorites      uint32 `json:"todayFavorites,omitempty"`      //Количество посетителей, добавивших сайт в "Избранное", за сегодня
	YesterdayHits       uint32 `json:"yesterdayHits,omitempty"`       //Количество хитов за вчера
	YesterdaySessions   uint32 `json:"yesterdaySessions,omitempty"`   //Количество сессий за вчера
	YesterdayEvents     uint32 `json:"yesterdayEvents,omitempty"`     //Количество событий за вчера
	YesterdayHosts      uint32 `json:"yesterdayHosts,omitempty"`      //Количество хостов за вчера
	YesterdayGuests     uint32 `json:"yesterdayGuests,omitempty"`     //Суммарное количество посетителей за вчера
	YesterdayNewGuests  uint32 `json:"yesterdayNewGuests,omitempty"`  //Количество новых посетителей за вчера
	YesterdayFavorites  uint32 `json:"yesterdayFavorites,omitempty"`  //Количество посетителей, добавивших сайт в "Избранное", за вчера
	BYesterdayHits      uint32 `json:"BYesterdayHits,omitempty"`      //Количество хитов за позавчера
	BYesterdaySessions  uint32 `json:"BYesterdaySessions,omitempty"`  //Количество сессий за позавчера
	BYesterdayEvents    uint32 `json:"BYesterdayEvents,omitempty"`    //Количество событий за позавчера
	BYesterdayHosts     uint32 `json:"BYesterdayHosts,omitempty"`     //Количество хостов за позавчера
	BYesterdayGuests    uint32 `json:"BYesterdayGuests,omitempty"`    //Суммарное количество посетителей за позавчера
	BYesterdayNewGuests uint32 `json:"BYesterdayNewGuests,omitempty"` //Количество новых посетителей за позавчера
	BYesterdayFavorites uint32 `json:"BYesterdayFavorites,omitempty"` //Количество посетителей, добавивших сайт в "Избранное", за позавчера
	PeriodHits          uint32 `json:"periodHits,omitempty"`          //Количество хитов за установленный период времени (filterDATE1], filterDATE2])
	PeriodSessions      uint32 `json:"periodSessions,omitempty"`      //Количество сессий за установленный период времени
	PeriodEvents        uint32 `json:"periodEvents,omitempty"`        //Количество событий за установленный период времени
	PeriodNewGuests     uint32 `json:"periodNewGuests,omitempty"`     //Количество новых посетителей за установленный период времени
	PeriodFavorites     uint32 `json:"periodFavorites,omitempty"`     //Количество посетителей, добавивших сайт в "Избранное", за установленный период времени
	OnlineGuests        uint32 `json:"onlineGuests,omitempty"`        //Количество посетителей в online
}

type TrafficDailyList struct {
	Uuid     uuid.UUID `json:"uuid,omitempty"`  //ID записи
	DateStat time.Time `json:"dateStat"`        //дата
	Day      uint8     `json:"day,omitempty"`   //день (1-31)
	Month    uint8     `json:"month,omitempty"` //Месяц (1-12)
	Year     uint8     `json:"year,omitempty"`  //год
	Wday     uint8     `json:"wday,omitempty"`  //Номер дня недели
	//(0 - понедельник,
	//1 - вторник, ...
	//6 - воскресенье)
	HITS          uint32 `json:"HITS,omitempty"`          //Количество хитов
	HOSTS         uint32 `json:"HOSTS,omitempty"`         //Количество хостов
	SESSIONS      uint32 `json:"SESSIONS,omitempty"`      //Количество сессий
	EVENTS        uint32 `json:"EVENTS,omitempty"`        //Количество событий
	GUESTS        uint32 `json:"GUESTS,omitempty"`        //Количество посетителей
	NewGuests     uint32 `json:"newGuests,omitempty"`     //Количество новых посетителей
	FAVORITES     uint32 `json:"FAVORITES,omitempty"`     //Количество посетителей, добавивших сайт в "Избранное"
	TotalHosts    uint32 `json:"totalHosts,omitempty"`    //Количество хостов
	AmAverageTime uint32 `json:"amAverageTime,omitempty"` //Среднее время длительности сессии
	Am1           uint32 `json:"am1,omitempty"`           //Количество сессий длительность которых менее 1 минуты
	Am1m3m        uint32 `json:"am1M3M,omitempty"`        //Количество сессий длительность которых от 1 до 3 минут
	Am3m6m        uint32 `json:"am3M6M,omitempty"`        //Количество сессий длительность которых от 3 до 6 минут
	Am6m9m        uint32 `json:"am6M9M,omitempty"`        //Количество сессий длительность которых от 6 до 9 минут
	Am9m12m       uint32 `json:"am9M12M,omitempty"`       //Количество сессий длительность которых от 9 до 12 минут
	Am12m15m      uint32 `json:"am12M15M,omitempty"`      //Количество сессий длительность которых от 12 до 15 минут
	Am15m18m      uint32 `json:"am15M18M,omitempty"`      //Количество сессий длительность которых от 15 до 18 минут
	Am18m21m      uint32 `json:"am18M21M,omitempty"`      //Количество сессий длительность которых от 18 до 21 минут
	Am21m24m      uint32 `json:"am21M24M,omitempty"`      //Количество сессий длительность которых от 21 до 24 минут
	Am24m         uint32 `json:"am24M,omitempty"`         //Количество сессий длительность которых более 24 минут
	AhAverageHits uint32 `json:"ahAverageHits,omitempty"` //Среднее количество хитов в сессии
	Ah1h          uint32 `json:"ah1H,omitempty"`          //Количество сессий в которых был только 1 хит
	Ah2h5h        uint32 `json:"ah2H5H,omitempty"`        //Количество сессий в которых было от 2 до 5 хитов
	Ah6h9h        uint32 `json:"ah6H9H,omitempty"`        //Количество сессий в которых было от 6 до 9 хитов
	Ah10h13h      uint32 `json:"ah10H13H,omitempty"`      //Количество сессий в которых было от 10 до 13 хитов
	Ah14h17h      uint32 `json:"ah14H17H,omitempty"`      //Количество сессий в которых было от 14 до 17 хитов
	Ah18h21h      uint32 `json:"ah18H21H,omitempty"`      //Количество сессий в которых было от 18 до 21 хитов
	Ah22h25h      uint32 `json:"ah22H25H,omitempty"`      //Количество сессий в которых было от 22 до 25 хитов
	Ah26h29h      uint32 `json:"ah26H29H,omitempty"`      //Количество сессий в которых было от 26 до 29 хитов
	Ah30h33h      uint32 `json:"ah30H33H,omitempty"`      //Количество сессий в которых было от 30 до 33 хитов
	Ah34h         uint32 `json:"ah34H,omitempty"`         //Количество сессий в которых было более 34 хитов
}

type TrafficPhraseList struct {
	PhraseValue       string `json:"phrase,omitempty"`            //Поисковая фраза
	TotalPhrases      uint32 `json:"totalPhrases,omitempty"`      //Суммарное количество заходов с данной поисковой фразой
	TodayPhrases      uint32 `json:"todayPhrases,omitempty"`      //Сколько раз сегодня заходили с данной поисковой фразой
	YesterdayPhrases  uint32 `json:"yesterdayPhrases,omitempty"`  //Сколько раз вчера заходили с данной поисковой фразой
	BYesterdayPhrases uint32 `json:"BYesterdayPhrases,omitempty"` //Сколько раз позавчера заходили с данной поисковой фразой
	PeriodPhrases     uint32 `json:"periodPhrases,omitempty"`     //Сколько раз заходили с данной поисковой фразой за установленный период времени (filterDATE1], filterDATE2])
}

type TrafficRefererList struct {
	SiteName           string `json:"siteName,omitempty"`           //Ссылающийся сайт
	TotalReferers      uint32 `json:"totalReferers,omitempty"`      //Суммарное количество заходов с данного ссылающегося сайта
	TodayReferers      uint32 `json:"todayReferers,omitempty"`      //Количество заходов со ссылающегося сайта за сегодня
	YesterdayReferers  uint32 `json:"yesterdayReferers,omitempty"`  //Количество заходов со ссылающегося сайта за вчера
	BYesterdayReferers uint32 `json:"BYesterdayReferers,omitempty"` //Количество заходов со ссылающегося сайта за позавчера
	PeriodReferers     uint32 `json:"periodReferers,omitempty"`     //Количество заходов со ссылающегося сайта за установленный период времени (filterDATE1], filterDATE2])
}

type TrafficSumListHour struct {
	HourHost0      uint32 `json:"hourHost0,omitempty"`      //Число хостов, зафиксированных с 0:00 до 1:00 часа
	HourHost23     uint32 `json:"hourHost23,omitempty"`     //Число хостов, зафиксированных с 23:00 до 0:00
	HourSession0   uint32 `json:"hourSession0,omitempty"`   //Число сессий, зафиксированных с 0:00 до 1:00 часа
	HourSession23  uint32 `json:"hourSession23,omitempty"`  //Число сессий, зафиксированных с 23:00 до 0:00
	HourHit0       uint32 `json:"hourHit0,omitempty"`       //Число хитов, зафиксированных с 0:00 до 1:00 часа
	HourHit23      uint32 `json:"hourHit23,omitempty"`      //Число хитов, зафиксированных с 23:00 до 0:00
	HourEvent0     uint32 `json:"hourEvent0,omitempty"`     //Число событий, зафиксированных с 0:00 до 1:00 часа
	HourEvent23    uint32 `json:"hourEvent23,omitempty"`    //Число событий, зафиксированных с 23:00 до 0:00
	HourGuest0     uint32 `json:"hourGuest0,omitempty"`     //Число посетителей, зафиксированных с 0:00 до 1:00 часа
	HourGuest23    uint32 `json:"hourGuest23,omitempty"`    //Число посетителей, зафиксированных с 23:00 до 0:00
	HourNewGuest0  uint32 `json:"hourNewGuest0,omitempty"`  //Число новых посетителей, зафиксированных с 0:00 до 1:00 часа
	HourNewGuest23 uint32 `json:"hourNewGuest23,omitempty"` //Число новых посетителей, зафиксированных с 23:00 до 0:00
	HourFavorite0  uint32 `json:"hourFavorite0,omitempty"`  //Число посетителей, добавивших сайт в "Избранное", зафиксированных с 0:00 до 1:00 часа
	HourFavorite23 uint32 `json:"hourFavorite23,omitempty"` //Число посетителей, добавивших сайт в "Избранное", зафиксированных с 23:00 до 0:00
}
