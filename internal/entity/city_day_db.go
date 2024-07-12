package entity

import "time"

type CityDayDB struct {
	Uuid      string    `ch:"uuid"`
	CityUuid  string    `ch:"city_uuid"`
	DateStat  time.Time `ch:"date_stat"`
	Sessions  uint32    `ch:"sessions"`
	NewGuests uint32    `ch:"new_guests"`
	Hits      uint32    `ch:"hits"`
	Events    uint32    `ch:"events"`
}
