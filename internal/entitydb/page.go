package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type Page struct {
	Uuid         uuid.UUID `ch:"uuid"`
	DateStat     time.Time `ch:"date_stat"`
	Dir          bool      `ch:"dir"`
	Url          string    `ch:"url"`
	Url404       bool      `ch:"url_404"`
	UrlHash      uint32    `ch:"url_hash"`
	SiteId       string    `ch:"site_id"`
	Counter      uint32    `ch:"counter"`
	EnterCounter uint32    `ch:"enter_counter"`
	ExitCounter  uint32    `ch:"exit_counter"`
	Sign         int8      `ch:"sign"`
	Version      uint32    `ch:"version"`
}
