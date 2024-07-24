package entityjson

import (
	"github.com/google/uuid"
	"time"
)

type StatEvent struct {
	Uuid          uuid.UUID `json:"uuid"`                    // ID события
	Event3        string    `json:"event3,omitempty"`        // Дополнительный параметр event3 события
	TypeUuid      uuid.UUID `json:"typeUuid,omitempty"`      // ID типа события
	DateEnter     time.Time `json:"dateEnter"`               // Время создания события
	Event1        string    `json:"event1,omitempty"`        // Идентификатор event1 типа события
	Event2        string    `json:"event2,omitempty"`        // Идентификатор event2 типа события
	Name          string    `json:"name,omitempty"`          // Название типа события
	Event         string    `json:"event,omitempty"`         // Event1 / event2, либо название типа события (если оно задано)
	Description   string    `json:"description,omitempty"`   // Описание типа события
	AdvUuid       uuid.UUID `json:"advUuid,omitempty"`       // ID рекламной кампании
	AdvBack       bool      `json:"advBack,omitempty"`       // Флаг прямого захода (N) или возврата (Y) по рекламной кампании
	CountryId     string    `json:"countryId,omitempty"`     // ID страны посетителя
	CountryName   string    `json:"countryName,omitempty"`   // Название страны посетителя
	SessionUuid   uuid.UUID `json:"sessionUuid,omitempty"`   // ID сессии
	GuestUuid     uuid.UUID `json:"guestUuid,omitempty"`     // ID посетителя
	HitUuid       uuid.UUID `json:"hitUuid,omitempty"`       // ID хита
	RefererUrl    string    `json:"refererUrl,omitempty"`    // Ссылающаяся страница
	RefererSiteId string    `json:"refererSiteId,omitempty"` // ID сайта для ссылающейся страницы
	URL           string    `json:"URL,omitempty"`           // Страница на которой было зафиксировано событие
	SiteId        string    `json:"siteId,omitempty"`        // ID сайта для страницы, на которой было зафиксировано событие
	RedirectUrl   string    `json:"redirectUrl,omitempty"`   // Страница, на которую был перенаправлен посетитель после фиксации события
	Money         float32   `json:"money,omitempty"`         // денежная сумма
	Chargeback    bool      `json:"chargeback,omitempty"`    // True - отрицательная денежная сумма; False - положительная денежная сумма
	Currency      string    `json:"currency,omitempty"`      // Трех символьный идентификатор валюты для денежной суммы
}
