package entitydb

import "time"

type EventDayDB struct {
	Uuid      string    `ch:"uuid"`
	DateStat  time.Time `ch:"date_stat"`
	DateLast  time.Time `ch:"date_last"`
	EventUuid uint32    `ch:"event_uuid"`
	Money     float64   `ch:"money"`
	Counter   uint32    `ch:"counter"`
}
