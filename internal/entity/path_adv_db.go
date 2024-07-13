package entity

import "time"

type PathAdvDB struct {
	Uuid                string    `ch:"uuid"`
	AdvUuid             string    `ch:"adv_uuid"`
	PathUuid            string    `ch:"path_uuid"`
	DateStat            time.Time `ch:"date_stat"`
	Counter             uint32    `ch:"counter"`
	CounterBack         uint32    `ch:"counter_back"`
	CounterFullPath     uint32    `ch:"counter_full_path"`
	CounterFullPathBack uint32    `ch:"counter_full_path_back"`
	Steps               uint32    `ch:"steps"`
}
