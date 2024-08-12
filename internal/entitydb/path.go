package entitydb

import "time"

type Path struct {
	Uuid            string    `ch:"uuid"`
	PathId          int32     `ch:"path_id"`
	ParentPathId    int32     `ch:"parent_path_id"`
	DateStat        time.Time `ch:"date_stat"`
	Counter         uint32    `ch:"counter"`
	CounterFullPath uint32    `ch:"counter_full_path"`
	CounterAbnormal uint32    `ch:"counter_abnormal"`
	Pages           string    `ch:"pages"`
	FirstPage       string    `ch:"first_page"`
	FirstPageSiteId string    `ch:"page_404"`
	FirstPage404    bool      `ch:"first_page_404"`
	PrevPage        string    `ch:"prev_page"`
	PrevPageHash    int32     `ch:"prev_page_hash"`
	LastPage        string    `ch:"last_page"`
	LastPage404     bool      `ch:"last_page_404"`
	LastPageSiteId  string    `ch:"last_page_site_id"`
	LastPageHash    int32     `ch:"last_page_hash"`
	Steps           uint32    `ch:"steps"`
	Sign            int32     `ch:"sign"`
	Version         uint32    `ch:"version"`
}
