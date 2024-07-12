package entity

type CountryDB struct {
	Uuid      string `ch:"uuid"`
	ShortName string `ch:"short_name"`
	Name      string `ch:"name"`
	Sessions  uint32 `ch:"sessions"`
	NewGuests uint32 `ch:"new_guests"`
	Hits      uint32 `ch:"hits"`
	Events    uint32 `ch:"events"`
}
