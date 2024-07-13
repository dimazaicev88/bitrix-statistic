package entity

type StatData struct {
	PHPSessionId      string `json:"phpsessid"`
	GuestHash         string `json:"guestHash"`
	Url               string `json:"url"`
	Referer           string `json:"referer"`
	Ip                string `json:"ip"`
	UserAgent         string `json:"userAgent"`
	UserId            int    `json:"userId"`
	UserLogin         string `json:"userLogin"`
	HttpXForwardedFor string `json:"httpXForwardedFor"`
	Error404          uint8  `json:"error404"`
	SiteId            string `json:"siteId"`
}
