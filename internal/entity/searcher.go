package entity

import "time"

type Searcher struct {
	Id              int
	DateCleanup     time.Time
	TotalHits       int
	SaveStatistic   string
	Active          string
	Name            string
	UserAgent       string
	DiagramDefault  string
	HitKeepDays     int
	DynamicKeepDays int
	Phrases         int
	PhrasesHits     int
	CheckActivity   string
}
