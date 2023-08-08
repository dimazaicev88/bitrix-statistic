package entity

import "database/sql"

type Hit struct {
	Id         int            `json:"ID,omitempty" db:"ID"`
	SessionId  int            `json:"SessionId,omitempty"`
	GuestId    sql.NullInt32  `json:"GuestId,omitempty"`
	NewGuest   string         `json:"NewGuest,omitempty"`
	UserId     sql.NullInt32  `json:"USER_ID,omitempty"`
	UserAuth   sql.NullString `json:"USER_AUTH,omitempty"`
	Url        sql.NullString `json:"Url,omitempty"`
	Url404     string         `json:"Url404,omitempty"`
	UrlFrom    sql.NullString `json:"URL_FROM,omitempty"`
	Ip         sql.NullString `json:"Ip,omitempty"`
	Method     sql.NullString `json:"METHOD,omitempty"`
	Cookies    sql.NullString `json:"COOKIES,omitempty"`
	UserAgent  sql.NullString `json:"UserAgent,omitempty"`
	StopListId sql.NullInt32  `json:"StopListId,omitempty"`
	CountryId  sql.NullInt32  `json:"CountryId,omitempty" db:"CountryId"`
	CityId     sql.NullInt32  `json:"CityId,omitempty"`
	RegionName sql.NullString `json:"REGION_NAME,omitempty"`
	CityName   sql.NullString `json:"CITY_NAME,omitempty"`
	SiteId     sql.NullString `json:"SiteId,omitempty"`
	DateHit    sql.NullString `json:"DATE_HIT,omitempty"`
}
