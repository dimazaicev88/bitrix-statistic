package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type (
	Page struct {
		Uuid         uuid.UUID `ch:"uuid"`
		DateStat     time.Time `ch:"date_stat"`
		Dir          bool      `ch:"dir"`
		Url          string    `ch:"url"`
		Url404       bool      `ch:"url_404"`
		UrlHash      int32     `ch:"url_hash"`
		SiteId       string    `ch:"site_id"`
		Counter      uint32    `ch:"counter"`
		EnterCounter uint32    `ch:"enter_counter"`
		ExitCounter  uint32    `ch:"exit_counter"`
		Sign         int8      `ch:"sign"`
		Version      uint32    `ch:"version"`
	}

	PageAdv struct {
		Uuid             uuid.UUID `ch:"uuid"`
		DateStat         time.Time `ch:"date_stat"`
		PageUuid         uuid.UUID `ch:"page_uuid"`
		AdvUuid          uuid.UUID `ch:"adv_uuid"`
		Counter          uint32    `ch:"counter"`
		EnterCounter     uint32    `ch:"enter_counter"`
		ExitCounter      uint32    `ch:"exit_counter"`
		CounterBack      uint32    `ch:"counter_back"`
		EnterCounterBack uint32    `ch:"enter_counter_back"`
		ExitCounterBack  uint32    `ch:"exit_counter_back"`
	}
)
