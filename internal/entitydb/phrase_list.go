package entitydb

import "time"

type PhraseList struct {
	Uuid         string    `ch:"uuid"`
	DateHit      time.Time `ch:"date_hit"`
	SearcherUuid string    `ch:"searcher_uuid"`
	RefererUuid  string    `ch:"referer_uuid"`
	Phrase       string    `ch:"phrase"`
	UrlFrom      string    `ch:"url_from"`
	UrlTo        string    `ch:"url_to"`
	UrlTo404     bool      `ch:"url_to_404"`
	SessionUuid  string    `ch:"session_uuid"`
	SiteId       string    `ch:"site_id"`
}
