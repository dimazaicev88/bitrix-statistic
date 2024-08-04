package entitydb

import (
	"time"
)

type Session struct {
	Uuid          string    `ch:"uuid"`
	GuestUuid     string    `ch:"guest_uuid"`
	IsNewGuest    bool      `ch:"new_guest"`
	UserId        uint32    `ch:"user_id"`
	IsUserAuth    bool      `ch:"user_auth"`
	Events        uint32    `ch:"events"`
	Hits          uint32    `ch:"hits"`
	Favorites     bool      `ch:"favorites"`
	UrlFrom       string    `ch:"url_from"`
	UrlTo         string    `ch:"url_to"`
	UrlTo404      bool      `ch:"url_to_404"`
	UrlLast       string    `ch:"url_last"`
	UrlLast404    bool      `ch:"url_last_404"`
	UserAgent     string    `ch:"user_agent"`
	DateStat      time.Time `ch:"date_stat"`
	DateFirst     time.Time `ch:"date_first"`
	DateLast      time.Time `ch:"date_last"`
	IpFirst       string    `ch:"ip_first"`
	IpLast        string    `ch:"ip_last"`
	FirstHitUuid  string    `ch:"first_hit_uuid"`
	LastHitUuid   string    `ch:"last_hit_uuid"`
	PhpSessionId  string    `ch:"phpsessid"`
	AdvUuid       string    `ch:"adv_uuid"`
	AdvBack       string    `ch:"adv_back"`
	Referer1      string    `ch:"referer1"`
	Referer2      string    `ch:"referer2"`
	Referer3      string    `ch:"referer3"`
	StopListUuid  string    `ch:"stop_list_uuid"`
	CountryId     string    `ch:"country_uuid"`
	FirstSiteUuid string    `ch:"first_site_uuid"`
	LastSiteUuid  string    `ch:"last_site_uuid"`
	CityId        string    `ch:"city_id"`
	Sign          int8      `ch:"sign"`
	Version       uint32    `ch:"version"`
}
