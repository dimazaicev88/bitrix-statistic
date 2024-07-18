package entitydb

import "time"

type AdvEventDayDB struct {
	Uuid        string    `ch:"uuid"`
	AdvUuid     string    `ch:"adv_uuid"`
	EventUuid   string    `ch:"event_uuid"`
	DateStat    time.Time `ch:"date_stat"`
	Counter     uint32    `ch:"counter"`
	CounterBack uint32    `ch:"counter_back"`
	Money       float64   `ch:"money"`
	MoneyBack   float64   `ch:"money_back"`
}
