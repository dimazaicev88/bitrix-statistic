package entity

import "time"

type Searcher struct {
	Id              int       `json:"id,omitempty" db:"id"`
	DateCleanup     time.Time `json:"date_cleanup"  db:"date_cleanup"`
	TotalHits       int       `json:"total_hits,omitempty" db:"total_hits"`
	SaveStatistic   string    `json:"save_statistic,omitempty" db:"save_statistic"`
	Active          string    `json:"active,omitempty" db:"active"`
	Name            string    `json:"name,omitempty" db:"name"`
	UserAgent       string    `json:"user_agent,omitempty" db:"user_agent"`
	DiagramDefault  string    `json:"diagram_default,omitempty" db:"diagram_default"`
	HitKeepDays     int       `json:"hit_keep_days,omitempty" db:"hit_keep_days"`
	DynamicKeepDays int       `json:"dynamic_keep_days,omitempty" db:"dynamic_keep_days"`
	Phrases         int       `json:"phrases,omitempty" db:"phrases"`
	PhrasesHits     int       `json:"phrases_hits,omitempty" db:"phrases_hits"`
	CheckActivity   string    `json:"check_activity,omitempty" db:"check_activity"`
}
