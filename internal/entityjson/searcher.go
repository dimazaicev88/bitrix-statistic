package entityjson

import (
	"time"
)

type Searcher struct {
	Uuid           string    `json:"uuid"`                     // ID поисковой системы
	NAME           string    `json:"NAME,omitempty"`           // Название
	UserAgent      string    `json:"userAgent,omitempty"`      // UserAgent
	DiagramDefault bool      `json:"diagramDefault,omitempty"` // True|False флаг "включать в круговую диаграмму и график по умолчанию"
	DateLast       time.Time `json:"dateLast"`                 // Дата последнего хита
	TotalHits      uint32    `json:"totalHits,omitempty"`      // Суммарное количество хитов
	TodayHits      uint32    `json:"todayHits,omitempty"`      // Количество хитов за сегодня
	YesterdayHits  uint32    `json:"yesterdayHits,omitempty"`  // Количество хитов за вчера
	BYesterdayHits uint32    `json:"BYesterdayHits,omitempty"` // Количество хитов за позавчера
	PeriodHits     uint32    `json:"periodHits,omitempty"`     // Количество хитов за установленный период времени (filter"DATE1_PERIOD"], filter"DATE2_PERIOD"])
}
