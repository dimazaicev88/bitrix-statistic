package entitydb

import (
	"time"
)

type Hit struct {
	Uuid         string    `ch:"uuid"`
	SessionUuid  string    `ch:"session_uuid"`
	AdvUuid      string    `ch:"adv_uuid"`
	DateHit      time.Time `ch:"date_hit"`
	GuestUuid    string    `ch:"guest_uuid"`
	NewGuest     bool      `ch:"new_guest"`
	UserId       uint32    `ch:"user_id"`
	IsUserAuth   bool      `ch:"user_auth"`
	Url          string    `ch:"url"`
	Url404       bool      `ch:"url_404"`
	UrlFrom      string    `ch:"url_from"`
	Ip           string    `ch:"ip"`
	Method       string    `ch:"method"`
	Cookies      string    `ch:"cookies"`
	UserAgent    string    `ch:"user_agent"`
	StopListUuid string    `ch:"stop_list_uuid"`
	CountryId    string    `ch:"country_id"`
	CityUuid     string    `ch:"city_uuid"`
	SiteId       string    `ch:"site_id"`
}
