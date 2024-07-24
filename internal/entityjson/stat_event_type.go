package entityjson

import (
	"github.com/google/uuid"
	"time"
)

type StatEventTypeGroupEvent1 struct {
	Event1            string    `json:"event1,omitempty"`            //Идентификатор event1 типа события
	DateEnter         time.Time `json:"dateEnter"`                   //Дата первого события, тип которого имеет данный идентификатор event1
	DateLast          time.Time `json:"dateLast"`                    //Дата последнего события, тип которого имеет данный идентификатор event1
	TotalCounter      uint32    `json:"totalCounter,omitempty"`      //Суммарное количество событий
	TodayCounter      uint32    `json:"todayCounter,omitempty"`      //Количество событий за сегодня
	YesterdayCounter  uint32    `json:"yesterdayCounter,omitempty"`  //Количество событий за вчера
	BYesterdayCounter uint32    `json:"BYesterdayCounter,omitempty"` //Количество событий за позавчера
	PeriodCounter     uint32    `json:"periodCounter,omitempty"`     //количество событий за период времени, указанный в фильтре (filterDATE1_PERIOD], filterDATE2_PERIOD])
	TotalMoney        float32   `json:"totalMoney,omitempty"`        //суммарная денежная сумма
	TodayMoney        float32   `json:"todayMoney,omitempty"`        //Денежная сумма за сегодня
	YesterdayMoney    float32   `json:"yesterdayMoney,omitempty"`    //Денежная сумма за вчера
	BYesterdayMoney   float32   `json:"BYesterdayMoney,omitempty"`   //денежная сумма за позавчера
	PeriodMoney       float32   `json:"periodMoney,omitempty"`       //денежная сумма за период времени, указанный в фильтре (filterDATE1_PERIOD], filterDATE2_PERIOD])
	Currency          string    `json:"currency,omitempty"`          //Трех символьный идентификатор валюты
}

type StatEventTypeGroupEvent2 struct {
	Event2            string    `json:"event2,omitempty"`            //Идентификатор event2 типа события
	DateEnter         time.Time `json:"dateEnter"`                   //Дата первого события, тип которого имеет данный идентификатор event2
	DateLast          time.Time `json:"dateLast"`                    //Дата последнего события, тип которого имеет данный идентификатор event2
	TotalCounter      uint32    `json:"totalCounter,omitempty"`      //Суммарное количество событий
	TodayCounter      uint32    `json:"todayCounter,omitempty"`      //Количество событий за сегодня
	YesterdayCounter  uint32    `json:"yesterdayCounter,omitempty"`  //Количество событий за вчера
	BYesterdayCounter uint32    `json:"bYesterdayCounter,omitempty"` //Количество событий за позавчера
	PeriodCounter     uint32    `json:"periodCounter,omitempty"`     //количество событий за период времени, указанный в фильтре (filterDATE1_PERIOD], filterDATE2_PERIOD])
	TotalMoney        float32   `json:"totalMoney,omitempty"`        //суммарная денежная сумма
	TodayMoney        float32   `json:"todayMoney,omitempty"`        //Денежная сумма за сегодня
	YesterdayMoney    float32   `json:"yesterdayMoney,omitempty"`    //Денежная сумма за вчера
	BYesterdayMoney   float32   `json:"BYesterdayMoney,omitempty"`   //денежная сумма за позавчера
	PeriodMoney       float32   `json:"periodMoney,omitempty"`       //денежная сумма за период времени, указанный в фильтре (filterDATE1_PERIOD], filterDATE2_PERIOD])
	Currency          string    `json:"currency,omitempty"`          //Трех символьный идентификатор валюты
}

type StatEventTypeGroup struct {
	Uuid              uuid.UUID `json:"uuid"`                        //ID типа события
	Event1            string    `json:"event1,omitempty"`            //идентификатор event1 типа события
	Event2            string    `json:"event2,omitempty"`            //идентификатор event2 типа события
	DiagramDefault    bool      `json:"diagramDefault,omitempty"`    //Y|N] флаг: включать ли данный тип события в круговую диаграмму и график по умолчанию
	Name              string    `json:"name,omitempty"`              //название типа события
	Event             string    `json:"event,omitempty"`             //event1 / event2, либо название типа события (если оно указано)
	Description       string    `json:"description,omitempty"`       //описание типа события
	TotalCounter      uint32    `json:"totalCounter,omitempty"`      //суммарное количество событий данного типа
	TodayCounter      uint32    `json:"todayCounter,omitempty"`      //количество событий данного типа за сегодня
	YesterdayCounter  uint32    `json:"yesterdayCounter,omitempty"`  //количество событий данного типа за вчера
	BYesterdayCounter uint32    `json:"BYesterdayCounter,omitempty"` //количество событий данного типа за позавчера
	PeriodCounter     uint32    `json:"periodCounter,omitempty"`     //количество событий данного типа за период времени, указанный в фильтре (filterDATE1_PERIOD], filterDATE2_PERIOD])
	TotalMoney        float32   `json:"totalMoney,omitempty"`        //суммарная денежная сумма по данному типу события
	TodayMoney        float32   `json:"todayMoney,omitempty"`        //денежная сумма по данному типу события за сегодня
	YesterdayMoney    float32   `json:"yesterdayMoney,omitempty"`    //денежная сумма по данному типу события за вчера
	BYesterdayMoney   float32   `json:"BYesterdayMoney,omitempty"`   //денежная сумма по данному типу события за позавчера
	PeriodMoney       float32   `json:"periodMoney,omitempty"`       //денежная сумма по данному типу события за период времени, указанный в фильтре (filterDATE1_PERIOD], filterDATE2_PERIOD])
	Currency          string    `json:"currency,omitempty"`          //трехсимвольный идентификатор валюты
}
