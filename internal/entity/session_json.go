package entity

type SessionJson struct {
	Id            int    `json:"id,omitempty"`
	GuestId       int    `json:"guest_id,omitempty"`
	NewGuest      string `json:"new_guest,omitempty"`
	UserId        int    `json:"user_id,omitempty"`
	UserAuth      string `json:"user_auth,omitempty"`
	CEvents       int    `json:"c_events,omitempty"`
	Hits          int    `json:"hits,omitempty"`
	Favorites     string `json:"favorites,omitempty"`
	UrlFrom       string `json:"url_from,omitempty"`
	UrlTo         string `json:"url_to,omitempty"`
	UrlTo404      string `json:"url_to_404,omitempty"`
	UrlLast       string `json:"url_last,omitempty"`
	UrlLast404    string `json:"url_last_404,omitempty"`
	UserAgent     string `json:"user_agent,omitempty"`
	DateStat      int64  `json:"date_stat,omitempty"`
	DateFirst     int64  `json:"date_first,omitempty"`
	DateLast      int64  `json:"date_last,omitempty"`
	IpFirst       string `json:"ip_first,omitempty"`
	IpFirstNumber string `json:"ip_first_number,omitempty"`
	IpLast        string `json:"ip_last,omitempty"`
	IpLastNumber  string `json:"ip_last_number,omitempty"`
	FirstHitId    int    `json:"first_hit_id,omitempty"`
	LastHitId     int    `json:"last_hit_id,omitempty"`
	PhpSessionId  string `json:"phpsessid,omitempty"`
	AdvId         int    `json:"adv_id,omitempty"`
	AdvBack       string `json:"adv_back,omitempty"`
	Referer1      string `json:"referer1,omitempty"`
	Referer2      string `json:"referer2,omitempty"`
	Referer3      string `json:"referer3,omitempty"`
	StopListId    int    `json:"stop_list_id,omitempty"`
	CountryId     string `json:"country_id,omitempty"`
	CityId        int    `json:"city_id,omitempty"`
	FirstSiteId   string `json:"first_site_id,omitempty"`
	LastSiteId    string `json:"last_site_id,omitempty"`
}
