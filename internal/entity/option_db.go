package entity

type Option struct {
	Name        string `ch:"name"`
	Value       string `ch:"value"`
	Description string `ch:"description"`
	SiteId      string `ch:"siteId"`
}
