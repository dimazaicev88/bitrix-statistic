package filters

type Filter struct {
	Fields         []string         `json:"fields,omitempty"`
	Skip           int              `json:"skip,omitempty"`
	Limit          int              `json:"limit,omitempty"`
	OrderBy        string           `json:"orderBy,omitempty"`
	Order          string           `json:"order,omitempty"`
	FilterOperator []FilterOperator `json:"filterOperator"`
}

type FilterOperator struct {
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
	Field    string      `json:"field"`
}

//Filter struct {
//	ID         null.Int `json:"ID"`
//	GuestId    null.Int `json:"GUEST_ID"`
//	SessionId  null.Int `json:"SESSION_ID"`
//	StopListId null.Int `json:"STOP_LIST_ID"`
//	Url        string   `json:"URL"`
//	Url404     string   `json:"URL_404"`
//	NewGuest   string   `json:"NEW_GUEST"`
//	Registered string   `json:"REGISTERED"`
//	Date1      string   `json:"DATE_1"`
//	Date2      string   `json:"DATE_2"`
//	Ip         string   `json:"IP"`
//	UserAgent  string   `json:"USER_AGENT"`
//	CountryId  null.Int `json:"COUNTRY_ID"`
//	CityId     null.Int `json:"CITY_ID"`
//	Cookie     string   `json:"COOKIE"`
//	User       string   `json:"USER"`
//	Country    string   `json:"COUNTRY"`
//	Region     string   `json:"REGION"`
//	City       string   `json:"CITY"`
//	Stop       string   `json:"STOP"`
//	SiteId     string   `json:"SITE_ID"`
//} `json:"filter"`
