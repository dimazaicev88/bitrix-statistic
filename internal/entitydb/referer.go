package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type Referer struct {
	Uuid      uuid.UUID `ch:"uuid"`
	DateFirst time.Time `ch:"date_first"`
	DateLast  time.Time `ch:"date_last"`
	SiteName  string    `ch:"site_name"`
	Sessions  uint32    `ch:"sessions"`
	Hits      uint32    `ch:"hits"`
}

type RefererList struct {
	Uuid        uuid.UUID `ch:"uuid"`
	RefererUuid uuid.UUID `ch:"referer_uuid"`
	DateHit     time.Time `ch:"date_hit"`
	Protocol    string    `ch:"protocol"`
	SiteName    string    `ch:"site_name"`
	UrlFrom     string    `ch:"url_from"`
	UrlTo       string    `ch:"url_to"`
	UrlTo404    bool      `ch:"url_to_404"`
	SessionUuid uuid.UUID `ch:"session_uuid"`
	AdvUuid     uuid.UUID `ch:"adv_uuid"`
	SiteId      string    `ch:"site_id"`
}
