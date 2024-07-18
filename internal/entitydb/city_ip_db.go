package entitydb

type CityIPDB struct {
	StartIp   string `ch:"start_ip"`
	EndIp     string `ch:"end_ip"`
	CountryId string `ch:"country_id"`
	CityUuid  string `ch:"city_uuid"`
}
