package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type Hit struct {
	Uuid         uuid.UUID `ch:"uuid"`
	SessionUuid  uuid.UUID `ch:"sessionUuid"`
	AdvUuid      uuid.UUID `ch:"advUuid"`
	PhpSessionId string    `ch:"phpSessionId"`
	DateHit      time.Time `ch:"dateHit"`
	GuestUuid    uuid.UUID `ch:"guestUuid"`
	IsNewGuest   bool      `ch:"isNewGuest"`
	UserId       uint32    `ch:"userId"`
	IsUserAuth   bool      `ch:"userAuth"`
	Url          string    `ch:"url"`
	Url404       bool      `ch:"url404"`
	UrlFrom      string    `ch:"urlFrom"`
	Ip           string    `ch:"ip"`
	Method       string    `ch:"method"`
	Cookies      string    `ch:"cookies"`
	UserAgent    string    `ch:"userAgent"`
	StopListUuid uuid.UUID `ch:"stopListUuid"`
	CountryId    string    `ch:"countryId"`
	CityUuid     uuid.UUID `ch:"cityUuid"`
	SiteId       string    `ch:"siteId"`
}
