package entitydb

type BrowserDB struct {
	Uuid      string `ch:"uuid"`
	UserAgent string `ch:"user_agent"`
}
