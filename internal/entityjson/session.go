package entityjson

import (
	"time"
)

type Session struct {
	Uuid         string    `json:"uuid"`                  // Uuid сессии
	GuestUuid    string    `json:"guestUuid,omitempty"`   // ID посетителя
	NewGuest     bool      `json:"newGuest,omitempty"`    // Флаг "новый посетитель" (True - новый; False - вернувшийся)
	UserId       uint32    `json:"userId,omitempty"`      // ID пользователя под которым последний раз был авторизован посетитель
	UserAuth     bool      `json:"userAuth,omitempty"`    // Флаг "авторизован ли посетитель в данной сессии" (Y - да; N - нет)
	Events       uint32    `json:"events,omitempty"`      // Количество событий произведенных в данной сессии
	Hits         uint32    `json:"hits,omitempty"`        // Количество хитов произведенных в данной сессии
	Favorites    bool      `json:"favorites,omitempty"`   // Флаг "добавлял ли посетитель сайт в "Избранное" в данной сессии
	StopListUuid string    `json:"stopListId,omitempty"`  // ID записи стоп-листа, под которую попал посетитель (если это имело место быть)
	CountryId    string    `json:"countryId,omitempty"`   // ID страны посетителя
	UserAgent    string    `json:"userAgent,omitempty"`   // UserAgent посетителя
	UrlFrom      string    `json:"urlFrom,omitempty"`     // Ссылающаяся страница для первого хита сессии
	DateStat     time.Time `json:"dateStat,omitempty"`    // Дата первого хита сессии
	DateFirst    time.Time `json:"dateFirst,omitempty"`   // Время первого хита сессии
	UrlTo        string    `json:"urlTo,omitempty"`       // Страница первого хита сессии
	UrlTo404     bool      `json:"urlTo404,omitempty"`    // Флаг 404 ошибки на первой странице сессии (Y - да; N - нет)
	FirstHitUuid string    `json:"firstHitId,omitempty"`  // ID первого хита
	FirstSiteId  string    `json:"firstSiteId,omitempty"` // ID сайта для первого хита сессии
	IpFirst      string    `json:"ipFirst,omitempty"`     // IP адрес посетителя на первом хите сессии (в виде: XXX.XXX.XXX.XXX)
	DateLast     time.Time `json:"dateLast,omitempty"`    // Время последнего хита сессии
	UrlLast      string    `json:"urlLast,omitempty"`     // Страница последнего хита сессии
	UrlLast404   bool      `json:"urlLast404,omitempty"`  // Флаг 404 ошибки на последней странице сессии (Y - да; N - нет)
	LastHitUuid  string    `json:"lastHitId,omitempty"`   // ID последнего хита
	LastSiteId   string    `json:"lastSiteId,omitempty"`  // ID сайта для последнего хита сессии
	IpLast       string    `json:"ipLast,omitempty"`      // IP адрес посетителя на последнем хите сессии (в виде: XXX.XXX.XXX.XXX)
	AdvUuid      string    `json:"advId,omitempty"`       // ID рекламной кампании
	AdvBack      bool      `json:"advBack,omitempty"`     // Флаг прямого захода (N) или возврата (Y) по рекламной кампании
	Referer1     string    `json:"referer1,omitempty"`    // Идентификатор referer1 рекламной кампании
	Referer2     string    `json:"referer2,omitempty"`    // Идентификатор referer2 рекламной кампании
	Referer3     string    `json:"referer3,omitempty"`    // Дополнительный параметр referer3 рекламной кампании
	SessionTime  uint64    `json:"sessionTime,omitempty"` // Разница во времени между первым и последним хитом сессии (сек.)
	Phpsessid    string    `json:"phpsessid,omitempty"`   // Идентификатор сессии выданный PHP
}
