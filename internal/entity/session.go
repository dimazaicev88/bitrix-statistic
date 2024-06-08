package entity

type Session struct {
	Id            int    `json:"id,omitempty" db:"id"`
	GuestId       int    `json:"guest_id,omitempty" db:"guest_id"`
	NewGuest      string `json:"new_guest,omitempty" db:"new_guest"`
	UserId        int    `json:"user_id,omitempty" db:"user_id"`
	UserAuth      string `json:"user_auth,omitempty" db:"user_auth"`
	CEvents       int    `json:"c_events,omitempty" db:"c_events"`
	Hits          int    `json:"hits,omitempty" db:"hits"`
	Favorites     string `json:"favorites,omitempty" db:"favorites"`
	UrlFrom       string `json:"url_from,omitempty" db:"url_from"`
	UrlTo         string `json:"url_to,omitempty" db:"url_to"`
	UrlTo404      string `json:"url_to_404,omitempty" db:"url_to_404"`
	UrlLast       string `json:"url_last,omitempty" db:"url_last"`
	UrlLast404    string `json:"url_last_404,omitempty" db:"url_last_404"`
	UserAgent     string `json:"user_agent,omitempty" db:"user_agent"`
	DateStat      int64  `json:"date_stat,omitempty" db:"date_stat"`
	DateFirst     int64  `json:"date_first,omitempty" db:"date_first"`
	DateLast      int64  `json:"date_last,omitempty" db:"date_last"`
	IpFirst       string `json:"ip_first,omitempty" db:"ip_first"`
	IpFirstNumber string `json:"ip_first_number,omitempty" db:"ip_first_number"`
	IpLast        string `json:"ip_last,omitempty" db:"ip_last"`
	IpLastNumber  string `json:"ip_last_number,omitempty" db:"ip_last_number"`
	FirstHitId    int    `json:"first_hit_id,omitempty" db:"first_hit_id"`
	LastHitId     int    `json:"last_hit_id,omitempty" db:"last_hit_id"`
	PhpSessionId  string `json:"phpsessid,omitempty" db:"phpsessid"`
	AdvId         int    `json:"adv_id,omitempty" db:"adv_id"`
	AdvBack       string `json:"adv_back,omitempty" db:"adv_back"`
	Referer1      string `json:"referer1,omitempty" db:"referer1"`
	Referer2      string `json:"referer2,omitempty" db:"referer2"`
	Referer3      string `json:"referer3,omitempty" db:"referer3"`
	StopListId    int    `json:"stop_list_id,omitempty" db:"stop_list_id"`
	CountryId     string `json:"country_id,omitempty" db:"country_id"`
	CityId        int    `json:"city_id,omitempty" db:"city_id"`
	FirstSiteId   string `json:"first_site_id,omitempty" db:"first_site_id"`
	LastSiteId    string `json:"last_site_id,omitempty" db:"last_site_id"`
}
