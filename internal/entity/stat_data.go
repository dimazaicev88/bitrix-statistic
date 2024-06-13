package entity

type StatData struct {
	PhpSessionId      string `json:"phpSessionId"`
	CookieToken       string `json:"cookieToken"`
	SessionToken      string `json:"sessionToken"`
	Url               string `json:"url"`
	Referer           string `json:"referer"`
	Ip                string `json:"ip"`
	UserAgent         string `json:"userAgent"`
	UserId            int    `json:"userId"`
	HttpXForwardedFor string `json:"httpXForwardedFor"`
	Error404          string `json:"error404"`
	SiteId            string `json:"siteId"`
}
