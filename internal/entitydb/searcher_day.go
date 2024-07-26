package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type SearcherDayHits struct {
	Uuid         uuid.UUID `ch:"uuid"`
	DateStat     time.Time `ch:"hit_day"`
	SearcherUuid uuid.UUID `ch:"searcher_uuid"`
	TotalHits    uint64    `ch:"total_hits"`
}
