package entitydb

type Browser struct {
	Uuid      string `ch:"uuid"`
	UserAgent string `ch:"user_agent"`
}
