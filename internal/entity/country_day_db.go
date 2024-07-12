package entity

import "time"

type CountryDayDB struct {
	uuid      string    `ch:"uuid"`
	countryId string    `ch:"country_id"`
	dateStat  time.Time `ch:"date_stat"`
	sessions  uint32    `ch:"sessions"`
	newGuests uint32    `ch:"new_guests"`
	hits      uint32    `ch:"hits"`
	events    uint32    `ch:"events"`
}
