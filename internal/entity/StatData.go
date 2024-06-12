package entity

type StatData struct {
	PhpSessionId      string `json:"phpSessionId"`
	Token             string `json:"token"`
	Url               string `json:"url"`
	Referer           string `json:"referer"`
	Ip                string `json:"ip"`
	UserAgent         string `json:"userAgent"`
	UserId            int    `json:"userId"`
	HttpXForwardedFor string `json:"httpXForwardedFor"`
}
