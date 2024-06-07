package entity

import "time"

type SearcherHit struct {
	Id          int       `json:"id,omitempty"`
	DateHit     time.Time `json:"date_hit"`
	SearcherId  int       `json:"searcher_id,omitempty"`
	Url         string    `json:"url,omitempty"`
	Url404      string    `json:"url_404,omitempty"`
	Ip          string    `json:"ip,omitempty"`
	UserAgent   string    `json:"user_agent,omitempty"`
	HitKeepDays int       `json:"hit_keep_days,omitempty"`
	SiteId      string    `json:"site_id,omitempty"`
}
