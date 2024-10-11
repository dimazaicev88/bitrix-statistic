package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type City struct {
	Uuid      uuid.UUID `ch:"uuid"`
	CountryId string    `ch:"country_id"`
	Region    string    `ch:"region"`
	Name      string    `ch:"name"`
	XmlId     string    `ch:"xml_id"`
	Sessions  uint32    `ch:"sessions"`
	NewGuests uint32    `ch:"new_guests"`
	Hits      uint32    `ch:"hits"`
	Events    uint32    `ch:"events"`
}

type CityDay struct {
	Uuid      uuid.UUID `ch:"uuid"`
	CityUuid  string    `ch:"city_uuid"`
	DateStat  time.Time `ch:"date_stat"`
	Sessions  uint32    `ch:"sessions"`
	NewGuests uint32    `ch:"new_guests"`
	Hits      uint32    `ch:"hits"`
	Events    uint32    `ch:"events"`
}

type CityIP struct {
	StartIp   string `ch:"start_ip"`
	EndIp     string `ch:"end_ip"`
	CountryId string `ch:"country_id"`
	CityUuid  string `ch:"city_uuid"`
}
