package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type Hit struct {
	Uuid         uuid.UUID `ch:"uuid"`
	SessionUuid  uuid.UUID `ch:"session_uuid"`
	AdvUuid      uuid.UUID `ch:"adv_uuid"`
	PhpSessionId string    `ch:"php_session_id"`
	DateHit      time.Time `ch:"date_hit"`
	GuestUuid    uuid.UUID `ch:"guest_uuid"`
	IsNewGuest   bool      `ch:"is_new_guest"`
	UserId       uint32    `ch:"user_id"`
	IsUserAuth   bool      `ch:"user_auth"`
	Url          string    `ch:"url"`
	Url404       bool      `ch:"url_404"`
	UrlFrom      string    `ch:"url_from"`
	Ip           string    `ch:"ip"`
	Method       string    `ch:"method"`
	Cookies      string    `ch:"cookies"`
	UserAgent    string    `ch:"user_agent"`
	Favorites    bool      `ch:"favorites"`
	StopListUuid uuid.UUID `ch:"stop_list_uuid"`
	CountryId    string    `ch:"country_id"`
	CityUuid     uuid.UUID `ch:"city_uuid"`
	SiteId       string    `ch:"site_id"`
}
