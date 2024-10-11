package entitydb

import "time"

type Country struct {
	Uuid      string `ch:"uuid"`
	ShortName string `ch:"short_name"`
	Name      string `ch:"name"`
	Sessions  uint32 `ch:"sessions"`
	NewGuests uint32 `ch:"new_guests"`
	Hits      uint32 `ch:"hits"`
	Events    uint32 `ch:"events"`
}

type CountryDay struct {
	uuid      string    `ch:"uuid"`
	countryId string    `ch:"country_id"`
	dateStat  time.Time `ch:"date_stat"`
	sessions  uint32    `ch:"sessions"`
	newGuests uint32    `ch:"new_guests"`
	hits      uint32    `ch:"hits"`
	events    uint32    `ch:"events"`
}
