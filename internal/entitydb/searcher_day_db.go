package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type SearcherDayDb struct {
	Id           uuid.UUID `ch:"uuid"`
	DateStat     time.Time `ch:"date_stat"`
	DateLast     time.Time `ch:"date_last"`
	SearcherUuid uuid.UUID `ch:"searcher_uuid"`
	TotalHits    uint64    `ch:"total_hits"`
}
