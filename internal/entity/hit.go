package entity

import "database/sql"

type Hit struct {
	Id         int            `json:"ID,omitempty" db:"ID"`
	SessionId  int            `json:"SESSION_ID,omitempty"`
	GuestId    sql.NullInt32  `json:"GUEST_ID,omitempty"`
	NewGuest   string         `json:"NEW_GUEST,omitempty"`
	UserId     sql.NullInt32  `json:"USER_ID,omitempty"`
	UserAuth   sql.NullString `json:"USER_AUTH,omitempty"`
	Url        sql.NullString `json:"URL,omitempty"`
	Url404     string         `json:"URL_404,omitempty"`
	UrlFrom    sql.NullString `json:"URL_FROM,omitempty"`
	Ip         sql.NullString `json:"IP,omitempty"`
	Method     sql.NullString `json:"METHOD,omitempty"`
	Cookies    sql.NullString `json:"COOKIES,omitempty"`
	UserAgent  sql.NullString `json:"USER_AGENT,omitempty"`
	StopListId sql.NullInt32  `json:"STOP_LIST_ID,omitempty"`
	CountryId  sql.NullInt32  `json:"COUNTRY_ID,omitempty" db:"COUNTRY_ID"`
	CityId     sql.NullInt32  `json:"CITY_ID,omitempty"`
	RegionName sql.NullString `json:"REGION_NAME,omitempty"`
	CityName   sql.NullString `json:"CITY_NAME,omitempty"`
	SiteId     sql.NullString `json:"SITE_ID,omitempty"`
	DateHit    sql.NullString `json:"DATE_HIT,omitempty"`
}
