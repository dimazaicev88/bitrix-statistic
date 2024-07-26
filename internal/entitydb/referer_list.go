package entitydb

import "time"

type RefererList struct {
	Uuid        string    `ch:"uuid"`
	RefererUuid string    `ch:"referer_uuid"`
	DateHit     time.Time `ch:"date_hit"`
	Protocol    string    `ch:"protocol"`
	SiteName    string    `ch:"site_name"`
	UrlFrom     string    `ch:"url_from"`
	UrlTo       string    `ch:"url_to"`
	UrlTo404    bool      `ch:"url_to_404"`
	SessionUuid string    `ch:"session_uuid"`
	AdvUuid     string    `ch:"adv_uuid"`
	SiteId      string    `ch:"site_id"`
}
