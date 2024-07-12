package entity

import (
	"time"
)

type HitJson struct {
	uuid         string    `ch:"uuid"`
	sessionUuid  string    `ch:"session_uuid"`
	dateHit      time.Time `ch:"date_hit"`
	guestUuid    string    `ch:"guest_uuid"`
	newGuest     bool      `ch:"new_guest"`
	userId       uint32    `ch:"user_id"`
	userAuth     bool      `ch:"user_auth"`
	url          string    `ch:"url"`
	url404       bool      `ch:"url_404"`
	urlFrom      string    `ch:"url_from"`
	ip           string    `ch:"ip"`
	method       string    `ch:"method"`
	cookies      string    `ch:"cookies"`
	userAgent    string    `ch:"user_agent"`
	stopListUuid string    `ch:"stop_list_uuid"`
	countryId    string    `ch:"country_id"`
	cityUuid     string    `ch:"city_uuid"`
	siteId       string    `ch:"site_id"`
}
