package entitydb

import "time"

type PathCache struct {
	Uuid                string    `ch:"uuid"`
	SessionUuid         string    `ch:"session_uuid"`
	DateHit             time.Time `ch:"date_hit"`
	PathId              int32     `ch:"path_id"`
	PathPages           string    `ch:"path_pages"`
	PathFirstPage       string    `ch:"path_first_page"`
	PathFirstPage404    bool      `ch:"path_first_page_404"`
	PathFirstPageSiteId string    `ch:"path_first_page_site_id"`
	PathLastPage        string    `ch:"path_last_page"`
	PathLastPage404     bool      `ch:"path_last_page_404"`
	PathLastPageSiteId  string    `ch:"path_last_page_site_id"`
	PathSteps           uint32    `ch:"path_steps"`
	IsLastPage          bool      `ch:"is_last_page"`
}
