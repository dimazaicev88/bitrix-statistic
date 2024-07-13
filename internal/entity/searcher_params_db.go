package entity

type SearcherParamsDB struct {
	Uuid         string `ch:"uuid"`
	SearcherUuid string `ch:"searcher_uuid"`
	Domain       string `ch:"domain"`
	Variable     string `ch:"variable"`
	CharSet      string `ch:"char_set"`
}
