package filters

import "github.com/volatiletech/null/v9"

type BitrixHitFilter struct {
	Filter struct {
		ID         null.Int `json:"ID"`
		GuestId    null.Int `json:"GUEST_ID"`
		SessionId  null.Int `json:"SESSION_ID"`
		StopListId null.Int `json:"STOP_LIST_ID"`
		Url        string   `json:"URL"`
		Url404     string   `json:"URL_404"`
		NewGuest   string   `json:"NEW_GUEST"`
		Registered string   `json:"REGISTERED"`
		Date1      string   `json:"DATE_1"`
		Date2      string   `json:"DATE_2"`
		Ip         string   `json:"IP"`
		UserAgent  string   `json:"USER_AGENT"`
		CountryId  null.Int `json:"COUNTRY_ID"`
		CityId     null.Int `json:"CITY_ID"`
		Cookie     string   `json:"COOKIE"`
		User       string   `json:"USER"`
		Country    string   `json:"COUNTRY"`
		Region     string   `json:"REGION"`
		City       string   `json:"CITY"`
		Stop       string   `json:"STOP"`
		SiteId     string   `json:"SITE_ID"`

		IdExactMatch         string `json:"ID_EXACT_MATCH"`
		UrlExactMatch        string `json:"URL_EXACT_MATCH"`
		UserExactMatch       string `json:"USER_EXACT_MATCH"`
		GuestIdExactMatch    string `json:"GUEST_ID_EXACT_MATCH"`
		SessionIdExactMatch  string `json:"SESSION_ID_EXACT_MATCH"`
		IpExactMatch         string `json:"IP_EXACT_MATCH"`
		UserAgentExactMatch  string `json:"USER_AGENT_EXACT_MATCH"`
		CountryExactMatch    string `json:"COUNTRY_EXACT_MATCH"`
		CountryIdExactMatch  string `json:"COUNTRY_ID_EXACT_MATCH"`
		RegionExactMatch     string `json:"REGION_EXACT_MATCH"`
		CityExactMatch       string `json:"CITY_EXACT_MATCH"`
		CityIdExactMatch     string `json:"CITY_ID_EXACT_MATCH"`
		StopListIdExactMatch string `json:"STOP_LIST_ID_EXACT_MATCH"`
		CookieExactMatch     string `json:"COOKIE_EXACT_MATCH"`
		Url404ExactMatch     string `json:"URL_404_EXACT_MATCH"`
		NewGuestExactMatch   string `json:"NEW_GUEST_EXACT_MATCH"`
	} `json:"filter"`

	OrderBy     string `json:"orderBy"`
	Order       string `json:"order"`
	FilterLogic string `json:"LOGIC"`
}

type FilterOperator struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
	Field    string `json:"field"`
}
