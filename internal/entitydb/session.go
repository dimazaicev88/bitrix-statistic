package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	Uuid         uuid.UUID `ch:"uuid"`
	GuestUuid    uuid.UUID `ch:"guest_uuid"`
	IsNewGuest   bool      `ch:"new_guest"`
	UserId       uint32    `ch:"user_id"`
	IsUserAuth   bool      `ch:"user_auth"`
	Events       uint32    `ch:"events"`
	Hits         uint32    `ch:"hits"`
	Favorites    bool      `ch:"favorites"`
	UrlFrom      string    `ch:"url_from"`
	UrlTo        string    `ch:"url_to"`
	UrlTo404     bool      `ch:"url_to_404"`
	UrlLast      string    `ch:"url_last"`
	UrlLast404   bool      `ch:"url_last_404"`
	UserAgent    string    `ch:"user_agent"`
	DateStat     time.Time `ch:"date_stat"`
	DateFirst    time.Time `ch:"date_first"`
	DateLast     time.Time `ch:"date_last"`
	IpFirst      string    `ch:"ip_first"`
	IpLast       string    `ch:"ip_last"`
	FirstHitUuid uuid.UUID `ch:"first_hit_uuid"`
	LastHitUuid  uuid.UUID `ch:"last_hit_uuid"`
	PhpSessionId string    `ch:"phpsessid"`
	AdvUuid      uuid.UUID `ch:"adv_uuid"`
	AdvBack      bool      `ch:"adv_back"`
	Referer1     string    `ch:"referer1"`
	Referer2     string    `ch:"referer2"`
	Referer3     string    `ch:"referer3"`
	StopListUuid uuid.UUID `ch:"stop_list_uuid"`
	CountryId    string    `ch:"country_id"`
	FirstSiteId  string    `ch:"first_site_id"`
	LastSiteId   string    `ch:"last_site_id"`
	CityId       string    `ch:"city_id"`
	Sign         int8      `ch:"sign"`
	Version      uint32    `ch:"version"`
}
