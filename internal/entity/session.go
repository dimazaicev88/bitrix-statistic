package entity

import "database/sql"

type Session struct {
	Id            int            `json:"ID,omitempty" db:"ID"`
	GuestId       int            `json:"GUEST_ID,omitempty" db:"GUEST_ID"`
	NewGuest      string         `json:"NEW_GUEST,omitempty" db:"GUEST_ID"`
	UserId        sql.NullInt32  `json:"USER_ID,omitempty" db:"USER_ID"`
	UserAuth      sql.NullString `json:"USER_AUTH,omitempty" db:"USER_AUTH"`
	CEvents       int            `json:"C_EVENTS,omitempty" db:"C_EVENTS"`
	Hits          int            `json:"HITS,omitempty" db:"HITS"`
	Favorites     string         `json:"FAVORITES,omitempty" db:"FAVORITES"`
	UrlFrom       sql.NullString `json:"URL_FROM,omitempty" db:"URL_FROM"`
	UrlTo         sql.NullString `json:"URL_TO,omitempty" db:"URL_TO"`
	UrlTo404      string         `json:"URL_TO_404,omitempty" db:"URL_TO_404"`
	UrlLast       sql.NullString `json:"URL_LAST,omitempty" db:"URL_LAST"`
	UrlLast404    string         `json:"URL_LAST_404,omitempty" db:"URL_LAST_404"`
	UserAgent     sql.NullString `json:"USER_AGENT,omitempty" db:"USER_AGENT"`
	DateStat      sql.NullTime   `json:"DATE_STAT,omitempty" db:"DATE_STAT"`
	DateFirst     sql.NullTime   `json:"DATE_FIRST,omitempty" db:"DATE_FIRST"`
	DateLast      sql.NullTime   `json:"DATE_LAST,omitempty" db:"DATE_LAST"`
	IpFirst       sql.NullString `json:"IP_FIRST,omitempty" db:"IP_FIRST"`
	IpFirstNumber sql.NullInt64  `json:"IP_FIRST_NUMBER,omitempty" db:"IP_FIRST_NUMBER"`
	IpLast        sql.NullString `json:"IP_LAST,omitempty" db:"IP_LAST"`
	IpLastNumber  sql.NullString `json:"IP_LAST_NUMBER,omitempty" db:"IP_LAST_NUMBER"`
	FirstHitId    sql.NullInt32  `json:"FIRST_HIT_ID,omitempty" db:"FIRST_HIT_ID"`
	LastHitId     sql.NullInt32  `json:"LAST_HIT_ID,omitempty" db:"LAST_HIT_ID"`
	PhpSessionId  string         `json:"PHPSESSID,omitempty" db:"PHPSESSID"`
	AdvId         sql.NullInt32  `json:"ADV_ID,omitempty" db:"ADV_ID"`
	AdvBack       sql.NullString `json:"ADV_BACK,omitempty" db:"ADV_BACK"`
	Referer1      sql.NullString `json:"REFERER1,omitempty" db:"REFERER1"`
	Referer2      sql.NullString `json:"REFERER2,omitempty" db:"REFERER2"`
	Referer3      sql.NullString `json:"REFERER3,omitempty" db:"REFERER3"`
	StopListId    sql.NullInt32  `json:"STOP_LIST_ID,omitempty" db:"STOP_LIST_ID"`
	CountryId     sql.NullString `json:"COUNTRY_ID,omitempty" db:"COUNTRY_ID"`
	CityId        sql.NullInt32  `json:"CITY_ID,omitempty" db:"CITY_ID"`
	FirstSiteId   sql.NullString `json:"FIRST_SITE_ID,omitempty" db:"FIRST_SITE_ID"`
	LastSiteId    sql.NullString `json:"LAST_SITE_ID,omitempty" db:"LAST_SITE_ID"`
}
