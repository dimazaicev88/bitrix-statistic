package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/go-clickhouse/ch"
	"time"
)

type Hit struct {
	ch.CHModel `ch:"partition:toYYYYMM(dateHit)"`

	Uuid         uuid.UUID `ch:"uuid"`
	PhpSessionId string    `ch:"phpSessionId"`
	Event1       string    `ch:"event1"`
	Event2       string    `ch:"event2"`
	DateHit      time.Time `ch:"dateHit,pk"`
	GuestHash    string    `ch:"guestHash"`
	IsNewGuest   bool      `ch:"isNewGuest"`
	UserId       uint32    `ch:"userId"`
	Url          string    `ch:"url"`
	Referer      string    `ch:"referrer"`
	Url404       bool      `ch:"url404"`
	UrlFrom      string    `ch:"urlFrom"`
	Ip           string    `ch:"ip"`
	Method       string    `ch:"method"`
	Cookies      string    `ch:"cookies"`
	UserAgent    string    `ch:"userAgent"`
	SiteId       string    `ch:"siteId"`
}
