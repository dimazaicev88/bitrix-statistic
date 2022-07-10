package entity

import "database/sql"

type Hit struct {
	Id           int            `json:"ID"`
	IdExactMatch int            `json:"IdExactMatch,omitempty"`
	SessionId    int            `json:"SESSION_ID"`
	GuestId      sql.NullInt32  `json:"GUEST_ID"`
	NewGuest     string         `json:"NEW_GUEST"`
	UserId       sql.NullInt32  `json:"USER_ID"`
	UserAuth     sql.NullString `json:"USER_AUTH"`
	Url          sql.NullString `json:"URL"`
	Url404       string         `json:"URL_404"`
	UrlFrom      sql.NullString `json:"URL_FROM"`
	Ip           sql.NullString `json:"IP"`
	Method       sql.NullString `json:"METHOD"`
	Cookies      sql.NullString `json:"COOKIES"`
	UserAgent    sql.NullString `json:"USER_AGENT"`
	StopListId   sql.NullInt32  `json:"STOP_LIST_ID"`
	CountryId    sql.NullInt32  `json:"COUNTRY_ID"`
	CityId       sql.NullInt32  `json:"CITY_ID"`
	RegionName   sql.NullString `json:"REGION_NAME"`
	CityName     sql.NullString `json:"CITY_NAME"`
	SiteId       sql.NullString `json:"SITE_ID"`
	DateHit      sql.NullString `json:"DATE_HIT"`
}
