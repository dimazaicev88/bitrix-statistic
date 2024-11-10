package models

import (
	"github.com/google/uuid"
	"time"
)

type Hit struct {
	Uuid         uuid.UUID `ch:"uuid"`
	PhpSessionId string    `ch:"phpSessionId"`
	Event1       string    `ch:"event1"`
	Event2       string    `ch:"event2"`
	Event3       string    `ch:"event3"`
	DateHit      time.Time `ch:"dateHit"`
	GuestHash    string    `ch:"guestHash"`
	IsNewGuest   bool      `ch:"isNewGuest"`
	UserId       uint32    `ch:"userId"`
	Url          string    `ch:"url"`
	Referer      string    `ch:"referrer"`
	Url404       bool      `ch:"url404"`
	Ip           string    `ch:"ip"`
	Method       string    `ch:"method"`
	Cookies      string    `ch:"cookies"`
	UserAgent    string    `ch:"userAgent"`
	SiteId       string    `ch:"siteId"`
}
