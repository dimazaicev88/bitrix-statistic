package entity

type Session struct {
	Id            int    `json:"ID,omitempty" db:"ID"`
	GuestId       int    `json:"GUEST_ID,omitempty" db:"GUEST_ID"`
	NewGuest      string `json:"NEW_GUEST,omitempty" db:"GUEST_ID"`
	UserId        int    `json:"USER_ID,omitempty" db:"USER_ID"`
	UserAuth      string `json:"USER_AUTH,omitempty" db:"USER_AUTH"`
	CEvents       int    `json:"C_EVENTS,omitempty" db:"C_EVENTS"`
	Hits          int    `json:"HITS,omitempty" db:"HITS"`
	Favorites     string `json:"FAVORITES,omitempty" db:"FAVORITES"`
	UrlFrom       string `json:"URL_FROM,omitempty" db:"URL_FROM"`
	UrlTo         string `json:"URL_TO,omitempty" db:"URL_TO"`
	UrlTo404      string `json:"URL_TO_404,omitempty" db:"URL_TO_404"`
	UrlLast       string `json:"URL_LAST,omitempty" db:"URL_LAST"`
	UrlLast404    string `json:"URL_LAST_404,omitempty" db:"URL_LAST_404"`
	UserAgent     string `json:"USER_AGENT,omitempty" db:"USER_AGENT"`
	DateStat      int64  `json:"DATE_STAT,omitempty" db:"DATE_STAT"`
	DateFirst     int64  `json:"DATE_FIRST,omitempty" db:"DATE_FIRST"`
	DateLast      int64  `json:"DATE_LAST,omitempty" db:"DATE_LAST"`
	IpFirst       string `json:"IP_FIRST,omitempty" db:"IP_FIRST"`
	IpFirstNumber string `json:"IP_FIRST_NUMBER,omitempty" db:"IP_FIRST_NUMBER"`
	IpLast        string `json:"IP_LAST,omitempty" db:"IP_LAST"`
	IpLastNumber  string `json:"IP_LAST_NUMBER,omitempty" db:"IP_LAST_NUMBER"`
	FirstHitId    int    `json:"FIRST_HIT_ID,omitempty" db:"FIRST_HIT_ID"`
	LastHitId     int    `json:"LAST_HIT_ID,omitempty" db:"LAST_HIT_ID"`
	PhpSessionId  string `json:"PHPSESSID,omitempty" db:"PHPSESSID"`
	AdvId         int    `json:"ADV_ID,omitempty" db:"ADV_ID"`
	AdvBack       string `json:"ADV_BACK,omitempty" db:"ADV_BACK"`
	Referer1      string `json:"REFERER1,omitempty" db:"REFERER1"`
	Referer2      string `json:"REFERER2,omitempty" db:"REFERER2"`
	Referer3      string `json:"REFERER3,omitempty" db:"REFERER3"`
	StopListId    int    `json:"STOP_LIST_ID,omitempty" db:"STOP_LIST_ID"`
	CountryId     string `json:"COUNTRY_ID,omitempty" db:"COUNTRY_ID"`
	CityId        int    `json:"CITY_ID,omitempty" db:"CITY_ID"`
	FirstSiteId   string `json:"FIRST_SITE_ID,omitempty" db:"FIRST_SITE_ID"`
	LastSiteId    string `json:"LAST_SITE_ID,omitempty" db:"LAST_SITE_ID"`
}
