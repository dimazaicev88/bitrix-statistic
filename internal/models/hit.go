package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/go-clickhouse/ch"
	"time"
)

type Hit struct {
	ch.CHModel `ch:"partition:toYYYYMM(time)"`

	Uuid         uuid.UUID `ch:"uuid"`
	PhpSessionId string    `ch:"phpSessionId"`
	Event1       string    `ch:"event1"`
	Event2       string    `ch:"event2"`
	DateHit      time.Time `ch:"dateHit,pk"`
	GuestHash    string    `ch:"guestHash"`
	IsNewGuest   bool      `ch:"isNewGuest"`
	UserId       uint32    `ch:"userId"`
	Url          string    `ch:"url"`
	Referer      string    `ch:"referer"`
	Url404       bool      `ch:"url404"`
	UrlFrom      string    `ch:"urlFrom"`
	Ip           string    `ch:"ip"`
	Method       string    `ch:"method"`
	Cookies      string    `ch:"cookies"`
	UserAgent    string    `ch:"userAgent"`
	Favorites    bool      `ch:"favorites"`
	CountryId    string    `ch:"countryId"`
	CityId       string    `ch:"cityId"`
	SiteId       string    `ch:"siteId"`
}
