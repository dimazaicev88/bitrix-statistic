package entitydb

import "time"

type PageAdv struct {
	Uuid             string    `ch:"uuid"`
	DateStat         time.Time `ch:"date_stat"`
	PageUuid         string    `ch:"page_uuid"`
	AdvUuid          string    `ch:"adv_uuid"`
	Counter          uint32    `ch:"counter"`
	EnterCounter     uint32    `ch:"enter_counter"`
	ExitCounter      uint32    `ch:"exit_counter"`
	CounterBack      uint32    `ch:"counter_back"`
	EnterCounterBack uint32    `ch:"enter_counter_back"`
	ExitCounterBack  uint32    `ch:"exit_counter_back"`
}
