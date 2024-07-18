package entitydb

import "time"

type SearcherDayDb struct {
	Id         int       `ch:"uuid"`
	DateStat   time.Time `ch:"date_stat"`
	DateLast   time.Time `ch:"date_last"`
	SearcherId int       `ch:"searcher_uuid"`
	TotalHits  int       `ch:"total_hits"`
}
