package entity

type Hit struct {
	Id         int    `json:"ID"`
	SessionId  int    `json:"SESSION_ID"`
	GuestId    int    `json:"GUEST_ID"`
	NewGuest   string `json:"NEW_GUEST"`
	UserId     int    `json:"USER_ID"`
	UserAuth   string `json:"USER_AUTH"`
	Url        string `json:"URL"`
	Url404     string `json:"URL_404"`
	UrlFrom    string `json:"URL_FROM"`
	Ip         string `json:"IP"`
	Method     string `json:"METHOD"`
	Cookies    string `json:"COOKIES"`
	UserAgent  string `json:"USER_AGENT"`
	StopListId int    `json:"STOP_LIST_ID"`
	CountryId  int    `json:"COUNTRY_ID"`
	CityId     int    `json:"CITY_ID"`
	RegionName string `json:"REGION_NAME"`
	CityName   string `json:"CITY_NAME"`
	SiteId     string `json:"SITE_ID"`
	DateHit    string `json:"DATE_HIT"`
}
