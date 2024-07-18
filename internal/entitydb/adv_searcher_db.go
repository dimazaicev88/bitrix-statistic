package entitydb

// AdvSearcherDB TODO Возможно нужно удалить
type AdvSearcherDB struct {
	Uuid         string `ch:"uuid"`
	AdvUuid      string `ch:"adv_uuid"`
	SearcherUuid string `ch:"searcher_uuid"`
}
