package entityjson

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Hit struct {
	Uuid        uuid.UUID        `json:"uuid,omitempty"`         // UUID хита
	SessionUuid uuid.UUID        `json:"sessionUuid,omitempty"`  // UUID сессии
	DateHit     time.Time        `json:"dateHit,omitempty"`      // Время хита
	GuestId     uuid.UUID        `json:"guestUuid,omitempty"`    // UUID посетителя
	NewGuest    bool             `json:"isNewGuest,omitempty"`   // Флаг "был ли это новый посетитель на сайте"
	UserId      sql.Null[uint32] `json:"userId,omitempty"`       // ID пользователя под которым посетитель был авторизован (в момент хита или до этого)
	UserAuth    bool             `json:"userAuth,omitempty"`     // Флаг "был ли посетитель авторизован в момент хита"
	Url         sql.NullString   `json:"url,omitempty"`          // Страница хита
	Url404      bool             `json:"url404,omitempty"`       // Была ли 404 ошибка на странице хита
	UrlFrom     sql.NullString   `json:"urlFrom,omitempty"`      // Страница откуда пришел посетитель
	Ip          sql.NullString   `json:"ip,omitempty"`           // IP адрес посетитель в момент хита
	Method      sql.NullString   `json:"method,omitempty"`       // HTTP метод отсылки данных
	Cookies     sql.NullString   `json:"cookies,omitempty"`      // Содержимое Cookie посетителя в момент хита
	UserAgent   sql.NullString   `json:"userAgent,omitempty"`    // UserAgent посетителя в момент хита
	StopListId  sql.NullInt32    `json:"stopListUuid,omitempty"` // ID записи стоп-листа под которую попал посетитель (если это имело место)
	CountryId   sql.NullString   `json:"countryId,omitempty" `   // ID страны (двух символьный идентификатор) посетителя сайта в момент хита (определяется по IP адресу)
	CountryName sql.NullString   `json:"countryName,omitempty"`  // Название страны посетителя сайта в момент хита (определяется по IP адресу)
	SiteId      sql.NullString   `json:"siteId,omitempty"`       // ID сайта (двух символьный идентификатор)
}
