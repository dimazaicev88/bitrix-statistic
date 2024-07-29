package entityjson

type StatData struct {
	PHPSessionId      string `json:"phpsessid"`
	GuestHash         string `json:"guestHash"`
	Url               string `json:"url"`
	Referer           string `json:"referer"`
	Ip                string `json:"ip"`
	UserAgent         string `json:"userAgent"`
	UserId            uint32 `json:"userId"`
	UserLogin         string `json:"userLogin"`
	HttpXForwardedFor string `json:"httpXForwardedFor"`
	IsError404        bool   `json:"isError404"`
	SiteId            string `json:"siteId"`
	Event1            string `json:"event1"`
	Event2            string `json:"event2"`
	IsUserAuth        bool   `json:"isUserAuth"`
	Method            string `json:"method"`
	Cookies           string `json:"cookies"`
}
