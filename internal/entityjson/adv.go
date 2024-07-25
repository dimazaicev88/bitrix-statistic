package entityjson

import (
	"time"
)

type Adv struct {
	Uuid                      string    `json:"uuid"`                      // ID
	Priority                  uint32    `json:"priority"`                  //приоритет
	Referer1                  string    `json:"referer1"`                  //идентификатор referer1
	Referer2                  string    `json:"referer2"`                  //идентификатор referer2
	Description               string    `json:"description"`               //описание
	EventsView                string    `json:"eventsView"`                //режим показа списка событий, возможные значения: link - ссылкой на список, list - списком event1 - сгруппированными по event1; event2 - сгруппированными по event2
	DateFirst                 time.Time `json:"dateFirst"`                 //Дата первого прямого захода
	DateLast                  time.Time `json:"dateLast"`                  //Дата последнего прямого захода или возврата
	AdvTime                   uint32    `json:"advTime"`                   //длительность РК в секундах
	Attent                    uint32    `json:"attent"`                    //Внимательность на прямом заходе
	AttentBack                uint32    `json:"attentBack"`                // Внимательность на возврате
	NewVisitors               float32   `json:"newVisitors"`               //Процент новых посетителей
	ReturnedVisitors          float32   `json:"returnEdvisitors"`          // Процент вернувшихся на сайт посетителей после прямого захода
	VisitorsPerDay            float32   `json:"visitorsPerDay"`            // Среднее кол-во посетителей в день
	Currency                  string    `json:"currency"`                  //Валюта в которой заданы финансовые показатели
	Cost                      float32   `json:"cost"`                      //затраты
	Revenue                   float32   `json:"revenue"`                   //доход
	Benefit                   float32   `json:"benefit"`                   // прибыль
	SessionCost               float32   `json:"sessionCost"`               //Стоимость сессии
	VisitorCost               float32   `json:"visitorCost"`               //Стоимость посетителя
	Roi                       float32   `json:"roi"`                       //рентабельность
	Guests                    uint32    `json:"guests"`                    //Суммарное кол-во посетителей на прямом заходе
	NewGuests                 uint32    `json:"newGuests"`                 //Суммарное кол-во новых посетителей на прямом заходе
	Favorites                 uint32    `json:"favorites"`                 // Суммарное кол-во посетителей, добавившие сайт в "Избранное" на прямом заходе
	Hosts                     uint32    `json:"hosts"`                     // Суммарное кол-во хостов на прямом заходе
	Sessions                  uint32    `json:"sessions"`                  //Суммарное кол-во сессий на прямом заходе
	Hits                      uint32    `json:"hits"`                      // Суммарное кол-во хитов на прямом заходе
	GuestsBack                uint32    `json:"guestsback"`                //Суммарное кол-во посетителей на возврате
	FavoritesBack             uint32    `json:"favoritesBack"`             // Суммарное кол-во посетителей, добавившие сайт в "Избранное" на возврате
	HostsBack                 uint32    `json:"hostsBack"`                 //Суммарное кол-во хостов на возврате
	SessionsBack              uint32    `json:"sessionsBack"`              // Суммарное кол-во сессий на возврате
	HitsBack                  uint32    `json:"hitsBack"`                  //Суммарное кол-во хитов на возврате
	GuestsToday               uint32    `json:"guestsToday"`               //Посетителей на прямом заходе за сегодня
	GuestsBackToday           uint32    `json:"guestsBackToday"`           //Посетителей на возврате за сегодня
	NewGuestsToday            uint32    `json:"newGuestsToday"`            //Новых посетителей на возврате за сегодня
	FavoritesToday            uint32    `json:"favoritesToday"`            //Посетителей, добавившие сайт в "Избранное" на прямом заходе за сегодня
	FavoritesBackToday        uint32    `json:"favoritesBackToday"`        //Посетителей, добавившие сайт в "Избранное" на возврате за сегодня
	HostsToday                uint32    `json:"hostsToday"`                //хостов на прямом заходе за сегодня
	HostsBackToday            uint32    `json:"hostsBackToday"`            // Хостов на возврате за сегодня
	SessionsToday             uint32    `json:"sessionsToday"`             // Сессий на прямом заходе за сегодня
	SessionsBackToday         uint32    `json:"sessionsBackToday"`         //Сессий на возврате за сегодня
	HitsToday                 uint32    `json:"hitsToday"`                 //хитов на прямом заходе за сегодня
	HitsBackToday             uint32    `json:"hitsBackToday"`             //Хитов на возврате за сегодня
	GuestsYesterday           uint32    `json:"guestsYesterday"`           //Посетителей на прямом заходе за вчера
	GuestsBackYesterday       uint32    `json:"guestsBackYesterday"`       // Посетителей на возврате за вчера
	NewGuestsYesterday        uint32    `json:"newGuestsYesterday"`        //Новых посетителей на возврате за вчера
	FavoritesYesterday        uint32    `json:"favoritesYesterday"`        // Посетителей, добавившие сайт в "Избранное" на прямом заходе за вчера
	FavoritesBackYesterday    uint32    `json:"favoritesBackYesterday"`    //Посетителей, добавившие сайт в "Избранное" на возврате за вчера
	HostsYesterday            uint32    `json:"hostsYesterday"`            //хостов на прямом заходе за вчера
	HostsBackYesterday        uint32    `json:"hostsBackYesterday"`        //Хостов на возврате за вчера
	SessionsYesterday         uint32    `json:"sessionsYesterday"`         //Сессий на прямом заходе за вчера
	SessionsBackYesterday     uint32    `json:"sessionsBackYesterday"`     // Сессий на возврате за вчера
	HitsYesterday             uint32    `json:"hitsYesterday"`             //Хитов на прямом заходе за вчера
	HitsBackYesterday         uint32    `json:"hitsBackYesterday"`         //Хитов на возврате за вчера
	GuestsBefYesterday        uint32    `json:"guestsbefyesterday"`        //Посетителей на прямом заходе за позавчера
	GuestsBackBefYesterday    uint32    `json:"guestsbackbefyesterday"`    // Посетителей на возврате за позавчера
	NewGuestsBefYesterday     uint32    `json:"newguestsbefyesterday"`     // Новых посетителей на возврате за позавчера
	FavoritesBefYesterday     uint32    `json:"favoritesbefyesterday"`     //Посетителей, добавившие сайт в "Избранное" на прямом заходе за позавчера
	FavoritesBackBefYesterday uint32    `json:"favoritesbackbefyesterday"` //Посетителей, добавившие сайт в "Избранное" на возврате за позавчера
	CHostsBefYesterday        uint32    `json:"chostsbefyesterday"`        //хостов на прямом заходе за позавчера
	HostsBackBefYesterday     uint32    `json:"hostsbackbefyesterday"`     // Хостов на возврате за позавчера
	SessionsBefYesterday      uint32    `json:"sessionsbefyesterday"`      //Сессий на прямом заходе за позавчера
	SessionsBackBefYesterday  uint32    `json:"sessionsbackbefyesterday"`  // Сессий на возврате за позавчера
	HitsBefYesterday          uint32    `json:"hitsbefyesterday"`          //хитов на прямом заходе за позавчера
	HitsBackBefYesterday      uint32    `json:"hitsbackbefyesterday"`      //Хитов на возврате за позавчера
	GuestsPeriod              uint32    `json:"guestsperiod"`              //Посетителей на прямом заходе за период
	GuestsBackPeriod          uint32    `json:"guestsbackperiod"`          //Посетителей на возврате за период
	NewGuestsPeriod           uint32    `json:"newguestsperiod"`           //Новых посетителей на возврате за период
	FavoritesPeriod           uint32    `json:"favoritesperiod"`           //Посетителей, добавившие сайт в "Избранное" на прямом заходе за период
	FavoritesBackPeriod       uint32    `json:"favoritesbackperiod"`       //Посетителей, добавившие сайт в "Избранное" на возврате за период
	HostsPeriod               uint32    `json:"hostsperiod"`               //Хостов на прямом заходе за период
	HostsBackPeriod           uint32    `json:"hostsbackperiod"`           //Хостов на возврате за период
	SessionsPeriod            uint32    `json:"sessionsperiod"`            //Сессий на прямом заходе за период
	SessionsBackPeriod        uint32    `json:"sessionsbackperiod"`        //Сессий на возврате за период
	HitsPeriod                uint32    `json:"hitsperiod"`                //Хитов на прямом заходе за период
	HitsBackPeriod            uint32    `json:"hitsbackperiod"`            //Хитов на возврате за период
}
