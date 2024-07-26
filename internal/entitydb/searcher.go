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
