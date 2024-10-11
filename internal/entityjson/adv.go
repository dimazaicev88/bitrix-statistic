package entityjson

import (
	"github.com/google/uuid"
	"time"
)

type Adv struct {
	Uuid                uuid.UUID `json:"uuid"`                // ID
	Priority            uint32    `json:"priority"`            //приоритет
	Referer1            string    `json:"referer1"`            //идентификатор referer1
	Referer2            string    `json:"referer2"`            //идентификатор referer2
	Description         string    `json:"description"`         //описание
	EventsView          string    `json:"eventsView"`          //режим показа списка событий, возможные значения: link - ссылкой на список, list - списком event1 - сгруппированными по event1; event2 - сгруппированными по event2
	DateFirst           time.Time `json:"dateFirst"`           //Дата первого прямого захода
	DateLast            time.Time `json:"dateLast"`            //Дата последнего прямого захода или возврата
	AdvTime             uint32    `json:"advTime"`             //длительность РК в секундах
	Attent              uint32    `json:"attent"`              //Внимательность на прямом заходе
	AttentBack          uint32    `json:"attentBack"`          // Внимательность на возврате
	NewVisitors         float32   `json:"newVisitors"`         //Процент новых посетителей
	ReturnedVisitors    float32   `json:"returnEdvisitors"`    // Процент вернувшихся на сайт посетителей после прямого захода
	VisitorsPerDay      float32   `json:"visitorsPerDay"`      // Среднее кол-во посетителей в день
	Currency            string    `json:"currency"`            //Валюта в которой заданы финансовые показатели
	Cost                float32   `json:"cost"`                //затраты
	Revenue             float32   `json:"revenue"`             //доход
	Benefit             float32   `json:"benefit"`             // прибыль
	SessionCost         float32   `json:"sessionCost"`         //Стоимость сессии
	VisitorCost         float32   `json:"visitorCost"`         //Стоимость посетителя
	Roi                 float32   `json:"roi"`                 //рентабельность
	Guests              uint32    `json:"guests"`              //Суммарное кол-во посетителей на прямом заходе
	NewGuests           uint32    `json:"newGuests"`           //Суммарное кол-во новых посетителей на прямом заходе
	Favorites           uint32    `json:"favorites"`           // Суммарное кол-во посетителей, добавившие сайт в "Избранное" на прямом заходе
	Hosts               uint32    `json:"hosts"`               // Суммарное кол-во хостов на прямом заходе
	Sessions            uint32    `json:"sessions"`            //Суммарное кол-во сессий на прямом заходе
	Hits                uint32    `json:"hits"`                // Суммарное кол-во хитов на прямом заходе
	GuestsBack          uint32    `json:"guestsback"`          //Суммарное кол-во посетителей на возврате
	FavoritesBack       uint32    `json:"favoritesBack"`       // Суммарное кол-во посетителей, добавившие сайт в "Избранное" на возврате
	HostsBack           uint32    `json:"hostsBack"`           //Суммарное кол-во хостов на возврате
	SessionsBack        uint32    `json:"sessionsBack"`        // Суммарное кол-во сессий на возврате
	HitsBack            uint32    `json:"hitsBack"`            //Суммарное кол-во хитов на возврате
	GuestsPeriod        uint32    `json:"guestsperiod"`        //Посетителей на прямом заходе за период
	GuestsBackPeriod    uint32    `json:"guestsbackperiod"`    //Посетителей на возврате за период
	NewGuestsPeriod     uint32    `json:"newguestsperiod"`     //Новых посетителей на возврате за период
	FavoritesPeriod     uint32    `json:"favoritesperiod"`     //Посетителей, добавившие сайт в "Избранное" на прямом заходе за период
	FavoritesBackPeriod uint32    `json:"favoritesbackperiod"` //Посетителей, добавившие сайт в "Избранное" на возврате за период
	HostsPeriod         uint32    `json:"hostsperiod"`         //Хостов на прямом заходе за период
	HostsBackPeriod     uint32    `json:"hostsbackperiod"`     //Хостов на возврате за период
	SessionsPeriod      uint32    `json:"sessionsperiod"`      //Сессий на прямом заходе за период
	SessionsBackPeriod  uint32    `json:"sessionsbackperiod"`  //Сессий на возврате за период
	HitsPeriod          uint32    `json:"hitsperiod"`          //Хитов на прямом заходе за период
	HitsBackPeriod      uint32    `json:"hitsbackperiod"`      //Хитов на возврате за период
}
