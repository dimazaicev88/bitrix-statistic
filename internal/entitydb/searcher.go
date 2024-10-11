package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type Searcher struct {
	Uuid            uuid.UUID `ch:"uuid"`
	DateCleanup     time.Time `ch:"date_cleanup"`
	TotalHits       uint32    `ch:"total_hits"`
	SaveStatistic   bool      `ch:"save_statistic"`
	Active          bool      `ch:"active"`
	Name            string    `ch:"name"`
	UserAgent       string    `ch:"user_agent"`
	DiagramDefault  bool      `ch:"diagram_default"`
	HitKeepDays     uint32    `ch:"hit_keep_days"`
	DynamicKeepDays uint32    `ch:"dynamic_keep_days"`
	Phrases         uint32    `ch:"phrases"`
	PhrasesHits     uint32    `ch:"phrases_hits"`
	CheckActivity   bool      `ch:"check_activity"`
}

type SearcherDayHits struct {
	Uuid         uuid.UUID `ch:"uuid"`
	DateStat     time.Time `ch:"hit_day"`
	SearcherUuid uuid.UUID `ch:"searcher_uuid"`
	TotalHits    uint32    `ch:"total_hits"`
}

type SearcherHit struct {
	Uuid         uuid.UUID `ch:"uuid"`
	DateHit      time.Time `ch:"date_hit"`
	SearcherUuid string    `ch:"searcher_uuid"`
	Url          string    `ch:"url"`
	Url404       bool      `ch:"url_404"`
	Ip           string    `ch:"ip"`
	UserAgent    string    `ch:"user_agent"`
	SiteId       string    `ch:"site_id"`
}

type SearcherParams struct {
	Uuid         uuid.UUID `ch:"uuid"`
	SearcherUuid uuid.UUID `ch:"searcher_uuid"`
	Domain       string    `ch:"domain"`
	Variable     string    `ch:"variable"`
	CharSet      string    `ch:"char_set"`
}

type PhraseList struct {
	Uuid         uuid.UUID `ch:"uuid"`
	DateHit      time.Time `ch:"date_hit"`
	SearcherUuid uuid.UUID `ch:"searcher_uuid"`
	RefererUuid  uuid.UUID `ch:"referer_uuid"`
	Phrase       string    `ch:"phrase"`
	UrlFrom      string    `ch:"url_from"`
	UrlTo        string    `ch:"url_to"`
	UrlTo404     bool      `ch:"url_to_404"`
	SessionUuid  uuid.UUID `ch:"session_uuid"`
	SiteId       string    `ch:"site_id"`
}

type SearcherPhraseStat struct {
	SearcherUuid uuid.UUID `ch:"searcher_uuid"`
	Phrases      uint32    `ch:"phrases"`
	PhrasesHits  uint32    `ch:"phrases_hits"`
}
