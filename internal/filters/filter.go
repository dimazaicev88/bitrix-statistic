package filters

type Filter struct {
	Select   []string               `json:"SELECT"`
	Where    string                 `json:"WHERE"`
	OrderBy  []string               `json:"ORDER_BY"`
	TypeSort string                 `json:"TYPE_SORT"`
	Params   map[string]interface{} `json:"PARAMS"`
}
