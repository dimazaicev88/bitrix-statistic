package entity

import "time"

type SearcherHitDb struct {
	Uuid        string    `ch:"id"`
	DateHit     time.Time `ch:"date_hit"`
	SearcherId  string    `ch:"searcher_id"`
	Url         string    `ch:"url"`
	Url404      bool      `ch:"url_404"`
	Ip          string    `ch:"ip"`
	UserAgent   string    `ch:"user_agent"`
	HitKeepDays uint32    `ch:"hit_keep_days"`
	SiteId      string    `ch:"site_id"`
}
