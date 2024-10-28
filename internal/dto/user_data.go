package dto

import "github.com/google/uuid"

type UserData struct {
	PHPSessionId      string    `json:"phpsessid"`
	GuestUuid         uuid.UUID `json:"guestUuid"`
	Url               string    `json:"url"`
	Referer           string    `json:"referer"`
	Ip                string    `json:"ip"`
	UserAgent         string    `json:"userAgent"`
	UserId            uint32    `json:"userId"`
	UserLogin         string    `json:"userLogin"`
	HttpXForwardedFor string    `json:"httpXForwardedFor"`
	IsError404        bool      `json:"isError404"`
	SiteId            string    `json:"siteId"`
	Lang              string    `json:"lang"`
	Event1            string    `json:"event1"`
	Event2            string    `json:"event2"`
	IsUserAuth        bool      `json:"isUserAuth"`
	Method            string    `json:"method"`
	Cookies           string    `json:"cookies"`
	IsFavorite        bool      `json:"isFavorite"`
}
