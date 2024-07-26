package entitydb

import "time"

type Path struct {
	Uuid            string    `ch:"uuid"`
	PathId          uint32    `ch:"path_id"`
	ParentPathId    uint32    `ch:"parent_path_id"`
	DateStat        time.Time `ch:"date_stat"`
	Counter         uint32    `ch:"counter"`
	CounterAbnormal uint32    `ch:"counter_abnormal"`
	CounterFullPath uint32    `ch:"counter_full_path"`
	Pages           string    `ch:"pages"`
	Page            string    `ch:"page"`
	Page404         bool      `ch:"page_404"`
	PageSiteId      string    `ch:"page_site_id"`
	PrevPage        string    `ch:"prev_page"`
	PrevPageHash    uint32    `ch:"prev_page_hash"`
	PageHash        uint32    `ch:"page_hash"`
	Steps           uint32    `ch:"steps"`
}
