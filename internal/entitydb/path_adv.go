package entitydb

import "time"

type PathAdv struct {
	Uuid                string    `ch:"uuid"`
	AdvUuid             string    `ch:"adv_uuid"`
	PathId              int32     `ch:"path_id"`
	DateStat            time.Time `ch:"date_stat"`
	Counter             uint32    `ch:"counter"`
	CounterBack         uint32    `ch:"counter_back"`
	CounterFullPath     uint32    `ch:"counter_full_path"`
	CounterFullPathBack uint32    `ch:"counter_full_path_back"`
	Steps               uint32    `ch:"steps"`
	Sign                int32     `ch:"sign"`
	Version             uint32    `ch:"version"`
}
