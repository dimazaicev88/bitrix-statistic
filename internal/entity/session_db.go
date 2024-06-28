package entity

type SessionDb struct {
	Id            int    `db:"id"`
	GuestId       int    `db:"guest_id"`
	NewGuest      string `db:"new_guest"`
	UserId        int    `db:"user_id"`
	UserAuth      string `db:"user_auth"`
	CEvents       int    `db:"c_events"`
	Hits          int    `db:"hits"`
	Favorites     string `db:"favorites"`
	UrlFrom       string `db:"url_from"`
	UrlTo         string `db:"url_to"`
	UrlTo404      string `db:"url_to_404"`
	UrlLast       string `db:"url_last"`
	UrlLast404    string `db:"url_last_404"`
	UserAgent     string `db:"user_agent"`
	DateStat      int64  `db:"date_stat"`
	DateFirst     int64  `db:"date_first"`
	DateLast      int64  `db:"date_last"`
	IpFirst       string `db:"ip_first"`
	IpFirstNumber string `db:"ip_first_number"`
	IpLast        string `db:"ip_last"`
	IpLastNumber  string `db:"ip_last_number"`
	FirstHitId    int    `db:"first_hit_id"`
	LastHitId     int    `db:"last_hit_id"`
	PhpSessionId  string `db:"phpsessid"`
	AdvId         int    `db:"adv_id"`
	AdvBack       string `db:"adv_back"`
	Referer1      string `db:"referer1"`
	Referer2      string `db:"referer2"`
	Referer3      string `db:"referer3"`
	StopListId    int    `db:"stop_list_id"`
	CountryId     string `db:"country_id"`
	CityId        int    `db:"city_id"`
	FirstSiteId   string `db:"first_site_id"`
	LastSiteId    string `db:"last_site_id"`
}
