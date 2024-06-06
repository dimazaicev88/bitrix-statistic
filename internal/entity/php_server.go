package entity

type PhpServer struct {
	HttpUserAgent     string `json:"HTTP_USER_AGENT"`
	RemoteAddr        string `json:"REMOTE_ADDR"`
	HttpXForwardedFor string `json:"HTTP_X_FORWARDED_FOR"`
}
