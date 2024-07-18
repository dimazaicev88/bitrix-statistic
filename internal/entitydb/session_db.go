package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type SessionDb struct {
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
	Url          string    `ch:"url"`
	Url404       bool      `ch:"url_last_404"`
	UserAgent    string    `ch:"user_agent"`
	DateStat     time.Time `ch:"date_stat"`
	Date         int64     `ch:"date"`
	Ip           string    `ch:"ip"`
	HitId        uuid.UUID `ch:"hit_uuid"`
	PhpSessionId string    `ch:"phpsessid"`
	AdvId        string    `ch:"adv_id"`
	AdvBack      string    `ch:"adv_back"`
	Referer1     string    `ch:"referer1"`
	Referer2     string    `ch:"referer2"`
	Referer3     string    `ch:"referer3"`
	StopListUuid uuid.UUID `ch:"stop_list_uuid"`
	CountryId    uuid.UUID `ch:"country_uuid"`
	CityUuid     string    `ch:"city_uuid"`
	SiteId       string    `ch:"site_id"`
}
