package dto

import "github.com/google/uuid"

type Hit struct {
	Uuid         uuid.UUID `json:"uuid,omitempty"`         // UUID хита
	SessionUuid  uuid.UUID `json:"sessionUuid,omitempty"`  // UUID сессии
	DateHit      string    `json:"dateHit,omitempty"`      // Время хита
	GuestUuid    uuid.UUID `json:"guestUuid,omitempty"`    // UUID посетителя
	NewGuest     bool      `json:"isNewGuest,omitempty"`   // Флаг "был ли это новый посетитель на сайте"
	UserId       uint32    `json:"userId,omitempty"`       // ID пользователя под которым посетитель был авторизован (в момент хита или до этого)
	UserAuth     bool      `json:"userAuth,omitempty"`     // Флаг "был ли посетитель авторизован в момент хита"
	Url          string    `json:"url,omitempty"`          // Страница хита
	Url404       bool      `json:"url404,omitempty"`       // Была ли 404 ошибка на странице хита
	UrlFrom      string    `json:"urlFrom,omitempty"`      // Страница откуда пришел посетитель
	Ip           string    `json:"ip,omitempty"`           // IP адрес посетитель в момент хита
	Method       string    `json:"method,omitempty"`       // HTTP метод отсылки данных
	Cookies      string    `json:"cookies,omitempty"`      // Содержимое Cookie посетителя в момент хита
	UserAgent    string    `json:"userAgent,omitempty"`    // UserAgent посетителя в момент хита
	StopListUuid uuid.UUID `json:"stopListUuid,omitempty"` // ID записи стоп-листа под которую попал посетитель (если это имело место)
	CountryId    string    `json:"countryId,omitempty" `   // ID страны (двух символьный идентификатор) посетителя сайта в момент хита (определяется по IP адресу)
	CountryName  string    `json:"countryName,omitempty"`  // Название страны посетителя сайта в момент хита (определяется по IP адресу)
	SiteId       string    `json:"siteId,omitempty"`       // ID сайта (двух символьный идентификатор)
}
