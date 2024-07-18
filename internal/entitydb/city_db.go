package entitydb

type CityDB struct {
	Uuid      string `ch:"uuid"`
	CountryId string `ch:"country_id"`
	Region    string `ch:"region"`
	Name      string `ch:"name"`
	XmlId     string `ch:"xml_id"`
	Sessions  uint32 `ch:"sessions"`
	NewGuests uint32 `ch:"new_guests"`
	Hits      uint32 `ch:"hits"`
	Events    uint32 `ch:"events"`
}
