package entitydb

type Option struct {
	Name        string      `ch:"name"`
	Value       interface{} `ch:"value"`
	Description string      `ch:"description"`
	SiteId      string      `ch:"siteId"`
}
