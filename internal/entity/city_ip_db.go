package entity

type CityIPDB struct {
	startIp   uint32 `ch:"start_ip"`
	endIp     uint32 `ch:"end_ip"`
	countryId string `ch:"country_id"`
	cityUuid  string `ch:"city_uuid"`
}
