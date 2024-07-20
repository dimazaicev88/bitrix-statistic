package entitydb

import "time"

type SearcherHitDb struct {
	Uuid       string    `ch:"uuid"`
	DateHit    time.Time `ch:"date_hit"`
	SearcherId string    `ch:"searcher_uuid"`
	Url        string    `ch:"url"`
	Url404     bool      `ch:"url_404"`
	Ip         string    `ch:"ip"`
	UserAgent  string    `ch:"user_agent"`
	SiteId     string    `ch:"site_id"`
}
