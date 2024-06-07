package entity

import "time"

type SearcherDay struct {
	Id         int
	DateStat   time.Time
	DateLast   time.Time
	SearcherId int
	TotalHits  int
}
