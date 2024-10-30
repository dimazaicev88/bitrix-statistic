package dto

import (
	"github.com/google/uuid"
	"time"
)

type (
	Adv struct {
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
		GuestsBack          uint32    `json:"guestsBack"`          //Суммарное кол-во посетителей на возврате
		FavoritesBack       uint32    `json:"favoritesBack"`       // Суммарное кол-во посетителей, добавившие сайт в "Избранное" на возврате
		HostsBack           uint32    `json:"hostsBack"`           //Суммарное кол-во хостов на возврате
		SessionsBack        uint32    `json:"sessionsBack"`        // Суммарное кол-во сессий на возврате
		HitsBack            uint32    `json:"hitsBack"`            //Суммарное кол-во хитов на возврате
		GuestsPeriod        uint32    `json:"guestsPeriod"`        //Посетителей на прямом заходе за период
		GuestsBackPeriod    uint32    `json:"guestsBackPeriod"`    //Посетителей на возврате за период
		NewGuestsPeriod     uint32    `json:"newGuestsPeriod"`     //Новых посетителей на возврате за период
		FavoritesPeriod     uint32    `json:"favoritesPeriod"`     //Посетителей, добавившие сайт в "Избранное" на прямом заходе за период
		FavoritesBackPeriod uint32    `json:"favoritesBackPeriod"` //Посетителей, добавившие сайт в "Избранное" на возврате за период
		HostsPeriod         uint32    `json:"hostsPeriod"`         //Хостов на прямом заходе за период
		HostsBackPeriod     uint32    `json:"hostsBackPeriod"`     //Хостов на возврате за период
		SessionsPeriod      uint32    `json:"sessionsPeriod"`      //Сессий на прямом заходе за период
		SessionsBackPeriod  uint32    `json:"sessionsBackPeriod"`  //Сессий на возврате за период
		HitsPeriod          uint32    `json:"hitsPeriod"`          //Хитов на прямом заходе за период
		HitsBackPeriod      uint32    `json:"hitsBackPeriod"`      //Хитов на возврате за период
	}

	AdvDynamicResult struct {
		AdvDynamic []AdvDynamic `json:"advDynamic"`
		AdvMaxMin  *AdvMaxMin   `json:"advMaxMin"`
	}

	AdvDynamic struct {
		DateStat      time.Time `json:"dateStat,omitempty"`      //Дата
		MonthDay      uint8     `json:"day"`                     //День (1-31)
		Month         uint8     `json:"month,omitempty"`         //Месяц (1-12)
		Year          uint16    `json:"year,omitempty"`          //Год
		Guests        uint32    `json:"guests,omitempty"`        //Посетителей на прямом заходе
		NewGuests     uint32    `json:"newGuests,omitempty"`     //Новых посетителей на прямом заходе
		Favorites     uint32    `json:"favorites,omitempty"`     //Посетителей, добавивших сайт в "избранное" на прямом заходе
		Hosts         uint32    `json:"hosts,omitempty"`         //Хостов на прямом заходе
		Sessions      uint32    `json:"sessions,omitempty"`      //Сессий на прямом заходе
		Hits          uint32    `json:"hits,omitempty"`          //Хитов на прямом заходе
		GuestsBack    uint32    `json:"guestsBack,omitempty"`    //Посетителей на возврате
		FavoritesBack uint32    `json:"favoritesBack,omitempty"` //Посетителей, добавивших сайт в "избранное" на возврате
		HostsBack     uint32    `json:"hostsBack,omitempty"`     //Хостов на возврате
		SessionsBack  uint32    `json:"sessionsBack,omitempty"`  //Сессий на возврате
		HitsBack      uint32    `json:"hitsBack,omitempty"`      //Хитов на возврате

	}

	AdvMaxMin struct {
		DateFirst time.Time `json:"dateFirst"`          //Минимальная дата
		MinDay    uint8     `json:"minDay,omitempty"`   //День минимальной даты (1-31)
		MinMonth  uint8     `json:"minMonth,omitempty"` //Месяц минимальной даты (1-12)
		MinYear   uint16    `json:"minYear,omitempty"`  //Год минимальной даты
		DateLast  time.Time `json:"dateLast"`           //Максимальная дата
		MaxDay    uint8     `json:"maxDay,omitempty"`   //День максимальной даты (1-31)
		MaxMonth  uint8     `json:"maxMonth,omitempty"` //Месяц максимальной даты (1-12)
		MaxYear   uint16    `json:"maxYear,omitempty"`  //Год максимальной даты
	}

	AdvEvent struct {
		EventTypeUuid           uuid.UUID `json:"eventTypeUuid,omitempty"`           //Uuid типа события
		Event1                  string    `json:"event1,omitempty"`                  //идентификатор Event1
		Event2                  string    `json:"event2,omitempty"`                  //идентификатор event2
		Name                    string    `json:"name,omitempty"`                    //название
		Event                   string    `json:"event,omitempty"`                   //название либо id] Event1 / event2
		Description             string    `json:"description,omitempty"`             //описание
		Counter                 uint64    `json:"counter,omitempty"`                 //Суммарное кол-во событий данного типа на прямом заходе по рк
		CounterBack             uint64    `json:"counterBack,omitempty"`             //Суммарное кол-во событий данного типа на возврате по рк
		CounterToday            uint64    `json:"counterToday,omitempty"`            //Кол-во событий данного типа на прямом заходе по рк за сегодня
		CounterYesterday        uint64    `json:"counterYesterday,omitempty"`        //Кол-во событий данного типа на прямом заходе по рк за вчера
		CounterBefYesterday     uint64    `json:"counterBefYesterday,omitempty"`     //Кол-во событий данного типа на прямом заходе по рк за позавчера
		CounterPeriod           uint64    `json:"counterPeriod,omitempty"`           //Кол-во событий данного типа на прямом заходе по рк за период
		CounterBackToday        uint64    `json:"counterBackToday,omitempty"`        //Кол-во событий данного типа на возврате по рк за сегодня
		CounterBackYesterday    uint64    `json:"counterBackYesterday,omitempty"`    //Кол-во событий данного типа на возврате по рк за вчера
		CounterBackBefYesterday uint64    `json:"counterBackBefYesterday,omitempty"` //Кол-во событий данного типа на возврате по рк за позавчера
		CounterBackPeriod       uint64    `json:"counterBackPeriod,omitempty"`       //Кол-во событий данного типа на возврате по рк за период
		Money                   float64   `json:"money,omitempty"`                   //Итоговая денежная сумма событий данного типа на прямом заходе по рк
		MoneyToday              float64   `json:"moneyToday,omitempty"`              //Денежная сумма событий данного типа на прямом заходе по рк за сегодня
		MoneyYesterday          float64   `json:"moneyYesterday,omitempty"`          //Денежная сумма событий данного типа на прямом заходе по рк за вчера
		MoneyBefYesterday       float64   `json:"moneyBefYesterday,omitempty"`       //Денежная сумма событий данного типа на прямом заходе по рк за позавчера
		MoneyPeriod             float64   `json:"moneyPeriod,omitempty"`             //Денежная сумма событий данного типа на прямом заходе по рк за период
		MoneyBack               float64   `json:"moneyBack,omitempty"`               //Итоговая денежная сумма событий данного на возврате по рк
		MoneyBackToday          float64   `json:"moneyBackToday,omitempty"`          //Денежная сумма событий данного типа на возврате по рк за сегодня
		MoneyBackYesterday      float64   `json:"moneyBackYesterday,omitempty"`      //Денежная сумма событий данного типа на возврате по рк за вчера
		MoneyBackBefYesterday   float64   `json:"moneyBackBefYesterday,omitempty"`   //Денежная сумма событий данного типа на возврате по рк за позавчера
		MoneyBackPeriod         float64   `json:"moneyBackPeriod,omitempty"`         //Денежная сумма событий данного типа на возврате по рк за период
	}

	Simple struct {
		AdvUuid     uuid.UUID `json:"advUuid"`
		Referer1    string    `json:"referer1"`    //идентификатор referer1
		Referer2    string    `json:"referer2"`    //идентификатор referer2
		Description string    `json:"description"` //описание
	}
)
