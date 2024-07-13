package entity

import "time"

type PathDB struct {
	Uuid            string    `ch:"uuid"`
	PathId          uint32    `ch:"path_id"`
	ParentPathId    uint32    `ch:"parent_path_id"`
	DateStat        time.Time `ch:"date_stat"`
	Counter         uint32    `ch:"counter"`
	CounterAbnormal uint32    `ch:"counter_abnormal"`
	CounterFullPath uint32    `ch:"counter_full_path"`
	Pages           string    `ch:"pages"`
	FirstPage       string    `ch:"first_page"`
	FirstPage404    bool      `ch:"first_page_404"`
	FirstPageSiteId string    `ch:"first_page_site_id"`
	PrevPage        string    `ch:"prev_page"`
	PrevPageHash    uint32    `ch:"prev_page_hash"`
	LastPage        string    `ch:"last_page"`
	LastPage404     bool      `ch:"last_page_404"`
	LastPageSiteId  string    `ch:"last_page_site_id"`
	LastPageHash    uint32    `ch:"last_page_hash"`
	Steps           uint32    `ch:"steps"`
}
