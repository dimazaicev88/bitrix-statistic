package dto

type UserData struct {
	PHPSessionId      string `json:"phpSessId"`
	GuestHash         string `json:"guestHash"`
	Url               string `json:"url"`
	Referer           string `json:"referer"`
	Ip                string `json:"ip"`
	UserAgent         string `json:"userAgent"`
	UserId            uint32 `json:"userId"`
	HttpXForwardedFor string `json:"httpXForwardedFor"`
	IsError404        bool   `json:"isError404"`
	SiteId            string `json:"siteId"`
	Lang              string `json:"lang"`
	Event1            string `json:"event1"`
	Event2            string `json:"event2"`
	Event3            string `json:"event3"`
	Method            string `json:"method"`
	Cookies           string `json:"cookies"`
}
