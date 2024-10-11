package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type (
	Path struct {
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

	PathAdv struct {
		AdvUuid             uuid.UUID `ch:"adv_uuid"`
		PathId              int32     `ch:"path_id"`
		DateStat            time.Time `ch:"date_stat"`
		Counter             uint32    `ch:"counter"`
		CounterBack         uint32    `ch:"counter_back"`
		CounterFullPath     uint32    `ch:"counter_full_path"`
		CounterFullPathBack uint32    `ch:"counter_full_path_back"`
		Steps               uint32    `ch:"steps"`
		Sign                int8      `ch:"sign"`
		Version             uint32    `ch:"version"`
	}

	PathCache struct {
		Uuid                string    `ch:"uuid"`
		SessionUuid         uuid.UUID `ch:"session_uuid"`
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
		Sign                int       `ch:"sign"`
		Version             uint32    `ch:"version"`
	}
)
