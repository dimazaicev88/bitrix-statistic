package entitydb

import "time"

type RefererDB struct {
	Uuid      string    `ch:"uuid"`
	DateFirst time.Time `ch:"date_first"`
	DateLast  time.Time `ch:"date_last"`
	SiteName  string    `ch:"site_name"`
	Sessions  uint32    `ch:"sessions"`
	Hits      uint32    `ch:"hits"`
}
